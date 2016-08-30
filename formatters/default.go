package formatters

import (
	"github.com/mrtazz/checkmake/rules"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
)

// DefaultFormatter is the formatter used by default for CLI output
type DefaultFormatter struct {
}

// NewDefaultFormatter returns a DefaultFormatter struct
func NewDefaultFormatter() *DefaultFormatter {
	return &DefaultFormatter{}
}

// Format is the function to call to get the formatted output
func (f *DefaultFormatter) Format(violations rules.RuleViolationList) {

	data := make([][]string, len(violations))

	for idx, val := range violations {
		data[idx] = []string{val.Rule,
			val.Violation,
			strconv.Itoa(val.LineNumber)}
	}

	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"Rule", "Description", "Line Number"})

	table.SetCenterSeparator(" ")
	table.SetColumnSeparator(" ")
	table.SetRowSeparator(" ")
	table.SetAutoWrapText(true)

	table.AppendBulk(data)
	table.Render()
}
