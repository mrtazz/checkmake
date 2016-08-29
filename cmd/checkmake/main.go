package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"github.com/mrtazz/checkmake/logger"
	"github.com/mrtazz/checkmake/parser"
	"github.com/mrtazz/checkmake/validator"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
	"strconv"
)

var (
	usage = `checkmake.

  Usage:
  checkmake [--debug] <makefile>
  checkmake -h | --help
  checkmake --version

  Options:
  -h --help     Show this screen.
  --version     Show version.
  --debug       Enable debug mode
`

	version   = ""
	buildTime = ""
	builder   = ""
	goversion = ""
)

func main() {

	args, err := docopt.Parse(usage, nil, true,
		fmt.Sprintf("%s %s built at %s by %s with %s",
			"checkmake", version, buildTime, builder, goversion), false)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if args["--debug"] == true {
		logger.SetLogLevel(logger.DebugLevel)
	}

	makefile, parseError := parser.Parse(args["<makefile>"].(string))

	if parseError != nil {
		log.Fatal(parseError)
		os.Exit(1)
	}

	violations := validator.Validate(makefile, validator.Config{})

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

	os.Exit(len(violations))
}
