package formatters

import (
	"bytes"
	"testing"

	"github.com/mrtazz/checkmake/parser"
	"github.com/mrtazz/checkmake/validator"
	"github.com/stretchr/testify/assert"
)

func TestDefaultFormatter(t *testing.T) {
	out := new(bytes.Buffer)
	formatter := DefaultFormatter{out: out}

	makefile, _ := parser.Parse("../fixtures/missing_phony.make")

	violations := validator.Validate(makefile, validator.Config{})
	formatter.Format(violations)

	assert.Regexp(t, `\s+RULE\s+DESCRIPTION\s+LINE NUMBER\s+`, out.String())
	assert.Regexp(t, `phonydeclared\s+Target "all" should be.+\s+18`, out.String())
	assert.Regexp(t, `\s+declared PHONY`, out.String())
}
