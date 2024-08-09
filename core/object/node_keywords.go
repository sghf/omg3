package object

import (
	"fmt"

	"github.com/opensvc/om3/core/keywords"
	"github.com/opensvc/om3/core/naming"
	"github.com/opensvc/om3/core/rawconfig"
	"github.com/opensvc/om3/daemon/daemonenv"
	"github.com/opensvc/om3/util/converters"
	"github.com/opensvc/om3/util/key"
)

const (
	DefaultNodeMaxParallel = 10
)

var nodePrivateKeywords = []keywords.Keyword{
	{
		Section: "node",
		Option:  "oci",
		Text:    keywords.NewText(fs, "text/kw/node/node.oci"),
	},
	{
		Section: "node",
		Option:  "uuid",
		Text:    keywords.NewText(fs, "text/kw/node/node.uuid"),
	},
	{
		Section:     "node",
		Option:      "prkey",
		DefaultText: keywords.NewText(fs, "text/kw/node/node.prkey.default"),
		Text:        keywords.NewText(fs, "text/kw/node/node.prkey"),
	},
	{
		Section: "node",
		Option:  "connect_to",
		Example: "1.2.3.4",
		Text:    keywords.NewText(fs, "text/kw/node/node.connect_to"),
	},
	{
		Section:   "node",
		Option:    "mem_bytes",
		Example:   "256mb",
		Converter: converters.Size,
		Text:      keywords.NewText(fs, "text/kw/node/node.mem_bytes"),
	},
	{
		Section:   "node",
		Option:    "mem_banks",
		Example:   "4",
		Converter: converters.Int,
		Text:      keywords.NewText(fs, "text/kw/node/node.mem_banks"),
	},
	{
		Section:   "node",
		Option:    "mem_slots",
		Example:   "4",
		Converter: converters.Int,
		Text:      keywords.NewText(fs, "text/kw/node/node.mem_slots"),
	},
	{
		Section: "node",
		Option:  "os_vendor",
		Example: "Digital",
		Text:    keywords.NewText(fs, "text/kw/node/node.os_vendor"),
	},
	{
		Section: "node",
		Option:  "os_release",
		Example: "5",
		Text:    keywords.NewText(fs, "text/kw/node/node.os_release"),
	},
	{
		Section: "node",
		Option:  "os_kernel",
		Example: "5.1234",
		Text:    keywords.NewText(fs, "text/kw/node/node.os_kernel"),
	},
	{
		Section: "node",
		Option:  "os_arch",
		Example: "5.1234",
		Text:    keywords.NewText(fs, "text/kw/node/node.os_arch"),
	},
	{
		Section: "node",
		Option:  "cpu_freq",
		Example: "3.2 Ghz",
		Text:    keywords.NewText(fs, "text/kw/node/node.cpu_freq"),
	},
	{
		Section:   "node",
		Option:    "cpu_threads",
		Example:   "4",
		Converter: converters.Int,
		Text:      keywords.NewText(fs, "text/kw/node/node.cpu_threads"),
	},
	{
		Section:   "node",
		Option:    "cpu_cores",
		Example:   "2",
		Converter: converters.Int,
		Text:      keywords.NewText(fs, "text/kw/node/node.cpu_cores"),
	},
	{
		Section:   "node",
		Option:    "cpu_dies",
		Example:   "1",
		Converter: converters.Int,
		Text:      keywords.NewText(fs, "text/kw/node/node.cpu_dies"),
	},
	{
		Section: "node",
		Option:  "cpu_model",
		Example: "Alpha EV5",
		Text:    keywords.NewText(fs, "text/kw/node/node.cpu_model"),
	},
	{
		Section: "node",
		Option:  "serial",
		Example: "abcdef0123456",
		Text:    keywords.NewText(fs, "text/kw/node/node.serial"),
	},
	{
		Section: "node",
		Option:  "bios_version",
		Example: "1.025",
		Text:    keywords.NewText(fs, "text/kw/node/node.bios_version"),
	},
	{
		Section: "node",
		Option:  "sp_version",
		Example: "1.026",
		Text:    keywords.NewText(fs, "text/kw/node/node.sp_version"),
	},
	{
		Section: "node",
		Option:  "enclosure",
		Example: "1",
		Text:    keywords.NewText(fs, "text/kw/node/node.enclosure"),
	},
	{
		Section: "node",
		Option:  "tz",
		Example: "+0200",
		Text:    keywords.NewText(fs, "text/kw/node/node.tz"),
	},
	{
		Section: "node",
		Option:  "manufacturer",
		Example: "Digital",
		Text:    keywords.NewText(fs, "text/kw/node/node.manufacturer"),
	},
	{
		Section: "node",
		Option:  "model",
		Example: "ds20e",
		Text:    keywords.NewText(fs, "text/kw/node/node.model"),
	},
	{
		Section: "array",
		Option:  "schedule",
		Text:    keywords.NewText(fs, "text/kw/node/array.schedule"),
	},
	{
		Section: "array",
		Types:   []string{"xtremio"},
		Option:  "name",
		Example: "array1",
		Text:    keywords.NewText(fs, "text/kw/node/array.xtremio.name"),
	},
	{
		Section: "backup",
		Option:  "schedule",
		Text:    keywords.NewText(fs, "text/kw/node/backup.schedule"),
	},
	{
		Section: "switch",
		Option:  "schedule",
		Text:    keywords.NewText(fs, "text/kw/node/switch.schedule"),
	},
}

