package parser

import (
	"testing"
)

func TestParseSimpleMakefile(t *testing.T) {

	_, err := Parse("../fixtures/simple.make")

	if err != nil {
		t.Errorf("Unable to parse 'fixtures/simple.make': %s", err.Error())
	}
}
