// Package reqphony implements the ruleset for making sure required phony
// targets are present
package reqphony

import (
	"fmt"
	"strings"

	"github.com/mrtazz/checkmake/parser"
	"github.com/mrtazz/checkmake/rules"
)

var (
	defaultRequired = []string{
		"all",
		"clean",
		"test",
	}
)

func init() {
	rules.RegisterRule(&ReqPhony{Required: defaultRequired})
}

// ReqPhony is an empty struct on which to call the rule functions
type ReqPhony struct {
	Required []string
}

// Name returns the name of the rule
func (r *ReqPhony) Name() string {
	return "reqphony"
}

// Description returns the description of the rule
func (r *ReqPhony) Description() string {
	return "Required phony targets must be present"
}

// Run executes the rule logic
func (r *ReqPhony) Run(makefile parser.Makefile, config rules.RuleConfig) rules.RuleViolationList {
	ret := rules.RuleViolationList{}

	ruleIndex := make(map[string]bool)

	for _, variable := range makefile.Variables {
		if variable.Name == "PHONY" {
			for _, phony := range strings.Split(variable.Assignment, " ") {
				ruleIndex[phony] = true
			}
		}
	}

	for _, reqRule := range r.Required {
		_, ok := ruleIndex[reqRule]
		if !ok {
			ret = append(ret, rules.RuleViolation{
				Rule:       "reqphony",
				Violation:  fmt.Sprintf("Missing required phony target %q", reqRule),
				LineNumber: 0,
			})
		}
	}

	return ret
}
