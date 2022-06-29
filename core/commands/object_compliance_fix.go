package commands

import (
	"github.com/spf13/cobra"
	"opensvc.com/opensvc/core/flag"
	"opensvc.com/opensvc/core/object"
	"opensvc.com/opensvc/core/objectaction"
	"opensvc.com/opensvc/core/path"
)

type (
	// CmdObjectComplianceFix is the cobra flag set of the sysreport command.
	CmdObjectComplianceFix struct {
		OptsGlobal
		object.OptsObjectComplianceFix
	}
)

// Init configures a cobra command and adds it to the parent command.
func (t *CmdObjectComplianceFix) Init(kind string, parent *cobra.Command, selector *string) {
	cmd := t.cmd(kind, selector)
	parent.AddCommand(cmd)
	flag.Install(cmd, t)
}

func (t *CmdObjectComplianceFix) cmd(kind string, selector *string) *cobra.Command {
	return &cobra.Command{
		Use:   "fix",
		Short: "run compliance fixes",
		Run: func(_ *cobra.Command, _ []string) {
			t.run(selector, kind)
		},
	}
}

func (t *CmdObjectComplianceFix) run(selector *string, kind string) {
	mergedSelector := mergeSelector(*selector, t.ObjectSelector, kind, "")
	objectaction.New(
		objectaction.LocalFirst(),
		objectaction.WithLocal(t.Local),
		objectaction.WithColor(t.Color),
		objectaction.WithFormat(t.Format),
		objectaction.WithObjectSelector(mergedSelector),
		objectaction.WithRemoteNodes(t.NodeSelector),
		objectaction.WithServer(t.Server),
		objectaction.WithRemoteAction("compliance fix"),
		objectaction.WithRemoteOptions(map[string]interface{}{
			"format":    t.Format,
			"force":     t.Force,
			"module":    t.Module,
			"moduleset": t.Moduleset,
			"attach":    t.Attach,
		}),
		objectaction.WithLocalRun(func(p path.T) (interface{}, error) {
			if o, err := object.NewSvc(p); err != nil {
				return nil, err
			} else {
				return o.ComplianceFix(t.OptsObjectComplianceFix)
			}
		}),
	).Do()
}
