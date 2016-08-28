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
	rules.RegisterRule(rules.Rule{
		Name:        "rule1",
		Description: "Every target without a body needs to be marked PHONY",
		Rule:        ruleset})
}

func ruleset(makefile parser.Makefile, config rules.RuleConfig) rules.RuleViolationList {
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
				Rule:      "rule1",
				Violation: fmt.Sprintf("Target '%s' should be marked PHONY.", rule.Target),
			})
		}
	}

	return ret
}
