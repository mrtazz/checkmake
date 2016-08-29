package testhelpers

import (
	"reflect"
	"testing"
)

// Expect provides a simple way to write unit test assertions
// gratefully taken and adapted from
// http://keighl.com/post/mocking-http-responses-in-golang/
func Expect(t *testing.T, a interface{}, b interface{}) {
	if reflect.DeepEqual(a, b) == false {
		t.Errorf("Expected: %v (type %v) - Got: %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}
