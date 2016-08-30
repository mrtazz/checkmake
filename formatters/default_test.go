package formatters

import (
	"bytes"
	"github.com/mrtazz/checkmake/parser"
	"github.com/mrtazz/checkmake/validator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultFormatter(t *testing.T) {
	out := new(bytes.Buffer)
	exp := "                                                        \n  RULE             DESCRIPTION             LINE NUMBER  \n                                                        \n  rule1   Target 'all' should be marked    18           \n          PHONY.                                        \n                                                        \n"
	formatter := DefaultFormatter{out: out}

	makefile, _ := parser.Parse("../fixtures/missing_phony.make")

	violations := validator.Validate(makefile, validator.Config{})
	formatter.Format(violations)

	assert.Equal(t, exp, out.String())

}
