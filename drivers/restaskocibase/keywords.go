package restaskocibase

import (
	"embed"

	"github.com/opensvc/om3/core/keywords"
	"github.com/opensvc/om3/util/converters"
)

var (
	//go:embed text
	fs embed.FS

	Keywords = []keywords.Keyword{
		{
			Attr:        "Name",
			DefaultText: keywords.NewText(fs, "text/kw/name.default"),
			Example:     "osvcprd..rundeck.container.db",
			Option:      "name",
			Scopable:    true,
			Text:        keywords.NewText(fs, "text/kw/name"),
		},
		{
			Attr:     "Hostname",
			Example:  "nginx1",
			Option:   "hostname",
			Scopable: true,
			Text:     keywords.NewText(fs, "text/kw/hostname"),
		},
		{
			Aliases:   []string{},
			Attr:      "DNSSearch",
			Converter: converters.List,
			Example:   "opensvc.com",
			Option:    "dns_search",
			Required:  false,
			Scopable:  true,
			Text:      keywords.NewText(fs, "text/kw/dns_search"),
		},
		{
			Aliases:  []string{"run_image"},
			Attr:     "Image",
			Example:  "ghcr.io/opensvc/pause",
			Option:   "image",
			Required: true,
			Scopable: true,
			Text:     keywords.NewText(fs, "text/kw/image"),
		},
		{
			Attr:       "ImagePullPolicy",
			Candidates: []string{"once", "always"},
			Example:    "once",
			Option:     "image_pull_policy",
			Scopable:   true,
			Text:       keywords.NewText(fs, "text/kw/image_pull_policy"),
		},
		{
			Attr:     "CWD",
			Option:   "cwd",
			Example:  "/opt/foo",
			Scopable: true,
			Text:     keywords.NewText(fs, "text/kw/cwd"),
		},
		{
			Aliases:   []string{"run_command"},
			Attr:      "Command",
			Converter: converters.Shlex,
			Example:   "/opt/tomcat/bin/catalina.sh",
			Option:    "command",
			Scopable:  true,
			Text:      keywords.NewText(fs, "text/kw/command"),
		},
		{
			Attr:      "RunArgs",
			Converter: converters.Shlex,
			Example:   "-v /opt/docker.opensvc.com/vol1:/vol1:rw -p 37.59.71.25:8080:8080",
			Option:    "run_args",
			Scopable:  true,
			Text:      keywords.NewText(fs, "text/kw/run_args"),
		},
		{
			Attr:      "Entrypoint",
			Converter: converters.Shlex,
			Example:   "/bin/sh",
			Option:    "entrypoint",
			Scopable:  true,
			Text:      keywords.NewText(fs, "text/kw/entrypoint"),
		},
		{
			Attr:      "Detach",
			Converter: converters.Bool,
			Default:   "true",
			Option:    "detach",
			Scopable:  true,
			Text:      keywords.NewText(fs, "text/kw/detach"),
		},
		{
			Attr:      "Remove",
			Converter: converters.Bool,
			Example:   "false",
			Option:    "rm",
			Scopable:  true,
			Text:      keywords.NewText(fs, "text/kw/rm"),
		},
		{
			Attr:      "Privileged",
			Converter: converters.Bool,
			Option:    "privileged",
			Scopable:  true,
			Text:      keywords.NewText(fs, "text/kw/privileged"),
		},
		{
			Attr:      "Init",
			Converter: converters.Bool,
			Default:   "true",
			Option:    "init",
			Scopable:  true,
			Text:      keywords.NewText(fs, "text/kw/init"),
		},
		{
			Attr:      "Interactive",
			Converter: converters.Bool,
			Option:    "interactive",
			Scopable:  true,
			Text:      keywords.NewText(fs, "text/kw/interactive"),
		},
		{
			Attr:      "TTY",
			Converter: converters.Bool,
			Option:    "tty",
			Scopable:  true,
			Text:      keywords.NewText(fs, "text/kw/tty"),
		},
		{
			Attr:      "VolumeMounts",
			Converter: converters.Shlex,
			Example:   "myvol1:/vol1 myvol2:/vol2:rw /localdir:/data:ro",
			Option:    "volume_mounts",
			Scopable:  true,
			Text:      keywords.NewText(fs, "text/kw/volume_mounts"),
		},
		{
			Attr:      "Env",
			Converter: converters.Shlex,
			Example:   "KEY=cert1/server.key PASSWORD=db/password",
			Option:    "environment",
			Scopable:  true,
			Text:      keywords.NewText(fs, "text/kw/environment"),
		},
		{
			Attr:      "ConfigsEnv",
			Converter: converters.Shlex,
			Example:   "CRT=cert1/server.crt PEM=cert1/server.pem",
			Option:    "configs_environment",
			Scopable:  true,
			Text:      keywords.NewText(fs, "text/kw/configs_environment"),
		},
		{
			Attr:      "Devices",
			Converter: converters.Shlex,
			Example:   "myvol1:/dev/xvda myvol2:/dev/xvdb",
			Option:    "devices",
			Scopable:  true,
			Text:      keywords.NewText(fs, "text/kw/devices"),
		},
		{
			Aliases:  []string{"net"},
			Attr:     "NetNS",
			Example:  "container#0",
			Option:   "netns",
			Scopable: true,
			Text:     keywords.NewText(fs, "text/kw/netns"),
		},
		{
			Attr:     "User",
			Example:  "guest",
			Option:   "user",
			Scopable: true,
			Text:     keywords.NewText(fs, "text/kw/user"),
		},
		{
			Attr:     "PIDNS",
			Example:  "container#0",
			Option:   "pidns",
			Scopable: true,
			Text:     keywords.NewText(fs, "text/kw/pidns"),
		},
		{
			Attr:     "IPCNS",
			Example:  "container#0",
			Option:   "ipcns",
			Scopable: true,
			Text:     keywords.NewText(fs, "text/kw/ipcns"),
		},
		{
			Attr:       "UTSNS",
			Candidates: []string{"", "host"},
			Example:    "container#0",
			Option:     "utsns",
			Scopable:   true,
			Text:       keywords.NewText(fs, "text/kw/utsns"),
		},
		{
			Attr:     "RegistryCreds",
			Example:  "creds-registry-opensvc-com",
			Option:   "registry_creds",
			Scopable: true,
			Text:     keywords.NewText(fs, "text/kw/registry_creds"),
		},
		{
			Attr:      "PullTimeout",
			Converter: converters.Duration,
			Default:   "2m",
			Example:   "2m",
			Option:    "pull_timeout",
			Scopable:  true,
			Text:      keywords.NewText(fs, "text/kw/pull_timeout"),
		},
		{
			Attr:      "SecretsEnv",
			Converter: converters.Shlex,
			Example:   "CRT=cert1/server.pem sec1/*",
			Option:    "secrets_environment",
			Scopable:  true,
			Text:      keywords.NewText(fs, "text/kw/secrets_environment"),
		},
		{
			Attr:      "ConfigsEnv",
			Converter: converters.Shlex,
			Example:   "PORT=http/port webapp/app1* {name}/* {name}-debug/settings",
			Option:    "configs_environment",
			Scopable:  true,
			Text:      keywords.NewText(fs, "text/kw/configs_environment"),
		},
	}
)
