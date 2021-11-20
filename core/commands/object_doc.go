package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"opensvc.com/opensvc/core/flag"
	"opensvc.com/opensvc/core/object"
	"opensvc.com/opensvc/core/objectaction"
	"opensvc.com/opensvc/core/path"
)

type (
	// CmdObjectDoc is the cobra flag set of the doc command.
	CmdObjectDoc struct {
		object.OptsDoc
	}
)

// Init configures a cobra command and adds it to the parent command.
func (t *CmdObjectDoc) Init(kind string, parent *cobra.Command, selector *string) {
	cmd := t.cmd(kind, selector)
	parent.AddCommand(cmd)
	flag.Install(cmd, &t.OptsDoc)
}

func (t *CmdObjectDoc) cmd(kind string, selector *string) *cobra.Command {
	return &cobra.Command{
		Use:   "doc",
		Short: "print the documentation of the selected keywords",
		Run: func(cmd *cobra.Command, args []string) {
			t.run(selector, kind)
		},
	}
}

func (t *CmdObjectDoc) run(selector *string, kind string) {
	mergedSelector := mergeSelector(*selector, t.Global.ObjectSelector, kind, "")
	objectaction.New(
		objectaction.LocalFirst(),
		objectaction.WithLocal(t.Global.Local),
		objectaction.WithColor(t.Global.Color),
		objectaction.WithFormat(t.Global.Format),
		objectaction.WithObjectSelector(mergedSelector),
		objectaction.WithRemoteNodes(t.Global.NodeSelector),
		objectaction.WithRemoteAction("doc"),
		objectaction.WithRemoteOptions(map[string]interface{}{
			"kw":     t.Keyword,
			"driver": t.Driver,
		}),
		objectaction.WithLocalRun(func(p path.T) (interface{}, error) {
			o, err := object.NewFromPath(p)
			if err != nil {
				return nil, err
			}
			c, ok := o.(object.Configurer)
			if !ok {
				return nil, fmt.Errorf("%s is not a configurer", o)
			}
			return c.Doc(t.OptsDoc)
		}),
	).Do()
}
