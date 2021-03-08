package object

import (
	"bytes"
	"encoding/json"
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"opensvc.com/opensvc/config"
)

type (
	// Path represents an opensvc object path-like identifier. Ex: ns1/svc/svc1
	Path struct {
		// Name is the name part of the path
		Name string
		// Namespace is the namespace part of the path
		Namespace string
		// Kind is the kinf part of the path
		Kind Kind
	}
)

const (
	// Separator is the character separating a path's namespace, kind and name
	Separator = "/"

	hostnameRegexStringRFC952 = `^[a-zA-Z]([a-zA-Z0-9\-]+[\.]?)*[a-zA-Z0-9]$` // https://tools.ietf.org/html/rfc952
	fqdnRegexStringRFC1123    = `^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\.?$`
)

var (

	// ErrPathInvalid is raised when the path allocator can not return a Path
	// because one of the path element is not valid.
	ErrPathInvalid = errors.New("invalid path")

	hostnameRegexRFC952 = regexp.MustCompile(hostnameRegexStringRFC952)
	fqdnRegexRFC1123    = regexp.MustCompile(fqdnRegexStringRFC1123)
)

// NewPath allocates a new path type from its elements
func NewPath(name string, namespace string, kind string) (Path, error) {
	var path Path
	name = strings.ToLower(name)
	namespace = strings.ToLower(namespace)
	kind = strings.ToLower(kind)
	if name == "" {
		return path, errors.Wrap(ErrPathInvalid, "name is empty")
	}
	if kind == "" {
		kind = "svc"
	}
	k := NewKind(kind)
	if k == KindInvalid {
		return path, errors.Wrapf(ErrPathInvalid, "invalid kind %s", kind)
	}
	if namespace == "" {
		namespace = "root"
	}
	if kind == "" {
		kind = "svc"
	}
	if !hostnameRegexRFC952.MatchString(name) {
		return path, errors.Wrapf(ErrPathInvalid, "invalid name %s (rfc952)", kind)
	}
	if !hostnameRegexRFC952.MatchString(namespace) {
		return path, errors.Wrapf(ErrPathInvalid, "invalid namespace %s (rfc952)", kind)
	}
	path.Namespace = namespace
	path.Name = name
	path.Kind = k
	return path, nil
}

func (t Path) String() string {
	var s string
	if t.Kind == KindInvalid {
		return ""
	}
	if t.Namespace != "" && t.Namespace != "root" {
		s += t.Namespace + Separator
	}
	if t.Kind != KindSvc || s != "" {
		s += t.Kind.String() + Separator
	}
	return s + t.Name
}

// NewPathFromString returns a new path struct from a path string representation
func NewPathFromString(s string) (Path, error) {
	var (
		name      string
		namespace string
		kind      string
	)
	s = strings.ToLower(s)
	l := strings.Split(s, Separator)
	switch len(l) {
	case 3:
		namespace = l[0]
		kind = l[1]
		name = l[2]
	case 2:
		namespace = "root"
		kind = l[0]
		name = l[1]
	case 1:
		namespace = "root"
		kind = "svc"
		name = l[0]
	}
	return NewPath(name, namespace, kind)
}

// MarshalJSON implements the json interface
func (t Path) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(t.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON implements the json interface
func (t *Path) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*t, err = NewPathFromString(j)
	if err != nil {
		return err
	}
	return nil
}

// NewObject allocates a new kinded object
func (t Path) NewObject() interface{} {
	switch t.Kind {
	case KindSvc:
		return NewSvc(t)
	case KindVol:
		return NewVol(t)
	case KindCfg:
		return NewCfg(t)
	case KindSec:
		return NewSec(t)
	case KindUsr:
		return NewUsr(t)
	default:
		return nil
	}
}

// ConfigFile returns the absolute path of an opensvc object configuration
// file.
func (t Path) ConfigFile() string {
	p := t.String()
	switch t.Namespace {
	case "", "root":
		p = fmt.Sprintf("%s/%s.conf", config.Node.Paths.Etc, p)
	default:
		p = fmt.Sprintf("%s/%s.conf", config.Node.Paths.EtcNs, p)
	}
	return filepath.FromSlash(p)
}
