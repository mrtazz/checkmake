// Package validator holds the basic engine that runs rule based checks
// against a Makefile struct
package validator

import (
	"fmt"
	"github.com/mrtazz/checkmake/logger"
	"github.com/mrtazz/checkmake/parser"
	"github.com/mrtazz/checkmake/rules"
	// rules register themselves via their package's init function, so we can
	// just blank import it
	_ "github.com/mrtazz/checkmake/rules/phonydeclared"
)

// Config is a struct to configure the validator
type Config struct {
	RuleConfigs rules.RuleConfigMap
}

// Validate let's you validate a passed in Makefile with the provided config
func Validate(makefile parser.Makefile, config Config) (ret rules.RuleViolationList) {

	rules := rules.GetRegisteredRules()

	for name, rule := range rules {
		logger.Debug(fmt.Sprintf("Running rule '%s'...", name))
		ret = append(ret, rule.Run(makefile, config.RuleConfigs[name])...)
	}

	return
}
