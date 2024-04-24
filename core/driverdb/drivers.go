package driverdb

import (
	// Uncomment to load
	_ "github.com/opensvc/om3/drivers/arrayfreenas"
	_ "github.com/opensvc/om3/drivers/arrayhoc"
	_ "github.com/opensvc/om3/drivers/arraypure"
	_ "github.com/opensvc/om3/drivers/pooldirectory"
	_ "github.com/opensvc/om3/drivers/poolfreenas"
	_ "github.com/opensvc/om3/drivers/poolpure"
	_ "github.com/opensvc/om3/drivers/poolshm"
	_ "github.com/opensvc/om3/drivers/poolvirtual"
	_ "github.com/opensvc/om3/drivers/poolzpool"
	_ "github.com/opensvc/om3/drivers/resappforking"
	_ "github.com/opensvc/om3/drivers/resappsimple"
	_ "github.com/opensvc/om3/drivers/rescertificatetls"
	_ "github.com/opensvc/om3/drivers/resdiskdisk"
	_ "github.com/opensvc/om3/drivers/resdiskloop"
	_ "github.com/opensvc/om3/drivers/resdisklv"
	_ "github.com/opensvc/om3/drivers/resdiskmd"
	_ "github.com/opensvc/om3/drivers/resdiskraw"
	_ "github.com/opensvc/om3/drivers/resdiskvg"
	_ "github.com/opensvc/om3/drivers/resexposeenvoy"
	_ "github.com/opensvc/om3/drivers/resfsdir"
	_ "github.com/opensvc/om3/drivers/resfsflag"
	_ "github.com/opensvc/om3/drivers/resfshost"
	_ "github.com/opensvc/om3/drivers/resfszfs"
	_ "github.com/opensvc/om3/drivers/resiphost"
	_ "github.com/opensvc/om3/drivers/resiproute"
	_ "github.com/opensvc/om3/drivers/resrouteenvoy"
	_ "github.com/opensvc/om3/drivers/ressharenfs"
	_ "github.com/opensvc/om3/drivers/ressyncrsync"
	_ "github.com/opensvc/om3/drivers/ressynczfs"
	_ "github.com/opensvc/om3/drivers/ressynczfssnap"
	_ "github.com/opensvc/om3/drivers/restaskdocker"
	_ "github.com/opensvc/om3/drivers/restaskhost"
	_ "github.com/opensvc/om3/drivers/resvhostenvoy"
	_ "github.com/opensvc/om3/drivers/resvol"
)
