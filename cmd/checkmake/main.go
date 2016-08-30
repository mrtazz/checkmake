package main

import (
	"fmt"
	"io"
	"log"
	"os"

	docopt "github.com/docopt/docopt-go"
	"github.com/mrtazz/checkmake/formatters"
	"github.com/mrtazz/checkmake/logger"
	"github.com/mrtazz/checkmake/parser"
	"github.com/mrtazz/checkmake/rules"
	"github.com/mrtazz/checkmake/validator"
	"github.com/olekukonko/tablewriter"
)

var (
	usage = `checkmake.

  Usage:
  checkmake [--debug] <makefile>
  checkmake -h | --help
  checkmake --version
  checkmake --list-rules

  Options:
  -h --help     Show this screen.
  --version     Show version.
  --debug       Enable debug mode
  --list-rules  List registered rules
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

	if args["--list-rules"] == true {
		listRules(os.Stdout)
		os.Exit(0)
	}

	makefile, parseError := parser.Parse(args["<makefile>"].(string))

	if parseError != nil {
		log.Fatal(parseError)
		os.Exit(1)
	}

	violations := validator.Validate(makefile, validator.Config{})

	formatter := formatters.NewDefaultFormatter()

	formatter.Format(violations)

	os.Exit(len(violations))
}

func listRules(w io.Writer) {
	data := [][]string{}
	for _, rule := range rules.GetRegisteredRules() {
		data = append(data, []string{rule.Name(), rule.Description()})
	}

	table := tablewriter.NewWriter(w)
	table.SetHeader([]string{"Name", "Description"})
	table.SetCenterSeparator(" ")
	table.SetColumnSeparator(" ")
	table.SetRowSeparator(" ")
	table.SetAutoWrapText(true)

	table.AppendBulk(data)
	table.Render()
}
