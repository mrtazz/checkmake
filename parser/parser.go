// Package parser implements all the parser functionality for Makefiles. This
// is supposed to be a parser with a very small feature set that just supports
// what is needed to do linting and checking and not actual full Makefile
// parsing. And it's handrolled because apparently GNU Make doesn't have a
// grammar (see http://www.mail-archive.com/help-make@gnu.org/msg02778.html)
package parser

import (
	"errors"
	"log"
	"regexp"
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

var (
	reFindRule             = regexp.MustCompile("^([a-zA-Z]+):(.*)")
	reFindSimpleVariable   = regexp.MustCompile("^([a-zA-Z]+) ?:=(.*)")
	reFindExpandedVariable = regexp.MustCompile("^([a-zA-Z]+) ?=(.*)")
)

// Parse is the main function to parse a Makefile from a file path string to a
// Makefile struct. This function should be kept fairly small and ideally most
// of the heavy lifting will live in the specific parsing functions below that
// know how to deal with individual lines.
func Parse(filepath string) (ret Makefile, err error) {

	var scanner *MakefileScanner
	scanner, err = NewMakefileScanner(filepath)
	if err != nil {
		return ret, err
	}

	for scanner.Scan() {
		switch {
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
// returned struct. The parsing of line details is done via regexing for now
// since it seems ok as a first pass but will likely have to change later into
// a proper lexer/parser setup.
func parseRuleOrVariable(scanner *MakefileScanner) (ret interface{}, err error) {
	line := scanner.Text()

	if matches := reFindRule.FindStringSubmatch(line); matches != nil {
		log.Printf("found rule in '%s' with matches %v\n", line, matches)
		ret = Rule{
			Target:       matches[1],
			Dependencies: strings.Split(matches[2], " ")}
	} else if matches := reFindSimpleVariable.FindStringSubmatch(line); matches != nil {
		ret = Variable{
			Name:           matches[1],
			Assignment:     matches[2],
			SimplyExpanded: true}
	} else if matches := reFindExpandedVariable.FindStringSubmatch(line); matches != nil {
		ret = Variable{
			Name:           matches[1],
			Assignment:     matches[2],
			SimplyExpanded: false}
	} else {
		log.Printf("didn't match '%s' with anything\n", line)
	}

	return
}
