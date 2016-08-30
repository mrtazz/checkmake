// Package formatters provides the base interface type for different output
// formatters to implement
package formatters

import (
	"github.com/mrtazz/checkmake/rules"
)

// Formatter is the base interface type to implement for formatters
type Formatter interface {
	Format(violations rules.RuleViolationList)
}
