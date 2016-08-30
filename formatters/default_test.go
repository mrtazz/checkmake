package formatters

import (
	"bytes"
	"github.com/mrtazz/checkmake/config"
	"github.com/mrtazz/checkmake/parser"
	"github.com/mrtazz/checkmake/validator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultFormatter(t *testing.T) {
	out := new(bytes.Buffer)
	exp := "                                                                \n      RULE                 DESCRIPTION             LINE NUMBER  \n                                                                \n  phonydeclared   Target '\"all\"' should be         18           \n                  declared PHONY.                               \n                                                                \n"
	formatter := DefaultFormatter{out: out}

	makefile, _ := parser.Parse("../fixtures/missing_phony.make")

	violations := validator.Validate(makefile, &config.Config{})
	formatter.Format(violations)

	assert.Equal(t, exp, out.String())

}
