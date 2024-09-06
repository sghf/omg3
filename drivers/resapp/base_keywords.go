package resapp

import (
	"github.com/opensvc/om3/core/keywords"
	"github.com/opensvc/om3/util/converters"
)

var (
	BaseKeywordTimeout = keywords.Keyword{
		Attr:      "Timeout",
		Converter: converters.Duration,
		Example:   "180",
		Option:    "timeout",
		Scopable:  true,
		Text:      keywords.NewText(fs, "text/kw/timeout"),
	}
	BaseKeywordStopTimeout = keywords.Keyword{
		Attr:      "StopTimeout",
		Converter: converters.Duration,
		Example:   "180",
		Option:    "stop_timeout",
		Scopable:  true,
		Text:      keywords.NewText(fs, "text/kw/stop_timeout"),
	}
	BaseKeywordSecretsEnv = keywords.Keyword{
		Attr:      "SecretsEnv",
		Converter: converters.Shlex,
		Example:   "CRT=cert1/server.pem sec1/*",
		Option:    "secrets_environment",
		Scopable:  true,
		Text:      keywords.NewText(fs, "text/kw/secrets_environment"),
	}
	BaseKeywordConfigsEnv = keywords.Keyword{
		Attr:      "ConfigsEnv",
		Converter: converters.Shlex,
		Example:   "PORT=http/port webapp/app1* {name}/* {name}-debug/settings",
		Option:    "configs_environment",
		Scopable:  true,
		Text:      keywords.NewText(fs, "text/kw/configs_environment"),
	}
	BaseKeywordEnv = keywords.Keyword{
		Attr:      "Env",
		Example:   "CRT=cert1/server.crt PEM=cert1/server.pem",
		Option:    "environment",
		Converter: converters.Shlex,
		Scopable:  true,
		Text:      keywords.NewText(fs, "text/kw/environment"),
	}
	BaseKeywordRetCodes = keywords.Keyword{
		Attr:     "RetCodes",
		Default:  "0:up 1:down",
		Example:  "0:up 1:down 3:warn 4: n/a 5:undef",
		Option:   "retcodes",
		Required: false,
		Scopable: true,
		Text:     keywords.NewText(fs, "text/kw/retcodes"),
	}
	BaseKeywordUmask = keywords.Keyword{
		Attr:      "Umask",
		Converter: converters.Umask,
		Example:   "022",
		Option:    "umask",
		Scopable:  true,
		Text:      keywords.NewText(fs, "text/kw/umask"),
	}

	BaseKeywords = []keywords.Keyword{
		BaseKeywordTimeout,
		BaseKeywordStopTimeout,
		BaseKeywordSecretsEnv,
		BaseKeywordConfigsEnv,
		BaseKeywordEnv,
		BaseKeywordRetCodes,
		BaseKeywordUmask,
	}
)
