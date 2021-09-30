package rescontainerdocker

import (
	"context"
	"time"

	"opensvc.com/opensvc/core/actionrollback"
	"opensvc.com/opensvc/core/drivergroup"
	"opensvc.com/opensvc/core/keywords"
	"opensvc.com/opensvc/core/manifest"
	"opensvc.com/opensvc/core/path"
	"opensvc.com/opensvc/core/provisioned"
	"opensvc.com/opensvc/core/resource"
	"opensvc.com/opensvc/core/status"
	"opensvc.com/opensvc/drivers/rescontainer"
	"opensvc.com/opensvc/util/converters"
)

const (
	driverGroup = drivergroup.Container
	driverName  = "docker"
)

type (
	T struct {
		resource.T
		Path               path.T         `json:"path"`
		SCSIReserv         bool           `json:"scsireserv"`
		PromoteRW          bool           `json:"promote_rw"`
		NoPreemptAbort     bool           `json:"NoPreemptAbort"`
		OsvcRootPath       string         `json:"osvc_root_path"`
		GuestOS            string         `json:"guest_os"`
		Name               string         `json:"name"`
		Hostname           string         `json:"hostname"`
		Image              string         `json:"image"`
		ImagePullPolicy    string         `json:"image_pull_policy"`
		Command            string         `json:"command"`
		RunArgs            string         `json:"run_args"`
		Entrypoint         string         `json:"entrypoint"`
		Detach             bool           `json:"detach"`
		Remove             bool           `json:"remove"`
		Privileged         bool           `json:"privileged"`
		Interactive        bool           `json:"interactive"`
		TTY                bool           `json:"tty"`
		VolumeMounts       []string       `json:"volume_mounts"`
		Environment        []string       `json:"environment"`
		SecretsEnvironment []string       `json:"secrets_environment"`
		ConfigsEnvironment []string       `json:"configs_environment"`
		Devices            []string       `json:"devices"`
		NetNS              string         `json:"netns"`
		UserNS             string         `json:"userns"`
		PIDNS              string         `json:"pidns"`
		IPCNS              string         `json:"ipcns"`
		UTSNS              string         `json:"utsns"`
		RegistryCreds      string         `json:"registry_creds"`
		PullTimeout        *time.Duration `json:"pull_timeout"`
		StartTimeout       *time.Duration `json:"start_timeout"`
		StopTimeout        *time.Duration `json:"stop_timeout"`
	}
)

func init() {
	resource.Register(driverGroup, driverName, New)
	resource.Register(driverGroup, "oci", New)
}

func New() resource.Driver {
	t := &T{}
	return t
}

