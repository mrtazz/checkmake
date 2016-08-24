// Package parser implements all the parser functionality for Makefiles
package parser

import (
	"errors"
	"strings"
)

// Makefile provides a data structure to describe a parsed Makefile
type Makefile struct {
	Rules     RuleList
	Variables VariableList
}

// Rule represents a Make rule
type Rule struct {
	Target       string
	Dependencies []string
	Body         []string
}

// RuleList represents a list of rules
type RuleList []Rule

// Variable represents a Make variable
type Variable struct {
	Name           string
	SimplyExpanded bool
	Assignment     string
}

// VariableList represents a list of variables
type VariableList []Variable

// Parse is the main function to parse a Makefile from a file path string to a
// Makefile struct
func Parse(filepath string) (ret Makefile, err error) {

	var scanner *MakefileScanner
	scanner, err = NewMakefileScanner(filepath)
	if err != nil {
		return ret, err
	}

	for scanner.Scan() {
		switch {
		case strings.HasPrefix(scanner.Text(), "\t"):
			// parse rule body here
		case strings.HasPrefix(scanner.Text(), "#"):
			// parse comments here, ignoring them for now
			break
		default:
			// parse target or variable here
			ruleOrVariable, parseError := parseRuleOrVariable(scanner)
			if parseError != nil {
				return ret, parseError
			}
			switch ruleOrVariable.(type) {
			case Rule:
				rule, found := ruleOrVariable.(Rule)
				if found != true {
					return ret, errors.New("Parse error")
				}
				ret.Rules = append(ret.Rules, rule)
			case Variable:
				variable, found := ruleOrVariable.(Variable)
				if found != true {
					return ret, errors.New("Parse error")
				}
				ret.Variables = append(ret.Variables, variable)
			}
		}
	}

	return ret, err
}

// parseRuleOrVariable gets the parsing scanner in a state where it resides on
// a line that could be a variable or a rule. The function parses the line and
// subsequent lines if there is a rule body to parse and returns an interface
// that is either a Variable or a Rule struct and leaves the scanner in a
// state where it resides on the first line after the content parsed into the
// returned struct.
func parseRuleOrVariable(scanner *MakefileScanner) (ret interface{}, err error) {
	return
}
