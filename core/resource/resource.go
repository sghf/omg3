package resource

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"opensvc.com/opensvc/core/drivergroup"
	"opensvc.com/opensvc/core/status"
)

type (
	// DriverID identifies a driver.
	DriverID struct {
		Group drivergroup.T
		Name  string
	}

	// Driver exposes what can be done with a resource
	Driver interface {
		Label() string
		Manifest() Manifest
		Start() error
		Stop() error
		Status() status.T

		// common
		SetRID(string)
		RID() string
		RSubset() string
		RLog() *Log
	}

	Aborter interface {
		Abort() bool
	}

	// T is the resource type, embedded in each drivers type
	T struct {
		Driver
		ResourceID string `json:"rid"`
		Subset     string `json:"subset"`
		Disable    bool   `json:"disable"`
		Log        Log    `json:"-"`
	}

	// OutputStatus is the structure representing the resource status,
	// which is embedded in the instance status.
	OutputStatus struct {
		Label  string      `json:"label"`
		Status status.T    `json:"status"`
		Subset string      `json:"subset,omitempty"`
		Type   string      `json:"type"`
		Log    []*LogEntry `json:"log,omitempty"`
	}
)

func (t DriverID) String() string {
	if t.Name == "" {
		return t.Group.String()
	}
	return fmt.Sprintf("%s.%s", t.Group, t.Name)
}

func ParseDriverID(s string) *DriverID {
	l := strings.SplitN(s, ".", 2)
	g := drivergroup.New(l[0])
	return &DriverID{
		Group: g,
		Name:  l[1],
	}
}

func NewDriverID(group drivergroup.T, name string) *DriverID {
	return &DriverID{
		Group: group,
		Name:  name,
	}
}

var drivers = make(map[DriverID]func() Driver)

func Register(group drivergroup.T, name string, f func() Driver) {
	driverID := NewDriverID(group, name)
	drivers[*driverID] = f
}

func (t DriverID) NewResourceFunc() func() Driver {
	drv, ok := drivers[t]
	if !ok {
		return nil
	}
	return drv
}

func (t T) String() string {
	return fmt.Sprintf("<Resource %s>", t.ResourceID)
}

// RSubset returns the resource subset name
func (t T) RSubset() string {
	return t.Subset
}

// RLog return a reference to the resource log
func (t *T) RLog() *Log {
	return &t.Log
}

// RID return a reference to the resource log
func (t T) RID() string {
	return t.ResourceID
}

// SetRID sets the resource identifier
func (t *T) SetRID(v string) {
	t.ResourceID = v
}

func formatResourceType(r Driver) string {
	m := r.Manifest()
	return fmt.Sprintf("%s.%s", m.Group, m.Name)
}

func formatResourceLabel(r Driver) string {
	return fmt.Sprintf("%s %s", formatResourceType(r), r.Label())
}

// Start activates a resource interfacer
func Start(r Driver) error {
	return r.Start()
}

// Stop deactivates a resource interfacer
func Stop(r Driver) error {
	return r.Stop()
}

// Status evaluates the status of a resource interfacer
func Status(r Driver) status.T {
	return r.Status()
}

func printStatus(r Driver) error {
	data := OutputStatus{
		Label:  formatResourceLabel(r),
		Type:   formatResourceType(r),
		Status: Status(r),
		Subset: r.RSubset(),
		Log:    r.RLog().Entries(),
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "    ")
	return enc.Encode(data)
}

func printManifest(r Driver) error {
	m := r.Manifest()
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "    ")
	return enc.Encode(m)
}

func printHelp(r Driver) error {
	fmt.Println(`Environment variables:
  RES_ACTION=start|stop|status|manifest

Stdin:
  json formatted context data
	`)
	return nil
}

// Action calls the resource method set as the RES_ACTION environment variable
func Action(r Driver) error {
	action := os.Getenv("RES_ACTION")
	switch action {
	case "status":
		return printStatus(r)
	case "stop":
		return Stop(r)
	case "start":
		return Start(r)
	case "manifest":
		return printManifest(r)
	default:
		return printHelp(r)
	}
}
