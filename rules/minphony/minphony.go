// Package minphony implements the ruleset for making sure required minimum
// phony targets are present
package minphony

import (
	"fmt"
	"strings"

	"github.com/mrtazz/checkmake/parser"
	"github.com/mrtazz/checkmake/rules"
)

var (
	defaultRequired = []string{
	}
)

func init() {
	rules.RegisterRule(&MinPhony{required: defaultRequired})
}

// MinPhony is an empty struct on which to call the rule functions
type MinPhony struct {
	required []string
}

// Name returns the name of the rule
func (r *MinPhony) Name() string {
	return "minphony"
}

// Description returns the description of the rule
func (r *MinPhony) Description() string {
	return "Minimum required phony targets must be present"
}

// Run executes the rule logic
func (r *MinPhony) Run(makefile parser.Makefile, _ rules.RuleConfig) rules.RuleViolationList {
	ret := rules.RuleViolationList{}

	ruleIndex := make(map[string]bool)
	ruleLineNumber := 0
	for _, variable := range makefile.Variables {
		if variable.Name == "PHONY" {
			ruleLineNumber = variable.LineNumber - 1
			for _, phony := range strings.Split(variable.Assignment, " ") {
				ruleIndex[phony] = true
			}
		}
	}

	for _, reqRule := range r.required {
		_, ok := ruleIndex[reqRule]
		if !ok {
			ret = append(ret, rules.RuleViolation{
				Rule:       "minphony",
				Violation:  fmt.Sprintf("Missing required phony target %q", reqRule),
				FileName:   makefile.FileName,
				LineNumber: ruleLineNumber,
			})
		}
	}

	return ret
}
