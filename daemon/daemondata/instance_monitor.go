package daemondata

import (
	"github.com/opensvc/om3/core/instance"
	"github.com/opensvc/om3/daemon/msgbus"
	"github.com/opensvc/om3/util/jsondelta"
)

// onInstanceMonitorDeleted delete .cluster.node.<node>.instance.<path>.monitor
func (d *data) onInstanceMonitorDeleted(m msgbus.InstanceMonitorDeleted) {
	d.statCount[idDelInstanceMonitor]++
	s := m.Path.String()
	if inst, ok := d.pending.Cluster.Node[d.localNode].Instance[s]; ok && inst.Monitor != nil {
		inst.Monitor = nil
		d.pending.Cluster.Node[d.localNode].Instance[s] = inst
		op := jsondelta.Operation{
			OpPath: jsondelta.OperationPath{"instance", s, "monitor"},
			OpKind: "remove",
		}
		d.pendingOps = append(d.pendingOps, op)
	}
}

// onInstanceMonitorUpdated updates .cluster.node.<node>.instance.<path>.monitor
func (d *data) onInstanceMonitorUpdated(m msgbus.InstanceMonitorUpdated) {
	d.statCount[idSetInstanceMonitor]++
	var op jsondelta.Operation
	s := m.Path.String()
	value := &m.Value
	if inst, ok := d.pending.Cluster.Node[d.localNode].Instance[s]; ok {
		inst.Monitor = value
		d.pending.Cluster.Node[d.localNode].Instance[s] = inst

	} else {
		d.pending.Cluster.Node[d.localNode].Instance[s] = instance.Instance{Monitor: value}
		op = jsondelta.Operation{
			OpPath:  jsondelta.OperationPath{"instance", s},
			OpValue: jsondelta.NewOptValue(struct{}{}),
			OpKind:  "replace",
		}
		d.pendingOps = append(d.pendingOps, op)
	}
	op = jsondelta.Operation{
		OpPath:  jsondelta.OperationPath{"instance", s, "monitor"},
		OpValue: jsondelta.NewOptValue(*value),
		OpKind:  "replace",
	}
	d.pendingOps = append(d.pendingOps, op)
}