// Manifest exposes to the core the input expected by the driver.
func (t T) Manifest() *manifest.T {
	m := manifest.New(driverGroup, driverName, t)
	m.AddContext([]manifest.Context{
		{
			Key:  "path",
			Attr: "Path",
			Ref:  "object.path",
		},
	}...)
	m.AddKeyword([]keywords.Keyword{
		{
			Option:      "name",
			Attr:        "Name",
			Scopable:    true,
			DefaultText: "<autogenerated>",
			Text:        "The name to assign to the container on docker run. If none is specified a ``<namespace>..<name>.container.<rid idx>`` name is automatically assigned.",
			Example:     "osvcprd..rundeck.container.db",
		},
		{
			Option:   "hostname",
			Attr:     "Hostname",
			Scopable: true,
			Example:  "nginx1",
			Text:     "Set the container hostname. If not set, a unique id is used.",
		},
		{
			Option:   "image",
			Attr:     "Image",
			Aliases:  []string{"run_image"},
			Scopable: true,
			Required: true,
			Example:  "google/pause",
			Text:     "The docker image pull, and run the container with.",
		},
		{
			Option:     "image_pull_policy",
			Attr:       "ImagePullPolicy",
			Scopable:   true,
			Candidates: []string{"once", "always"},
			Example:    "once",
			Text:       "The docker image pull policy. ``always`` pull upon each container start, ``once`` pull if not already pulled (default).",
		},
		{
			Option:   "command",
			Attr:     "Command",
			Aliases:  []string{"run_command"},
			Scopable: true,
			Example:  "/opt/tomcat/bin/catalina.sh",
			Text:     "The command to execute in the docker container on run.",
		},
		{
			Option:   "run_args",
			Attr:     "RunArgs",
			Scopable: true,
			Example:  "-v /opt/docker.opensvc.com/vol1:/vol1:rw -p 37.59.71.25:8080:8080",
			Text:     "Extra arguments to pass to the docker run command, like volume and port mappings.",
		},
		{
			Option:   "entrypoint",
			Attr:     "Entrypoint",
			Scopable: true,
			Example:  "/bin/sh",
			Text:     "The script or binary executed in the container. Args must be set in :kw:`command`.",
		},
		{
			Option:    "detach",
			Attr:      "Detach",
			Scopable:  true,
			Converter: converters.Bool,
			Default:   "true",
			Text:      "Run container in background. Set to ``false`` only for init containers, alongside :kw:`start_timeout` and the :c-tag:`nostatus` tag.",
		},
		{
			Option:    "remove",
			Attr:      "Remove",
			Scopable:  true,
			Converter: converters.Bool,
			Example:   "false",
			Text:      "If set to ``true``, add :opt:`--rm` to the docker run args and make sure the instance is removed on resource stop.",
		},
		{
			Option:    "privileged",
			Attr:      "Privileged",
			Scopable:  true,
			Converter: converters.Bool,
			Text:      "Give extended privileges to the container.",
		},
		{
			Option:    "interactive",
			Attr:      "Interactive",
			Scopable:  true,
			Converter: converters.Bool,
			Text:      "Keep stdin open even if not attached. To use if the container entrypoint is a shell.",
		},
		{
			Option:    "tty",
			Attr:      "TTY",
			Scopable:  true,
			Converter: converters.Bool,
			Text:      "Allocate a pseudo-tty.",
		},
		{
			Option:    "volume_mounts",
			Attr:      "VolumeMounts",
			Scopable:  true,
			Converter: converters.Shlex,
			Text:      "The whitespace separated list of ``<volume name|local dir>:<containerized mount path>:<mount options>``. When the source is a local dir, the default <mount option> is rw. When the source is a volume name, the default <mount option> is taken from volume access.",
			Example:   "myvol1:/vol1 myvol2:/vol2:rw /localdir:/data:ro",
		},
		{
			Option:    "environment",
			Attr:      "Environment",
			Scopable:  true,
			Converter: converters.Shlex,
			Text:      "A whitespace separated list of ``<var>=<secret name>/<key path>``. A shell expression spliter is applied, so double quotes can be around ``<secret name>/<key path>`` only or whole ``<var>=<secret name>/<key path>``. Variables are uppercased.",
			Example:   "KEY=cert1/server.key PASSWORD=db/password",
		},
		{
			Option:    "configs_environment",
			Attr:      "ConfigsEnvironment",
			Scopable:  true,
			Converter: converters.Shlex,
			Text:      "A whitespace separated list of ``<var>=<config name>/<key path>``. A shell expression spliter is applied, so double quotes can be around ``<config name>/<key path>`` only or whole ``<var>=<config name>/<key path>``. Variables are uppercased.",
			Example:   "CRT=cert1/server.crt PEM=cert1/server.pem",
		},
		{
			Option:    "devices",
			Attr:      "Devices",
			Scopable:  true,
			Converter: converters.Shlex,
			Text:      "The whitespace separated list of ``<host devpath>:<containerized devpath>``, specifying the host devices the container should have access to.",
			Example:   "myvol1:/dev/xvda myvol2:/dev/xvdb",
		},
		{
			Option:   "netns",
			Attr:     "NetNS",
			Aliases:  []string{"net"},
			Scopable: true,
			Example:  "container#0",
			Text:     "Sets the :cmd:`docker run --net` argument. The default is ``none`` if :opt:`--net` is not specified in :kw:`run_args`, meaning the container will have a private netns other containers can share. A :c-res:`ip.netns` or :c-res:`ip.cni` resource can configure an ip address in this container. A container with ``netns=container#0`` will share the container#0 netns. In this case agent format a :opt:`--net=container:<name of container#0 docker instance>`. ``netns=host`` shares the host netns.",
		},
		{
			Option:   "userns",
			Attr:     "UserNS",
			Scopable: true,
			Example:  "container#0",
			Text:     "Sets the :cmd:`docker run --userns` argument. If not set, the container will have a private userns other containers can share. A container with ``userns=host`` will share the host's userns.",
		},
		{
			Option:   "pidns",
			Attr:     "PIDNS",
			Scopable: true,
			Example:  "container#0",
			Text:     "Sets the :cmd:`docker run --pid` argument. If not set, the container will have a private pidns other containers can share. Usually a pidns sharer will run a google/pause image to reap zombies. A container with ``pidns=container#0`` will share the container#0 pidns. In this case agent format a :opt:`--pid=container:<name of container#0 docker instance>`. Use ``pidns=host`` to share the host's pidns.",
		},
		{
			Option:   "ipcns",
			Attr:     "IPCNS",
			Scopable: true,
			Example:  "container#0",
			Text:     "Sets the :cmd:`docker run --ipc` argument. If not set, the docker daemon's default value is used. ``ipcns=none`` does not mount /dev/shm. ``ipcns=private`` creates a ipcns other containers can not share. ``ipcns=shareable`` creates a netns other containers can share. ``ipcns=container#0`` will share the container#0 ipcns.",
		},
		{
			Option:     "utsns",
			Attr:       "UTSNS",
			Scopable:   true,
			Candidates: []string{"", "host"},
			Example:    "container#0",
			Text:       "Sets the :cmd:`docker run --uts` argument. If not set, the container will have a private utsns. A container with ``utsns=host`` will share the host's hostname.",
		},
		{
			Option:   "registry_creds",
			Attr:     "RegistryCreds",
			Scopable: true,
			Example:  "creds-registry-opensvc-com",
			Text:     "The name of a secret in the same namespace having a config.json key which value is used to login to the container image registry. If not specified, the node-level registry credential store is used.",
		},
		{
			Option:    "pull_timeout",
			Attr:      "PullTimeout",
			Scopable:  true,
			Converter: converters.Duration,
			Text:      "Wait for <duration> before declaring the container action a failure.",
			Example:   "2m",
			Default:   "2m",
		},
		{
			Option:    "start_timeout",
			Attr:      "StartTimeout",
			Scopable:  true,
			Converter: converters.Duration,
			Text:      "Wait for <duration> before declaring the container action a failure.",
			Example:   "5s",
			Default:   "1m5s",
		},
		{
			Option:    "stop_timeout",
			Attr:      "StopTimeout",
			Scopable:  true,
			Converter: converters.Duration,
			Text:      "Wait for <duration> before declaring the container action a failure.",
			Example:   "2m",
			Default:   "2m30s",
		},
		rescontainer.KWSCSIReserv,
		rescontainer.KWPromoteRW,
		rescontainer.KWNoPreemptAbort,
		rescontainer.KWOsvcRootPath,
		rescontainer.KWGuestOS,
	}...)
	return m
}

func (t T) Start(ctx context.Context) error {
	return nil
}

func (t T) Stop(ctx context.Context) error {
	return nil
}

func (t *T) Status(ctx context.Context) status.T {
	return status.NotApplicable
}

func (t T) Label() string {
	return t.Image
}

func (t T) Provision(ctx context.Context) error {
	return nil
}

func (t T) Unprovision(ctx context.Context) error {
	return nil
}

func (t T) Provisioned() (provisioned.T, error) {
	return provisioned.NotApplicable, nil
}

func (t T) create(ctx context.Context) error {
	actionrollback.Register(ctx, func() error {
		t.Log().Info().Msgf("rollback")
		return nil
	})
	return nil
}

// ContainerName formats a docker container name
func (t T) ContainerName() string {
	if t.Name != "" {
		return t.Name
	}
	s := ""
	if t.Path.Namespace != "" {
		s = t.Path.Namespace + ".."
	}
	s = s + t.Path.Name + "." + t.ResourceID.String()
	return s
}
