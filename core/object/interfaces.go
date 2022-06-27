package object

import (
	"os"
	"os/user"

	"opensvc.com/opensvc/core/instance"
	"opensvc.com/opensvc/core/keyop"
	"opensvc.com/opensvc/core/rawconfig"
	"opensvc.com/opensvc/core/resource"
	"opensvc.com/opensvc/core/resourceset"
	"opensvc.com/opensvc/core/schedule"
	"opensvc.com/opensvc/core/xconfig"
	"opensvc.com/opensvc/util/key"
	"opensvc.com/opensvc/util/render/tree"
	"opensvc.com/opensvc/util/timestamp"
)

type (
	// Renderer is implemented by data type stored in ActionResults.Data.
	Renderer interface {
		Render() string
	}
	// LoadTreeNoder is implemented by data type stored in ActionResults.Data.
	LoadTreeNoder interface {
		LoadTreeNode(*tree.Node)
	}

	// SecureKeystorer is implemented by encrypting Keystore object kinds (usr, sec).
	SecureKeystorer interface {
		GenCert(OptsGenCert) error
	}

	// Keystorer is implemented by Keystore object kinds (usr, sec, cfg).
	Keystorer interface {
		Add(OptsAdd) error
		Change(OptsAdd) error
		Decode(OptsDecode) ([]byte, error)
		Keys(OptsKeys) ([]string, error)
		MatchingKeys(string) ([]string, error)
		Remove(OptsRemove) error
		EditKey(OptsEditKey) error
		Install(OptsInstall) error

		InstallKey(string, string, *os.FileMode, *os.FileMode, *user.User, *user.Group) error
	}

	// Baser is implemented by all object kinds.
	Baser interface {
		Status(OptsStatus) (instance.Status, error)
		IsVolatile() bool
		ResourceSets() resourceset.L
		Resources() resource.Drivers
		ResourceByID(rid string) resource.Driver
	}

	// Actor is implemented by object kinds supporting start, stop, ...
	Actor interface {
		Freezer
		Restart(OptsStart) error
		Run(OptsRun) error
		Start(OptsStart) error
		Stop(OptsStop) error
		Provision(OptsProvision) error
		Unprovision(OptsUnprovision) error
		SetProvisioned(OptsSetProvisioned) error
		SetUnprovisioned(OptsSetUnprovisioned) error
		SyncResync(OptsSyncResync) error
	}

	Enterer interface {
		Enter(OptsEnter) error
	}

	// Freezer is implemented by object kinds supporting freeze and thaw.
	Freezer interface {
		Freeze() error
		Unfreeze() error
		Thaw() error
		Frozen() timestamp.T
	}

	// Configurer is implemented by object kinds supporting get, set, unset, eval, edit, ...
	Configurer interface {
		ConfigFile() string
		Config() *xconfig.T
		EditConfig(OptsEditConfig) error
		PrintConfig(OptsPrintConfig) (rawconfig.T, error)
		ValidateConfig(OptsValidateConfig) (xconfig.ValidateAlerts, error)
		Eval(OptsEval) (interface{}, error)
		Get(OptsGet) (interface{}, error)
		Set(OptsSet) error
		Unset(OptsUnset) error
		Delete(OptsDelete) error
		SetKeys(kops ...keyop.T) error
		UnsetKeys(kws ...key.T) error
		Doc(OptsDoc) (string, error)
	}

	// ResourceLister provides a method to list and filter resources
	ResourceLister interface {
		Resources() resource.Drivers
		IsDesc() bool
	}

	scheduler interface {
		Schedules() schedule.Table
	}
)
