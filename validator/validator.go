// Package validator holds the basic engine that runs rule based checks
// against a Makefile struct
package validator

import (
	"github.com/mrtazz/checkmake/parser"
	"github.com/mrtazz/checkmake/rules"
)

// Config is a struct to configure the validator
type Config struct {
	RuleConfigs rules.RuleConfigMap
}

// Validate let's you validate a passed in Makefile with the provided config
func Validate(makefile parser.Makefile, config Config) (ret rules.RuleViolationList) {

	rules := rules.GetRegisteredRules()

	for name, rule := range rules {
		ret = append(ret, rule.Rule(makefile, config.RuleConfigs[name])...)
	}

	return
}
