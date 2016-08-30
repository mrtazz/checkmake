package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"github.com/mrtazz/checkmake/config"
	"github.com/mrtazz/checkmake/formatters"
	"github.com/mrtazz/checkmake/logger"
	"github.com/mrtazz/checkmake/parser"
	"github.com/mrtazz/checkmake/rules"
	"github.com/mrtazz/checkmake/validator"
	"log"
	"os"
)

var (
	usage = `checkmake.

  Usage:
  checkmake [--debug|--config=<configPath>] <makefile>
  checkmake -h | --help
  checkmake --version

  Options:
  -h --help               Show this screen.
  --version               Show version.
  --debug                 Enable debug mode
  --config=<configPath>   Configuration file to read
`

	version   = ""
	buildTime = ""
	builder   = ""
	goversion = ""

	configPath = "checkmake.ini"
)

func main() {

	args, err := docopt.Parse(usage, nil, true,
		fmt.Sprintf("%s %s built at %s by %s with %s",
			"checkmake", version, buildTime, builder, goversion), false)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	formatter, violations := parseArgsAndGetFormatter(args)

	formatter.Format(violations)

	os.Exit(len(violations))
}

func parseArgsAndGetFormatter(args map[string]interface{}) (formatters.Formatter,
	rules.RuleViolationList) {
	if args["--debug"] == true {
		logger.SetLogLevel(logger.DebugLevel)
	}

	makefile, parseError := parser.Parse(args["<makefile>"].(string))

	if parseError != nil {
		log.Fatal(parseError)
		os.Exit(1)
	}

	if args["--config"] != nil {
		configPath = args["--config"].(string)
	}

	cfg, cfgError := config.NewConfigFromFile(configPath)

	if cfgError != nil {
		logger.Info(fmt.Sprintf("Unable to parse config file %q, running with defaults",
			configPath))
	}

	violations := validator.Validate(makefile, cfg)

	formatter := formatters.NewDefaultFormatter()

	return formatter, violations
}
