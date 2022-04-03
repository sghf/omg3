package commands

import (
	"github.com/spf13/cobra"
	"opensvc.com/opensvc/core/entrypoints/nodeaction"
	"opensvc.com/opensvc/core/flag"
	"opensvc.com/opensvc/core/object"
)

type (
	// CmdNodeComplianceShowRuleset is the cobra flag set of the sysreport command.
	CmdNodeComplianceShowRuleset struct {
		object.OptsNodeComplianceShowRuleset
	}
)

// Init configures a cobra command and adds it to the parent command.
func (t *CmdNodeComplianceShowRuleset) Init(parent *cobra.Command) {
	cmd := t.cmd()
	parent.AddCommand(cmd)
	flag.Install(cmd, &t.OptsNodeComplianceShowRuleset)
}

func (t *CmdNodeComplianceShowRuleset) cmd() *cobra.Command {
	return &cobra.Command{
		Use:     "ruleset",
		Short:   "Show compliance rules applying to this node.",
		Aliases: []string{"rulese", "rules", "rule", "rul", "ru"},
		Run: func(_ *cobra.Command, _ []string) {
			t.run()
		},
	}
}

func (t *CmdNodeComplianceShowRuleset) run() {
	nodeaction.New(
		nodeaction.WithLocal(t.Global.Local),
		nodeaction.WithRemoteNodes(t.Global.NodeSelector),
		nodeaction.WithFormat(t.Global.Format),
		nodeaction.WithColor(t.Global.Color),
		nodeaction.WithServer(t.Global.Server),
		nodeaction.WithRemoteAction("compliance show ruleset"),
		nodeaction.WithRemoteOptions(map[string]interface{}{
			"format":  t.Global.Format,
			"ruleset": t.Ruleset,
		}),
		nodeaction.WithLocalRun(func() (interface{}, error) {
			return object.NewNode().ComplianceShowRuleset(t.OptsNodeComplianceShowRuleset)
		}),
	).Do()
}