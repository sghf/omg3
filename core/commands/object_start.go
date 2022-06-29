package commands

import (
	"github.com/spf13/cobra"
	"opensvc.com/opensvc/core/flag"
	"opensvc.com/opensvc/core/object"
	"opensvc.com/opensvc/core/objectaction"
	"opensvc.com/opensvc/core/path"
)

type (
	// CmdObjectStart is the cobra flag set of the start command.
	CmdObjectStart struct {
		OptsGlobal
		OptsAsync
		object.OptsStart
	}
)

// Init configures a cobra command and adds it to the parent command.
func (t *CmdObjectStart) Init(kind string, parent *cobra.Command, selector *string) {
	cmd := t.cmd(kind, selector)
	parent.AddCommand(cmd)
	flag.Install(cmd, t)
}

func (t *CmdObjectStart) cmd(kind string, selector *string) *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "start the selected objects",
		Run: func(cmd *cobra.Command, args []string) {
			t.run(selector, kind)
		},
	}
}

func (t *CmdObjectStart) run(selector *string, kind string) {
	mergedSelector := mergeSelector(*selector, t.ObjectSelector, kind, "")
	objectaction.New(
		objectaction.WithObjectSelector(mergedSelector),
		objectaction.WithRID(t.RID),
		objectaction.WithTag(t.Tag),
		objectaction.WithSubset(t.Subset),
		objectaction.WithLocal(t.Local),
		objectaction.WithFormat(t.Format),
		objectaction.WithColor(t.Color),
		objectaction.WithRemoteNodes(t.NodeSelector),
		objectaction.WithRemoteAction("start"),
		objectaction.WithAsyncTarget("started"),
		objectaction.WithAsyncWatch(t.Watch),
		objectaction.WithDigest(),
		objectaction.WithLocalRun(func(p path.T) (interface{}, error) {
			o, err := object.NewActorFromPath(p)
			if err != nil {
				return nil, err
			}
			return nil, o.Start(t.OptsStart)
		}),
	).Do()
}
