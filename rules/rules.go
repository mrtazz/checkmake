// Package rules contains specific rules as subpackages to check a Makefile against
package rules

import (
	"github.com/mrtazz/checkmake/parser"
)

// Rule is the type of a rule function
type Rule interface {
	Name() string
	Description() string
	Run(parser.Makefile, RuleConfig) RuleViolationList
}

// RuleViolation represents a basic validation failure
type RuleViolation struct {
	Rule       string
	Violation  string
	FileName   string
	LineNumber int
}

// RuleViolationList is a list of Violation types and the return type of a
// Rule function
type RuleViolationList []RuleViolation

// RuleConfig is a simple string/string map to hold key/value configuration
// for rules.
type RuleConfig map[string]string

// RuleConfigMap is a map that stores RuleConfig maps keyed by the rule name
type RuleConfigMap map[string]RuleConfig

// RuleRegistry is the type to hold rules keyed by their name
type RuleRegistry map[string]Rule

var (
	ruleRegistry RuleRegistry
)

func init() {
	ruleRegistry = make(RuleRegistry)
}

// RegisterRule let's you register a rule for inclusion in the validator
func RegisterRule(r Rule) {
	ruleRegistry[r.Name()] = r
}

// GetRegisteredRules returns the internal ruleRegistry
func GetRegisteredRules() RuleRegistry {
	return ruleRegistry
}
