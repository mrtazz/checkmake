package formatters

import (
	"bytes"
	"testing"
	"text/template"

	"github.com/mrtazz/checkmake/config"
	"github.com/mrtazz/checkmake/parser"
	"github.com/mrtazz/checkmake/validator"
	"github.com/stretchr/testify/assert"
)

func TestCustomFormatter(t *testing.T) {
	out := new(bytes.Buffer)

	tmpl, _ := template.New("test").Parse("{{.LineNumber}}:{{.Rule}}:{{.Violation}}")
	formatter := CustomFormatter{template: tmpl, out: out}

	makefile, _ := parser.Parse("../fixtures/missing_phony.make")

	violations := validator.Validate(makefile, &config.Config{})
	formatter.Format(violations)

	assert.Regexp(t, `0:minphony:Missing required phony target "all"`, out.String())
	assert.Regexp(t, `0:minphony:Missing required phony target "test"`, out.String())
	assert.Regexp(t, `18:phonydeclared:Target "all" should be declared PHONY.`, out.String())
}
