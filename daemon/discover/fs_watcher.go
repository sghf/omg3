package discover

import (
	"context"
	"io/fs"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"

	"opensvc.com/opensvc/core/path"
	"opensvc.com/opensvc/core/rawconfig"
	"opensvc.com/opensvc/daemon/msgbus"
	"opensvc.com/opensvc/util/file"
	"opensvc.com/opensvc/util/pubsub"
)

const (
	delayExistAfterRemove = 100 * time.Millisecond
)

func (d *discover) fsWatcherStart() (func(), error) {
	log := d.log.With().Str("func", "fsWatch").Logger()
	bus := pubsub.BusFromContext(d.ctx)
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Error().Err(err).Msg("start")
		return func() {}, err
	}
	cleanup := func() {
		if err = watcher.Close(); err != nil {
			log.Error().Err(err).Msg("close")
		}
	}
	d.fsWatcher = watcher
	for _, filename := range []string{rawconfig.Paths.Etc, rawconfig.Paths.Etc + "/namespaces"} {
		if err := d.fsWatcher.Add(filename); err != nil {
			log.Error().Err(err).Msgf("add %s", filename)
			cleanup()
			return func() {}, err
		} else {
			log.Info().Msgf("add dir %s", filename)
		}
	}
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(d.ctx)
	stop := func() {
		log.Debug().Msg("stop")
		cancel()
		wg.Wait()
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer cleanup()
		log.Info().Msg("started")
		nodeConf := filepath.Join(rawconfig.Paths.Etc, "node.conf")
		const createDeleteMask = fsnotify.Create | fsnotify.Remove
		const needReAddMask = fsnotify.Remove | fsnotify.Rename
		const updateMask = fsnotify.Remove | fsnotify.Rename | fsnotify.Write | fsnotify.Create | fsnotify.Chmod

		// Add directory watch for:
		//  var/node/
		varNodeDir := filepath.Join(rawconfig.Paths.Var, "node")
		nodeFrozenFile := filepath.Join(varNodeDir, "frozen")
		if err := d.fsWatcher.Add(varNodeDir); err != nil {
			log.Error().Err(err).Msgf("add dir watch %s", varNodeDir)
		} else {
			log.Info().Msgf("add dir watch %s", varNodeDir)
		}
		if !file.ModTime(nodeFrozenFile).IsZero() {
			log.Info().Msgf("detect %s initially exists", nodeFrozenFile)
			msgbus.PubFrozenFileUpdate(bus, "", msgbus.FrozenFileUpdated{Filename: nodeFrozenFile})
		}

		//
		// Add directory watch for:
		//  etc/
		//  etc/namespaces/
		//  etc/namespaces/*
		//
		// Add config file watches
		//  etc/*.conf
		//  etc/namespaces/*/*.conf
		//
		err = filepath.WalkDir(
			rawconfig.Paths.Etc,
			func(filename string, entry fs.DirEntry, err error) error {
				switch {
				case entry.IsDir():
					if strings.HasPrefix(filename, rawconfig.Paths.Etc+"/namespaces/") {
						if err := d.fsWatcher.Add(filename); err != nil {
							log.Error().Err(err).Msgf("add dir watch %s", filename)
						} else {
							log.Info().Msgf("add dir watch %s", filename)
						}
					}
				default:
					if !strings.HasSuffix(filename, ".conf") {
						return nil
					}
					var (
						p   path.T
						err error
					)
					if filename == nodeConf {
						// pass
					} else if p, err = cfgFilenameToPath(filename); err != nil {
						log.Warn().Err(err).Msgf("do not watch invalid config file %s", filename)
						return nil
					}
					if err := watcher.Add(filename); err != nil {
						log.Error().Err(err).Msgf("add file %s", filename)
					} else {
						log.Debug().Msgf("add file %s", filename)
					}
					msgbus.PubCfgFileUpdate(bus, "fs_watcher emit cfgfile.update", msgbus.CfgFileUpdated{Path: p, Filename: filename})
				}
				return nil
			},
		)
		if err != nil {
			log.Error().Err(err).Msg("walk")
		}

		// watcher-events handler loop
		for {
			select {
			case <-ctx.Done():
				log.Info().Msg("stopped")
				return
			case e := <-watcher.Errors:
				log.Error().Err(e).Msg("")
			case event := <-watcher.Events:
				log.Debug().Msgf("event: %s", event)
				filename := event.Name
				switch {
				case strings.HasSuffix(filename, "frozen"):
					var (
						p   path.T
						err error
					)
					if filename == nodeFrozenFile {
						// pass
					} else if p, err = frozenFilenameToPath(filename); err != nil {
						log.Warn().Err(err).Msgf("%s", filename)
						continue
					}
					switch {
					case event.Op&fsnotify.Remove != 0:
						log.Debug().Msgf("detect removed file %s (%s)", filename, event.Op)
						msgbus.PubFrozenFileRemove(bus, p.String(), msgbus.FrozenFileRemoved{Path: p, Filename: filename})
					case event.Op&updateMask != 0:
						if event.Op&needReAddMask != 0 {
							time.Sleep(delayExistAfterRemove)
							if !file.Exists(filename) {
								log.Info().Msg("file removed")
								return
							} else {
								if err := watcher.Add(filename); err != nil {
									log.Error().Err(err).Msgf("re-add file watch %s", filename)
								} else {
									log.Debug().Msgf("re-add file watch %s", filename)
								}
							}
						}
						log.Debug().Msgf("detect updated file %s (%s)", filename, event.Op)
						msgbus.PubFrozenFileUpdate(bus, p.String(), msgbus.FrozenFileUpdated{Path: p, Filename: filename})
					}
				case strings.HasSuffix(filename, ".conf"):
					var (
						p   path.T
						err error
					)
					if filename == nodeConf {
						rawconfig.LoadSections()
					} else if p, err = cfgFilenameToPath(filename); err != nil {
						log.Warn().Err(err).Msgf("%s", filename)
						continue
					}
					switch {
					case event.Op&fsnotify.Remove != 0:
						log.Debug().Msgf("detect removed file %s (%s)", filename, event.Op)
						msgbus.PubCfgFileRemove(bus, p.String(), msgbus.CfgFileRemoved{Path: p, Filename: filename})
					case event.Op&updateMask != 0:
						if event.Op&needReAddMask != 0 {
							time.Sleep(delayExistAfterRemove)
							if !file.Exists(filename) {
								log.Info().Msg("file removed")
								return
							} else {
								if err := watcher.Add(filename); err != nil {
									log.Error().Err(err).Msgf("re-add file watch %s", filename)
								} else {
									log.Debug().Msgf("re-add file watch %s", filename)
								}
							}
						}
						log.Debug().Msgf("detect updated file %s (%s)", filename, event.Op)
						msgbus.PubCfgFileUpdate(bus, p.String(), msgbus.CfgFileUpdated{Path: p, Filename: filename})
					}
				}

			}
		}
	}()
	return stop, nil
}

func cfgFilenameToPath(filename string) (path.T, error) {
	return filenameToPath(filename, rawconfig.Paths.Etc, ".conf")
}

func frozenFilenameToPath(filename string) (path.T, error) {
	return filenameToPath(filename, rawconfig.Paths.Var, "")
}

func filenameToPath(filename, prefix, suffix string) (path.T, error) {
	svcName := strings.TrimPrefix(filename, prefix+"/")
	svcName = strings.TrimPrefix(svcName, "namespaces/")
	svcName = strings.TrimSuffix(svcName, suffix)
	if len(svcName) == 0 {
		return path.T{}, errors.New("skipped null filename")
	}
	return path.Parse(svcName)
}