var nodeCommonKeywords = []keywords.Keyword{
	{
		Section:   "node",
		Option:    "secure_fetch",
		Default:   "true",
		Converter: converters.Bool,
		Text:      keywords.NewText(fs, "text/kw/node/node.secure_fetch"),
	},
	{
		Section:   "node",
		Option:    "min_avail_mem",
		Default:   "2%",
		Converter: converters.Size,
		Text:      keywords.NewText(fs, "text/kw/node/node.min_avail_mem"),
	},
	{
		Section:   "node",
		Option:    "min_avail_swap",
		Default:   "10%",
		Converter: converters.Size,
		Text:      keywords.NewText(fs, "text/kw/node/node.min_avail_swap"),
	},
	{
		Section:    "node",
		Option:     "env",
		Default:    "TST",
		Candidates: rawconfig.Envs,
		Text:       keywords.NewText(fs, "text/kw/node/node.env"),
	},
	{
		Section:   "node",
		Option:    "max_parallel",
		Default:   fmt.Sprintf("%d", DefaultNodeMaxParallel),
		Converter: converters.Int,
		Text:      keywords.NewText(fs, "text/kw/node/node.max_parallel"),
	},
	{
		Section:   "node",
		Option:    "allowed_networks",
		Default:   "10.0.0.0/8 172.16.0.0/24 192.168.0.0/16",
		Converter: converters.List,
		Text:      keywords.NewText(fs, "text/kw/node/node.allowed_networks"),
	},
	{
		Section: "node",
		Option:  "loc_country",
		Example: "fr",
		Text:    keywords.NewText(fs, "text/kw/node/node.loc_country"),
	},
	{
		Section: "node",
		Option:  "loc_city",
		Example: "Paris",
		Text:    keywords.NewText(fs, "text/kw/node/node.loc_city"),
	},
	{
		Section: "node",
		Option:  "loc_zip",
		Example: "75017",
		Text:    keywords.NewText(fs, "text/kw/node/node.loc_zip"),
	},
	{
		Section: "node",
		Option:  "loc_addr",
		Example: "7 rue blanche",
		Text:    keywords.NewText(fs, "text/kw/node/node.loc_addr"),
	},
	{
		Section: "node",
		Option:  "loc_building",
		Example: "Crystal",
		Text:    keywords.NewText(fs, "text/kw/node/node.loc_building"),
	},
	{
		Section: "node",
		Option:  "loc_floor",
		Example: "21",
		Text:    keywords.NewText(fs, "text/kw/node/node.loc_floor"),
	},
	{
		Section: "node",
		Option:  "loc_room",
		Example: "102",
		Text:    keywords.NewText(fs, "text/kw/node/node.loc_room"),
	},
	{
		Section: "node",
		Option:  "loc_rack",
		Example: "R42",
		Text:    keywords.NewText(fs, "text/kw/node/node.loc_rack"),
	},
	{
		Section: "node",
		Option:  "sec_zone",
		Example: "dmz1",
		Text:    keywords.NewText(fs, "text/kw/node/node.sec_zone"),
	},
	{
		Section: "node",
		Option:  "team_integ",
		Example: "TINT",
		Text:    keywords.NewText(fs, "text/kw/node/node.team_integ"),
	},
	{
		Section: "node",
		Option:  "team_support",
		Example: "TSUP",
		Text:    keywords.NewText(fs, "text/kw/node/node.team_support"),
	},
	{
		Section: "node",
		Option:  "asset_env",
		Example: "Production",
		Text:    keywords.NewText(fs, "text/kw/node/node.asset_env"),
	},
	{
		Section: "node",
		Option:  "dbopensvc",
		Example: "https://collector.opensvc.com",
		Text:    keywords.NewText(fs, "text/kw/node/node.dbopensvc"),
	},
	{
		Section:   "node",
		Option:    "dbinsecure",
		Converter: converters.Bool,
		Text:      keywords.NewText(fs, "text/kw/node/node.dbinsecure"),
	},
	{
		Section:     "node",
		Option:      "dbcompliance",
		Example:     "https://collector.opensvc.com",
		DefaultText: keywords.NewText(fs, "text/kw/node/node.dbcompliance.default"),
		Text:        keywords.NewText(fs, "text/kw/node/node.dbcompliance"),
	},
	{
		Section:   "node",
		Option:    "dblog",
		Converter: converters.Bool,
		Default:   "true",
		Text:      keywords.NewText(fs, "text/kw/node/node.dblog"),
	},
	{
		Section: "node",
		Option:  "branch",
		Example: "1.9",
		Text:    keywords.NewText(fs, "text/kw/node/node.branch"),
	},
	{
		Section: "node",
		Option:  "repo",
		Example: "http://opensvc.repo.corp",
		Text:    keywords.NewText(fs, "text/kw/node/node.repo"),
	},
	{
		Section: "node",
		Option:  "repopkg",
		Example: "http://repo.opensvc.com",
		Text:    keywords.NewText(fs, "text/kw/node/node.repopkg"),
	},
	{
		Section: "node",
		Option:  "repocomp",
		Example: "http://compliance.repo.corp",
		Text:    keywords.NewText(fs, "text/kw/node/node.repocomp"),
	},
	{
		Section: "node",
		Option:  "ruser",
		Default: "root",
		Example: "root opensvc@node1",
		Text:    keywords.NewText(fs, "text/kw/node/node.ruser"),
	},
	{
		Section:   "node",
		Option:    "maintenance_grace_period",
		Converter: converters.Duration,
		Default:   "60",
		Text:      keywords.NewText(fs, "text/kw/node/node.maintenance_grace_period"),
	},
	{
		Section:   "node",
		Option:    "rejoin_grace_period",
		Converter: converters.Duration,
		Default:   "90s",
		Text:      keywords.NewText(fs, "text/kw/node/node.rejoin_grace_period"),
	},
	{
		Section:   "node",
		Option:    "ready_period",
		Converter: converters.Duration,
		Default:   "5s",
		Text:      keywords.NewText(fs, "text/kw/node/node.ready_period"),
	},
	{
		Section: "dequeue_actions",
		Option:  "schedule",
		Text:    keywords.NewText(fs, "text/kw/node/dequeue_actions.schedule"),
	},
	{
		Section: "sysreport",
		Option:  "schedule",
		Default: "~00:00-06:00",
		Text:    keywords.NewText(fs, "text/kw/node/sysreport.schedule"),
	},
	{
		Section: "compliance",
		Option:  "schedule",
		Default: "02:00-06:00",
		Text:    keywords.NewText(fs, "text/kw/node/compliance.schedule"),
	},
	{
		Section:   "compliance",
		Option:    "auto_update",
		Converter: converters.Bool,
		Default:   "false",
		Text:      keywords.NewText(fs, "text/kw/node/compliance.auto_update"),
	},
	{
		Section: "checks",
		Option:  "schedule",
		Default: "~00:00-06:00",
		Text:    keywords.NewText(fs, "text/kw/node/checks.schedule"),
	},
	{
		Section: "packages",
		Option:  "schedule",
		Default: "~00:00-06:00",
		Text:    keywords.NewText(fs, "text/kw/node/packages.schedule"),
	},
	{
		Section: "patches",
		Option:  "schedule",
		Default: "~00:00-06:00",
		Text:    keywords.NewText(fs, "text/kw/node/patches.schedule"),
	},
	{
		Section: "asset",
		Option:  "schedule",
		Default: "~00:00-06:00",
		Text:    keywords.NewText(fs, "text/kw/node/asset.schedule"),
	},
	{
		Section: "disks",
		Option:  "schedule",
		Default: "~00:00-06:00",
		Text:    keywords.NewText(fs, "text/kw/node/disks.schedule"),
	},
	{
		Section: "listener",
		Option:  "crl",
		Example: "https://crl.opensvc.com",
		Default: rawconfig.Paths.CACRL,
		Text:    keywords.NewText(fs, "text/kw/node/listener.crl"),
	},
	{
		Section: "listener",
		Option:  "dns_sock_uid",
		Default: "953",
		Text:    keywords.NewText(fs, "text/kw/node/listener.dns_sock_uid"),
	},
	{
		Section: "listener",
		Option:  "dns_sock_gid",
		Default: "953",
		Text:    keywords.NewText(fs, "text/kw/node/listener.dns_sock_gid"),
	},
	{
		Section:  "listener",
		Option:   "addr",
		Aliases:  []string{"tls_addr"},
		Scopable: true,
		Default:  "",
		Example:  "1.2.3.4",
		Text:     keywords.NewText(fs, "text/kw/node/listener.addr"),
	},
	{
		Section:   "listener",
		Option:    "port",
		Aliases:   []string{"tls_port"},
		Scopable:  true,
		Converter: converters.Int,
		Default:   fmt.Sprintf("%d", daemonenv.HTTPPort),
		Text:      keywords.NewText(fs, "text/kw/node/listener.port"),
	},
	{
		Section: "listener",
		Option:  "openid_well_known",
		Example: "https://keycloak.opensvc.com/auth/realms/clusters/.well-known/openid-configuration",
		Text:    keywords.NewText(fs, "text/kw/node/listener.openid_well_known"),
	},
	{
		Section: "syslog",
		Option:  "facility",
		Default: "daemon",
		Text:    keywords.NewText(fs, "text/kw/node/syslog.facility"),
	},
	{
		Section:    "syslog",
		Option:     "level",
		Default:    "info",
		Candidates: []string{"critical", "error", "warning", "info", "debug"},
		Text:       keywords.NewText(fs, "text/kw/node/syslog.level"),
	},
	{
		Section:     "syslog",
		Option:      "host",
		DefaultText: keywords.NewText(fs, "text/kw/node/syslog.host.default"),
		Text:        keywords.NewText(fs, "text/kw/node/syslog.host"),
	},
	{
		Section: "syslog",
		Option:  "port",
		Default: "514",
		Text:    keywords.NewText(fs, "text/kw/node/syslog.port"),
	},
	{
		Section:  "cluster",
		Option:   "vip",
		Example:  "192.168.99.12/24@eth0",
		Scopable: true,
		Text:     keywords.NewText(fs, "text/kw/node/cluster.vip"),
	},
	{
		Section:   "cluster",
		Option:    "dns",
		Converter: converters.List,
		Scopable:  true,
		Text:      keywords.NewText(fs, "text/kw/node/cluster.dns"),
	},
	{
		Section:     "cluster",
		Option:      "ca",
		Converter:   converters.List,
		DefaultText: keywords.NewText(fs, "text/kw/node/cluster.ca.default"),
		Text:        keywords.NewText(fs, "text/kw/node/cluster.ca"),
	},
	{
		Section:     "cluster",
		Option:      "cert",
		DefaultText: keywords.NewText(fs, "text/kw/node/cluster.cert.default"),
		Text:        keywords.NewText(fs, "text/kw/node/cluster.cert"),
	},
	{
		Section:     "cluster",
		Option:      "id",
		Scopable:    true,
		DefaultText: keywords.NewText(fs, "text/kw/node/cluster.id.default"),
		Text:        keywords.NewText(fs, "text/kw/node/cluster.id"),
	},
	{
		Section: "cluster",
		Option:  "name",
		Text:    keywords.NewText(fs, "text/kw/node/cluster.name"),
	},
	{
		Section:     "cluster",
		Option:      "secret",
		Scopable:    true,
		DefaultText: keywords.NewText(fs, "text/kw/node/cluster.secret.default"),
		Text:        keywords.NewText(fs, "text/kw/node/cluster.secret"),
	},
	{
		Section:   "cluster",
		Option:    "nodes",
		Converter: converters.List,
		Text:      keywords.NewText(fs, "text/kw/node/cluster.nodes"),
	},
	{
		Section:   "cluster",
		Option:    "drpnodes",
		Converter: converters.List,
		Text:      keywords.NewText(fs, "text/kw/node/cluster.drpnodes"),
	},
	{
		Section:   "cluster",
		Option:    "quorum",
		Converter: converters.Bool,
		Default:   "false",
		Text:      keywords.NewText(fs, "text/kw/node/cluster.quorum"),
	},
	{
		Section:    "node",
		Option:     "split_action",
		Scopable:   true,
		Candidates: []string{"crash", "reboot", "disabled"},
		Default:    "crash",
		Text:       keywords.NewText(fs, "text/kw/node/node.split_action"),
	},
	{
		Section:  "arbitrator",
		Option:   "uri",
		Aliases:  []string{"name"},
		Required: true,
		Example:  "http://www.opensvc.com",
		Text:     keywords.NewText(fs, "text/kw/node/arbitrator.uri"),
	},
	{
		Section:   "arbitrator",
		Option:    "insecure",
		Converter: converters.Bool,
		Default:   "false",
		Text:      keywords.NewText(fs, "text/kw/node/arbitrator.insecure"),
	},
	{
		Section:   "stonith",
		Option:    "cmd",
		Scopable:  true,
		Converter: converters.Shlex,
		Required:  true,
		Example:   "/bin/true",
		Text:      keywords.NewText(fs, "text/kw/node/stonith.cmd"),
	},
	{
		Section:    "hb",
		Option:     "type",
		Candidates: []string{"unicast", "multicast", "disk", "relay"},
		Required:   true,
		Text:       keywords.NewText(fs, "text/kw/node/hb.type"),
	},
	{
		Section:     "hb",
		Option:      "addr",
		Types:       []string{"unicast"},
		Scopable:    true,
		Example:     "1.2.3.4",
		DefaultText: keywords.NewText(fs, "text/kw/node/hb.unicast.addr.default"),
		Text:        keywords.NewText(fs, "text/kw/node/hb.unicast.addr"),
	},
	{
		Section:     "hb",
		Option:      "intf",
		Types:       []string{"unicast"},
		Scopable:    true,
		Example:     "eth0",
		DefaultText: keywords.NewText(fs, "text/kw/node/hb.unicast.intf.default"),
		Text:        keywords.NewText(fs, "text/kw/node/hb.unicast.intf"),
	},
	{
		Section:   "hb",
		Option:    "port",
		Types:     []string{"unicast"},
		Converter: converters.Int,
		Scopable:  true,
		Default:   "10000",
		Text:      keywords.NewText(fs, "text/kw/node/hb.unicast.port"),
	},
	{
		Section:   "hb",
		Option:    "timeout",
		Converter: converters.Duration,
		Scopable:  true,
		Default:   "15s",
		Text:      keywords.NewText(fs, "text/kw/node/hb.timeout"),
	},
	{
		Section:   "hb",
		Option:    "interval",
		Converter: converters.Duration,
		Scopable:  true,
		Default:   "5s",
		Text:      keywords.NewText(fs, "text/kw/node/hb.interval"),
	},
	{
		Section:  "hb",
		Option:   "addr",
		Types:    []string{"multicast"},
		Scopable: true,
		Default:  "224.3.29.71",
		Text:     keywords.NewText(fs, "text/kw/node/hb.multicast.addr"),
	},
	{
		Section:     "hb",
		Option:      "intf",
		Types:       []string{"multicast"},
		Scopable:    true,
		Example:     "eth0",
		DefaultText: keywords.NewText(fs, "text/kw/node/hb.multicast.intf.default"),
		Text:        keywords.NewText(fs, "text/kw/node/hb.multicast.intf"),
	},
	{
		Section:   "hb",
		Option:    "port",
		Types:     []string{"multicast"},
		Converter: converters.Int,
		Scopable:  true,
		Default:   "10000",
		Text:      keywords.NewText(fs, "text/kw/node/hb.multicast.port"),
	},
	{
		Section:     "hb",
		Option:      "nodes",
		Types:       []string{"unicast"},
		Scopable:    true,
		Converter:   converters.List,
		DefaultText: keywords.NewText(fs, "text/kw/node/hb.unicast.nodes.default"),
		Text:        keywords.NewText(fs, "text/kw/node/hb.unicast.nodes"),
	},
	{
		Section:  "hb",
		Option:   "dev",
		Types:    []string{"disk"},
		Scopable: true,
		Required: true,
		Example:  "/dev/mapper/36589cfc000000e03957c51dabab8373a",
		Text:     keywords.NewText(fs, "text/kw/node/hb.disk.dev"),
	},
	{
		Section:   "hb",
		Option:    "insecure",
		Types:     []string{"relay"},
		Converter: converters.Bool,
		Default:   "false",
		Text:      keywords.NewText(fs, "text/kw/node/hb.relay.insecure"),
	},
	{
		Section:  "hb",
		Option:   "relay",
		Types:    []string{"relay"},
		Required: true,
		Example:  "relaynode1",
		Text:     keywords.NewText(fs, "text/kw/node/hb.relay.relay"),
	},
	{
		Section: "hb",
		Option:  "username",
		Types:   []string{"relay"},
		Default: "relay",
		Text:    keywords.NewText(fs, "text/kw/node/hb.relay.username"),
	},
	{
		Section: "hb",
		Option:  "password",
		Types:   []string{"relay"},
		Default: "system/sec/relay",
		Text:    keywords.NewText(fs, "text/kw/node/hb.relay.password"),
	},
	{
		Section: "cni",
		Option:  "plugins",
		Default: "/opt/cni/bin",
		Text:    keywords.NewText(fs, "text/kw/node/cni.plugins"),
		Example: "/var/lib/opensvc/cni/bin",
	},
	{
		Section: "cni",
		Option:  "config",
		Default: "/opt/cni/net.d",
		Text:    keywords.NewText(fs, "text/kw/node/cni.config"),
		Example: "/var/lib/opensvc/cni/net.d",
	},
	{
		Section:    "pool",
		Option:     "type",
		Default:    "directory",
		Candidates: []string{"directory", "loop", "vg", "zpool", "freenas", "share", "shm", "symmetrix", "virtual", "dorado", "hoc", "drbd", "pure"},
		Text:       keywords.NewText(fs, "text/kw/node/pool.type"),
	},
	{
		Section: "pool",
		Option:  "status_schedule",
		Text:    keywords.NewText(fs, "text/kw/node/pool.status_schedule"),
	},
	{
		Section:  "pool",
		Option:   "mnt_opt",
		Scopable: true,
		Text:     keywords.NewText(fs, "text/kw/node/pool.mnt_opt"),
	},
	{
		Section:  "pool",
		Types:    []string{"freenas", "symmetrix", "dorado", "hoc", "pure"},
		Option:   "array",
		Scopable: true,
		Required: true,
		Text:     keywords.NewText(fs, "text/kw/node/pool.array"),
	},
	{
		Section: "pool",
		Types:   []string{"hoc", "pure"},
		Option:  "label_prefix",
		Text:    keywords.NewText(fs, "text/kw/node/pool.hoc.label_prefix"),
	},
	{
		Section:   "pool",
		Types:     []string{"pure"},
		Option:    "delete_now",
		Converter: converters.Bool,
		Default:   "true",
		Text:      keywords.NewText(fs, "text/kw/node/delete_now.pure.pod"),
	},
	{
		Section: "pool",
		Types:   []string{"pure"},
		Option:  "pod",
		Text:    keywords.NewText(fs, "text/kw/node/pool.pure.pod"),
	},
	{
		Section: "pool",
		Types:   []string{"pure"},
		Option:  "volumegroup",
		Text:    keywords.NewText(fs, "text/kw/node/pool.pure.volumegroup"),
	},
	{
		Section: "array",
		Types:   []string{"hoc"},
		Option:  "wwid_prefix",
		Text:    keywords.NewText(fs, "text/kw/node/array.hoc.wwid_prefix"),
	},
	{
		Section: "pool",
		Types:   []string{"hoc"},
		Option:  "volume_id_range_from",
		Text:    keywords.NewText(fs, "text/kw/node/pool.hoc.volume_id_range_from"),
	},
	{
		Section: "pool",
		Types:   []string{"hoc"},
		Option:  "volume_id_range_to",
		Text:    keywords.NewText(fs, "text/kw/node/pool.hoc.volume_id_range_to"),
	},
	{
		Section: "pool",
		Types:   []string{"hoc"},
		Option:  "vsm_id",
		Default: "",
		Text:    keywords.NewText(fs, "text/kw/node/pool.hoc.vsm_id"),
	},
	{
		Section:  "pool",
		Types:    []string{"symmetrix"},
		Option:   "srp",
		Required: true,
		Text:     keywords.NewText(fs, "text/kw/node/pool.symmetrix.srp"),
	},
	{
		Section: "pool",
		Types:   []string{"symmetrix"},
		Option:  "slo",
		Text:    keywords.NewText(fs, "text/kw/node/pool.symmetrix.slo"),
	},
	{
		Section:   "pool",
		Types:     []string{"symmetrix"},
		Option:    "srdf",
		Converter: converters.Bool,
		Default:   "false",
		Text:      keywords.NewText(fs, "text/kw/node/pool.symmetrix.srdf"),
	},
	{
		Section: "pool",
		Types:   []string{"symmetrix"},
		Option:  "rdfg",
		Text:    keywords.NewText(fs, "text/kw/node/pool.symmetrix.rdfg"),
	},
	{
		Section:  "pool",
		Types:    []string{"freenas", "dorado", "hoc", "pure"},
		Option:   "diskgroup",
		Required: true,
		Text:     keywords.NewText(fs, "text/kw/node/pool.diskgroup"),
	},
	{
		Section:   "pool",
		Types:     []string{"freenas"},
		Option:    "insecure_tpc",
		Converter: converters.Bool,
		Default:   "false",
		Text:      keywords.NewText(fs, "text/kw/node/pool.freenas.insecure_tpc"),
	},
	{
		Section:    "pool",
		Types:      []string{"freenas"},
		Option:     "compression",
		Default:    "inherit",
		Candidates: []string{"inherit", "none", "lz4", "gzip-1", "gzip-2", "gzip-3", "gzip-4", "gzip-5", "gzip-6", "gzip-7", "gzip-8", "gzip-9", "zle", "lzjb"},
		Text:       keywords.NewText(fs, "text/kw/node/pool.freenas.compression"),
	},
	{
		Section:   "pool",
		Types:     []string{"freenas"},
		Option:    "sparse",
		Converter: converters.Bool,
		Default:   "false",
		Text:      keywords.NewText(fs, "text/kw/node/pool.freenas.sparse"),
	},
	{
		Section:   "pool",
		Types:     []string{"freenas"},
		Option:    "blocksize",
		Default:   "512",
		Converter: converters.Size,
		Text:      keywords.NewText(fs, "text/kw/node/pool.freenas.blocksize"),
	},
	{
		Section:  "pool",
		Types:    []string{"vg"},
		Option:   "name",
		Required: true,
		Text:     keywords.NewText(fs, "text/kw/node/pool.vg.name"),
	},
	{
		Section: "pool",
		Types:   []string{"drbd"},
		Option:  "vg",
		Text:    keywords.NewText(fs, "text/kw/node/pool.drbd.vg"),
	},
	{
		Section:  "pool",
		Types:    []string{"zpool"},
		Option:   "name",
		Required: true,
		Text:     keywords.NewText(fs, "text/kw/node/pool.zpool.name"),
	},
	{
		Section: "pool",
		Types:   []string{"drbd"},
		Option:  "zpool",
		Text:    keywords.NewText(fs, "text/kw/node/pool.drbd.zpool"),
	},
	{
		Section: "pool",
		Option:  "path",
		Types:   []string{"drbd"},
		Text:    keywords.NewText(fs, "text/kw/node/pool.drbd.path"),
	},
	{
		Section: "pool",
		Option:  "path",
		Types:   []string{"share"},
		Default: "{var}/pool/share",
		Text:    keywords.NewText(fs, "text/kw/node/pool.share.path"),
	},
	{
		Section: "pool",
		Option:  "path",
		Types:   []string{"directory"},
		Default: "{var}/pool/directory",
		Text:    keywords.NewText(fs, "text/kw/node/pool.directory.path"),
	},
	{
		Section: "pool",
		Option:  "template",
		Types:   []string{"virtual"},
		Text:    keywords.NewText(fs, "text/kw/node/pool.virtual.template"),
		Example: "templates/vol/mpool-over-loop",
	},
	{
		Section:   "pool",
		Option:    "volume_env",
		Types:     []string{"virtual"},
		Converter: converters.List,
		Example:   "container#1.name:container_name env.foo:foo",
		Text:      keywords.NewText(fs, "text/kw/node/pool.virtual.volume_env"),
	},
	{
		Section:   "pool",
		Option:    "optional_volume_env",
		Types:     []string{"virtual"},
		Converter: converters.List,
		Example:   "container#1.name:container_name env.foo:foo",
		Text:      keywords.NewText(fs, "text/kw/node/pool.virtual.optional_volume_env"),
	},
	{
		Section:   "pool",
		Option:    "capabilities",
		Types:     []string{"virtual"},
		Converter: converters.List,
		Default:   "roo rwo rox rwx",
		Text:      keywords.NewText(fs, "text/kw/node/pool.virtual.capabilities"),
	},
	{
		Section: "pool",
		Option:  "path",
		Types:   []string{"loop"},
		Default: "{var}/pool/loop",
		Text:    keywords.NewText(fs, "text/kw/node/pool.loop.path"),
	},
	{
		Section: "pool",
		Option:  "pool_id",
		Types:   []string{"hoc"},
		Default: "",
		Text:    keywords.NewText(fs, "text/kw/node/pool.hoc.pool_id"),
	},
	{
		Section: "pool",
		Option:  "fs_type",
		Types:   []string{"freenas", "dorado", "hoc", "symmetrix", "drbd", "loop", "vg", "pure"},
		Default: "xfs",
		Text:    keywords.NewText(fs, "text/kw/node/pool.fs_type"),
	},
	{
		Section: "pool",
		Option:  "mkfs_opt",
		Example: "-O largefile",
		Text:    keywords.NewText(fs, "text/kw/node/pool.mkfs_opt"),
	},
	{
		Section: "pool",
		Option:  "mkblk_opt",
		Text:    keywords.NewText(fs, "text/kw/node/pool.mkblk_opt"),
	},
	{
		Section:   "hook",
		Option:    "events",
		Converter: converters.List,
		Text:      keywords.NewText(fs, "text/kw/node/hook.events"),
	},
	{
		Section:   "hook",
		Option:    "command",
		Converter: converters.Shlex,
		Text:      keywords.NewText(fs, "text/kw/node/hook.command"),
	},
	{
		Section:    "network",
		Option:     "type",
		Candidates: []string{"bridge", "routed_bridge"},
		Default:    "bridge",
		Text:       keywords.NewText(fs, "text/kw/node/network.type"),
	},
	{
		Section:  "network",
		Types:    []string{"routed_bridge"},
		Option:   "subnet",
		Scopable: true,
		Text:     keywords.NewText(fs, "text/kw/node/network.routed_bridge.subnet"),
	},
	{
		Section:  "network",
		Types:    []string{"routed_bridge"},
		Option:   "gateway",
		Scopable: true,
		Text:     keywords.NewText(fs, "text/kw/node/network.routed_bridge.gateway"),
	},
	{
		Section:   "network",
		Types:     []string{"routed_bridge"},
		Option:    "ips_per_node",
		Converter: converters.Int,
		Default:   "1024",
		Text:      keywords.NewText(fs, "text/kw/node/network.routed_bridge.ips_per_node"),
	},
	{
		Section:   "network",
		Types:     []string{"routed_bridge"},
		Option:    "tables",
		Default:   "main",
		Converter: converters.List,
		Text:      keywords.NewText(fs, "text/kw/node/network.routed_bridge.tables"),
		Example:   "main custom1 custom2",
	},
	{
		Section:     "network",
		Types:       []string{"routed_bridge"},
		Option:      "addr",
		Scopable:    true,
		DefaultText: keywords.NewText(fs, "text/kw/node/network.routed_bridge.addr.default"),
		Text:        keywords.NewText(fs, "text/kw/node/network.routed_bridge.addr"),
	},
	{
		Section:    "network",
		Types:      []string{"routed_bridge"},
		Option:     "tunnel",
		Default:    "auto",
		Candidates: []string{"auto", "always", "never"},
		Text:       keywords.NewText(fs, "text/kw/node/network.routed_bridge.tunnel"),
	},
	{
		Section: "network",
		Types:   []string{"bridge", "routed_bridge"},
		Option:  "network",
		Text:    keywords.NewText(fs, "text/kw/node/network.network"),
	},
	{
		Section:    "switch",
		Option:     "type",
		Candidates: []string{"brocade"},
		Required:   true,
		Text:       keywords.NewText(fs, "text/kw/node/switch.type"),
	},
	{
		Section: "switch",
		Types:   []string{"brocade"},
		Option:  "name",
		Example: "sansw1.my.corp",
		Text:    keywords.NewText(fs, "text/kw/node/switch.brocade.name"),
	},
	{
		Section:    "switch",
		Types:      []string{"brocade"},
		Option:     "method",
		Default:    "ssh",
		Candidates: []string{"telnet", "ssh"},
		Example:    "ssh",
		Text:       keywords.NewText(fs, "text/kw/node/switch.brocade.method"),
	},
	{
		Section:  "switch",
		Types:    []string{"brocade"},
		Option:   "username",
		Required: true,
		Example:  "admin",
		Text:     keywords.NewText(fs, "text/kw/node/switch.brocade.username"),
	},
	{
		Section: "switch",
		Types:   []string{"brocade"},
		Option:  "password",
		Example: "mysec/password",
		Text:    keywords.NewText(fs, "text/kw/node/switch.brocade.password"),
	},
	{
		Section: "switch",
		Types:   []string{"brocade"},
		Option:  "key",
		Example: "/path/to/key",
		Text:    keywords.NewText(fs, "text/kw/node/switch.brocade.key"),
	},
	{
		Section:    "array",
		Option:     "type",
		Candidates: []string{"freenas", "hds", "eva", "nexenta", "vioserver", "centera", "symmetrix", "emcvnx", "netapp", "hp3par", "ibmds", "ibmsvc", "xtremio", "dorado", "hoc"},
		Required:   true,
		Text:       keywords.NewText(fs, "text/kw/node/array.type"),
	},
	{
		Section:   "pool",
		Types:     []string{"dorado", "hoc"},
		Option:    "compression",
		Converter: converters.Bool,
		Default:   "false",
		Text:      keywords.NewText(fs, "text/kw/node/pool.compression"),
	},
	{
		Section: "pool",
		Types:   []string{"freenas"},
		Option:  "dedup",
		Default: "off",
		Text:    keywords.NewText(fs, "text/kw/node/pool.freenas.dedup"),
	},
	{
		Section:   "pool",
		Types:     []string{"dorado", "hoc"},
		Option:    "dedup",
		Converter: converters.Bool,
		Default:   "false",
		Text:      keywords.NewText(fs, "text/kw/node/pool.dedup"),
	},
	{
		Section:  "pool",
		Types:    []string{"dorado"},
		Option:   "hypermetrodomain",
		Required: false,
		Example:  "HyperMetroDomain_000",
		Text:     keywords.NewText(fs, "text/kw/node/pool.dorado.hypermetrodomain"),
	},
	{
		Section:  "array",
		Types:    []string{"dorado", "freenas", "hoc", "pure", "xtremio"},
		Option:   "api",
		Required: true,
		Example:  "https://array.opensvc.com/api/v1.0",
		Text:     keywords.NewText(fs, "text/kw/node/array.api"),
	},
	{
		Section: "array",
		Types:   []string{"hoc"},
		Option:  "http_proxy",
		Example: "http://proxy.mycorp:3158",
		Text:    keywords.NewText(fs, "text/kw/node/array.hoc.http_proxy"),
	},
	{
		Section: "array",
		Types:   []string{"hoc"},
		Option:  "https_proxy",
		Example: "https://proxy.mycorp:3158",
		Text:    keywords.NewText(fs, "text/kw/node/array.hoc.https_proxy"),
	},
	{
		Section:   "array",
		Types:     []string{"hoc"},
		Option:    "retry",
		Default:   "30",
		Converter: converters.Int,
		Text:      keywords.NewText(fs, "text/kw/node/array.hoc.retry"),
	},
	{
		Section:   "array",
		Types:     []string{"hoc"},
		Option:    "delay",
		Default:   "10s",
		Converter: converters.Duration,
		Text:      keywords.NewText(fs, "text/kw/node/array.hoc.delay"),
	},
	{
		Section:    "array",
		Types:      []string{"hoc"},
		Option:     "model",
		Required:   true,
		Example:    "VSP G350",
		Candidates: []string{"VSP G370", "VSP G700", "VSP G900", "VSP F370", "VSP F700", "VSP F900", "VSP G350", "VSP F350", "VSP G800", "VSP F800", "VSP G400", "VSP G600", "VSP F400", "VSP F600", "VSP G200", "VSP G1000", "VSP G1500", "VSP F1500", "Virtual Storage Platform", "HUS VM"},
		Text:       keywords.NewText(fs, "text/kw/node/array.hoc.model"),
	},
	{
		Section:  "array",
		Types:    []string{"centera", "eva", "hds", "ibmds", "ibmsvc", "freenas", "netapp", "nexenta", "vioserver", "xtremio", "dorado", "hoc"},
		Option:   "username",
		Required: true,
		Example:  "root",
		Text:     keywords.NewText(fs, "text/kw/node/array.username,required"),
	},
	{
		Section: "array",
		Types:   []string{"emcvnx", "hp3par", "symmetrix"},
		Option:  "username",
		Example: "root",
		Text:    keywords.NewText(fs, "text/kw/node/array.username,optional"),
	},
	{
		Section:  "array",
		Types:    []string{"pure"},
		Option:   "client_id",
		Example:  "bd2c75d0-f0d5-11ee-a362-8b0f2d1b83d7",
		Required: true,
		Text:     keywords.NewText(fs, "text/kw/node/array.pure.client_id"),
	},
	{
		Section:  "array",
		Types:    []string{"pure"},
		Option:   "key_id",
		Example:  "df80ae3a-f0d5-11ee-94c9-b7c8d2f57c4f",
		Required: true,
		Text:     keywords.NewText(fs, "text/kw/node/array.pure.key_id"),
	},
	{
		Section:   "array",
		Types:     []string{"pure", "hoc"},
		Option:    "insecure",
		Example:   "true",
		Default:   "false",
		Converter: converters.Bool,
		Text:      keywords.NewText(fs, "text/kw/node/array.pure.insecure"),
	},
	{
		Section:  "array",
		Types:    []string{"pure"},
		Option:   "issuer",
		Example:  "opensvc",
		Required: true,
		Text:     keywords.NewText(fs, "text/kw/node/array.pure.issuer"),
	},
	{
		Section:  "array",
		Types:    []string{"pure"},
		Option:   "secret",
		Example:  "system/sec/array1",
		Required: true,
		Text:     keywords.NewText(fs, "text/kw/node/array.pure.secret"),
	},
	{
		Section:  "array",
		Types:    []string{"pure"},
		Option:   "username",
		Example:  "opensvc",
		Required: true,
		Text:     keywords.NewText(fs, "text/kw/node/array.pure.username"),
	},
	{
		Section:  "array",
		Types:    []string{"centera", "eva", "hds", "freenas", "nexenta", "xtremio", "dorado", "hoc"},
		Option:   "password",
		Example:  "system/sec/array1",
		Required: true,
		Text:     keywords.NewText(fs, "text/kw/node/array.password,required"),
	},
	{
		Section: "array",
		Types:   []string{"emcvnx", "symmetrix"},
		Option:  "password",
		Example: "system/sec/array1",
		Text:    keywords.NewText(fs, "text/kw/node/array.password,optional"),
	},
	{
		Section:   "array",
		Types:     []string{"freenas", "dorado", "hoc"},
		Option:    "timeout",
		Converter: converters.Duration,
		Example:   "10s",
		Default:   "120s",
		Text:      keywords.NewText(fs, "text/kw/node/array.timeout"),
	},
	{
		Section: "array",
		Types:   []string{"dorado", "hoc"},
		Option:  "name",
		Example: "a09",
		Text:    keywords.NewText(fs, "text/kw/node/array.name"),
	},
	{
		Section: "array",
		Types:   []string{"symmetrix"},
		Option:  "name",
		Example: "00012345",
		Text:    keywords.NewText(fs, "text/kw/node/array.symmetrix.name"),
	},
	{
		Section: "array",
		Types:   []string{"symmetrix"},
		Option:  "symcli_path",
		Default: "/usr/symcli",
		Example: "/opt/symcli",
		Text:    keywords.NewText(fs, "text/kw/node/array.symmetrix.symcli_path"),
	},
	{
		Section: "array",
		Types:   []string{"symmetrix"},
		Option:  "symcli_connect",
		Example: "MY_SYMAPI_SERVER",
		Text:    keywords.NewText(fs, "text/kw/node/array.symmetrix.symcli_connect"),
	},
	{
		Section:  "array",
		Types:    []string{"centera", "netapp"},
		Option:   "server",
		Required: true,
		Example:  "centera1",
		Text:     keywords.NewText(fs, "text/kw/node/array.server"),
	},
	{
		Section:  "array",
		Types:    []string{"centera"},
		Option:   "java_bin",
		Required: true,
		Example:  "/opt/java/bin/java",
		Text:     keywords.NewText(fs, "text/kw/node/array.centera.java_bin"),
	},
	{
		Section:  "array",
		Types:    []string{"centera"},
		Option:   "jcass_dir",
		Required: true,
		Example:  "/opt/centera/LIB",
		Text:     keywords.NewText(fs, "text/kw/node/array.centera.jcass_dir"),
	},
	{
		Section:    "array",
		Types:      []string{"emcvnx"},
		Option:     "method",
		Default:    "secfile",
		Candidates: []string{"secfile", "credentials"},
		Example:    "secfile",
		Text:       keywords.NewText(fs, "text/kw/node/array.emcvnx.secfile"),
	},
	{
		Section:  "array",
		Types:    []string{"emcvnx"},
		Option:   "spa",
		Required: true,
		Example:  "array1-a",
		Text:     keywords.NewText(fs, "text/kw/node/array.emcvnx.spa"),
	},
	{
		Section:  "array",
		Types:    []string{"emcvnx"},
		Option:   "spb",
		Required: true,
		Example:  "array1-b",
		Text:     keywords.NewText(fs, "text/kw/node/array.emcvnx.spb"),
	},
	{
		Section: "array",
		Types:   []string{"emcvnx"},
		Option:  "scope",
		Default: "0",
		Example: "1",
		Text:    keywords.NewText(fs, "text/kw/node/array.emcvnx.scope"),
	},
	{
		Section:  "array",
		Types:    []string{"eva"},
		Option:   "manager",
		Required: true,
		Example:  "evamanager.mycorp",
		Text:     keywords.NewText(fs, "text/kw/node/array.eva.manager"),
	},
	{
		Section: "array",
		Types:   []string{"eva"},
		Option:  "bin",
		Example: "/opt/sssu/bin/sssu",
		Text:    keywords.NewText(fs, "text/kw/node/array.eva.bin"),
	},
	{
		Section: "array",
		Types:   []string{"hds"},
		Option:  "bin",
		Example: "/opt/hds/bin/HiCommandCLI",
		Text:    keywords.NewText(fs, "text/kw/node/array.hds.bin"),
	},
	{
		Section: "array",
		Types:   []string{"hds"},
		Option:  "jre_path",
		Example: "/opt/java",
		Text:    keywords.NewText(fs, "text/kw/node/array.hds.jre_path"),
	},
	{
		Section: "array",
		Types:   []string{"hds"},
		Option:  "name",
		Example: "HUSVM.1234",
		Text:    keywords.NewText(fs, "text/kw/node/array.hds.name"),
	},
	{
		Section:  "array",
		Types:    []string{"hds"},
		Option:   "url",
		Required: true,
		Example:  "https://hdsmanager/",
		Text:     keywords.NewText(fs, "text/kw/node/array.hds.url"),
	},
	{
		Section:    "array",
		Types:      []string{"hp3par"},
		Option:     "method",
		Default:    "ssh",
		Candidates: []string{"proxy", "cli", "ssh"},
		Example:    "ssh",
		Text:       keywords.NewText(fs, "text/kw/node/array.hp3par.method"),
	},
	{
		Section:     "array",
		Types:       []string{"hp3par"},
		Option:      "manager",
		Example:     "mymanager.mycorp",
		DefaultText: keywords.NewText(fs, "text/kw/node/array.hp3par.manager.default"),
		Text:        keywords.NewText(fs, "text/kw/node/array.hp3par.manager"),
	},
	{
		Section: "array",
		Types:   []string{"hp3par"},
		Option:  "key",
		Example: "/path/to/key",
		Text:    keywords.NewText(fs, "text/kw/node/array.hp3par.key"),
	},
	{
		Section: "array",
		Types:   []string{"hp3par"},
		Option:  "pwf",
		Example: "/path/to/pwf",
		Text:    keywords.NewText(fs, "text/kw/node/array.hp3par.pwf"),
	},
	{
		Section: "array",
		Types:   []string{"hp3par"},
		Option:  "cli",
		Default: "3parcli",
		Example: "/path/to/pwf",
		Text:    keywords.NewText(fs, "text/kw/node/array.hp3par.cli"),
	},
	{
		Section:  "array",
		Types:    []string{"ibmds"},
		Option:   "hmc1",
		Required: true,
		Example:  "hmc1.mycorp",
		Text:     keywords.NewText(fs, "text/kw/node/array.ibmds.hmc1"),
	},
	{
		Section:  "array",
		Types:    []string{"ibmds"},
		Option:   "hmc2",
		Required: true,
		Example:  "hmc2.mycorp",
		Text:     keywords.NewText(fs, "text/kw/node/array.ibmds.hmc2"),
	},
	{
		Section:  "array",
		Types:    []string{"netapp", "ibmsvc", "vioserver"},
		Option:   "key",
		Required: true,
		Example:  "/path/to/key",
		Text:     keywords.NewText(fs, "text/kw/node/array.key,required"),
	},
	{
		Section:   "array",
		Types:     []string{"nexenta"},
		Option:    "port",
		Default:   "2000",
		Converter: converters.Int,
		Example:   "2000",
		Text:      keywords.NewText(fs, "text/kw/node/array.nexenta.port"),
	},
}

var nodeKeywordStore = keywords.Store(append(nodePrivateKeywords, nodeCommonKeywords...))

func (t Node) KeywordLookup(k key.T, sectionType string) keywords.Keyword {
	return keywordLookup(nodeKeywordStore, k, naming.KindInvalid, sectionType)
}
