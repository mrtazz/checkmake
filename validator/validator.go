// Package validator holds the basic engine that runs rule based checks
// against a Makefile struct
package validator

import (
	"fmt"

	"github.com/mrtazz/checkmake/config"
	"github.com/mrtazz/checkmake/logger"
	"github.com/mrtazz/checkmake/parser"
	"github.com/mrtazz/checkmake/rules"
	// rules register themselves via their package's init function, so we can
	// just blank import it
	_ "github.com/mrtazz/checkmake/rules/minphony"
	_ "github.com/mrtazz/checkmake/rules/phonydeclared"
)

// Validate let's you validate a passed in Makefile with the provided config
func Validate(makefile parser.Makefile, cfg *config.Config) (ret rules.RuleViolationList) {

	rules := rules.GetRegisteredRules()

	for name, rule := range rules {
		logger.Debug(fmt.Sprintf("Running rule '%s'...", name))
		ruleConfig := cfg.GetRuleConfig(name)
		if ruleConfig["disabled"] != "true" {
			ret = append(ret, rule.Run(makefile, ruleConfig)...)
		}
	}

	return
}
