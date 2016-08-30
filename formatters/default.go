package formatters

import (
	"github.com/mrtazz/checkmake/rules"
	"github.com/olekukonko/tablewriter"
	"io"
	"os"
	"strconv"
)

// DefaultFormatter is the formatter used by default for CLI output
type DefaultFormatter struct {
	out io.Writer
}

// NewDefaultFormatter returns a DefaultFormatter struct
func NewDefaultFormatter() *DefaultFormatter {
	return &DefaultFormatter{out: os.Stdout}
}

// Format is the function to call to get the formatted output
func (f *DefaultFormatter) Format(violations rules.RuleViolationList) {

	data := make([][]string, len(violations))

	for idx, val := range violations {
		data[idx] = []string{val.Rule,
			val.Violation,
			strconv.Itoa(val.LineNumber)}
	}

	table := tablewriter.NewWriter(f.out)

	table.SetHeader([]string{"Rule", "Description", "Line Number"})

	table.SetCenterSeparator(" ")
	table.SetColumnSeparator(" ")
	table.SetRowSeparator(" ")
	table.SetAutoWrapText(true)

	table.AppendBulk(data)
	table.Render()
}
