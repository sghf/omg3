package compliance

import (
	"fmt"

	"opensvc.com/opensvc/core/collector"
	"opensvc.com/opensvc/util/hostname"
)

type (
	Rulesets map[string]Ruleset
	Ruleset  struct {
		Filter string
		Name   string
		Vars   []Var
	}
)

func (t T) GetRulesets() (Rulesets, error) {
	rulesets := make(Rulesets)
	err := t.collectorClient.CallFor(&rulesets, "comp_get_ruleset", hostname.Hostname())
	if err != nil {
		return nil, err
	}
	return rulesets, nil
}

func (t Ruleset) Render() string {
	buff := fmt.Sprintf(" %s", t.Name)
	if t.Filter != "" {
		buff += fmt.Sprintf(" (%s)", t.Filter)
	}
	buff += "\n"
	for _, v := range t.Vars {
		buff += fmt.Sprintf("  %s\n", v)
	}
	return buff
}

func (t Rulesets) Render() string {
	buff := "rulesets:\n"
	for _, rset := range t {
		buff += rset.Render()
	}
	return buff
}

func (t T) ListRulesets(filter string) ([]string, error) {
	var err error
	data := make([]string, 0)
	if filter == "" {
		filter = "%"
	}
	err = t.collectorClient.CallFor(&data, "comp_list_rulesets", filter, hostname.Hostname())
	if err != nil {
		return data, err
	}
	return data, nil
}

func (t T) AttachRuleset(s string) error {
	response, err := t.collectorClient.Call("comp_attach_ruleset", hostname.Hostname(), s)
	if err != nil {
		return err
	}
	collector.LogSimpleResponse(response, t.log)
	return nil
}

func (t T) DetachRuleset(s string) error {
	response, err := t.collectorClient.Call("comp_detach_ruleset", hostname.Hostname(), s)
	if err != nil {
		return err
	}
	collector.LogSimpleResponse(response, t.log)
	return nil
}
