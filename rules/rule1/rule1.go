// Package rule1 implements the ruleset for making sure all targets that don't
// have a rule body are marked PHONY
package rule1

import (
	"fmt"
	"github.com/mrtazz/checkmake/parser"
	"github.com/mrtazz/checkmake/rules"
	"strings"
)

func init() {
	rules.RegisterRule(&Rule1{})
}

// Rule1 is an empty struct on which to call the rule functions
type Rule1 struct {
}

// Name returns the name of the rule
func (r *Rule1) Name() string {
	return "rule1"
}

// Description returns the description of the rule
func (r *Rule1) Description() string {
	return "Every target without a body needs to be marked PHONY"
}

// Run executes the rule logic
func (r *Rule1) Run(makefile parser.Makefile, config rules.RuleConfig) rules.RuleViolationList {
	ret := rules.RuleViolationList{}

	ruleIndex := make(map[string]bool)

	for _, variable := range makefile.Variables {
		if variable.Name == "PHONY" {
			for _, phony := range strings.Split(variable.Assignment, " ") {
				ruleIndex[phony] = true
			}
		}
	}

	for _, rule := range makefile.Rules {
		_, ok := ruleIndex[rule.Target]
		if len(rule.Body) == 0 && ok == false {
			ret = append(ret, rules.RuleViolation{
				Rule:       "rule1",
				Violation:  fmt.Sprintf("Target '%s' should be marked PHONY.", rule.Target),
				LineNumber: rule.LineNumber,
			})
		}
	}

	return ret
}
