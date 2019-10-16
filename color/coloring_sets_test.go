package color

import (
	"reflect"
	"testing"
)

// TestSets verifies that set can be used
func TestSets(t *testing.T) {
	for idx, sp := range Prototypes {
		r := sp.Handler(sp.ANSIStyles)
		tp := reflect.TypeOf(r).String()
		expected := "[]string"
		if tp != expected {
			t.Errorf("returned incorrect type for idx %v, expected: %v, given: %v", idx, expected, tp)
		}
	}
}
