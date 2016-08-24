# Parser

Checkmake includes a simple parser for Makefiles. The idea here is to build it
up over time and add features as required for validations. The base structure
returned by the parser is a struct that looks like this:

```go
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

```

Providing the most basic building blocks to run validations on.
