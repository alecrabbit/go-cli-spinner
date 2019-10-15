package spinner

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/mattn/go-runewidth"
)

// TestNew verifies that the returned instance is of the proper type
func TestNew(t *testing.T) {
	for i := 0; i < len(CharSets); i++ {
		s, _ := New()
		tp := reflect.TypeOf(s).String()
		if tp != "*spinner.Spinner" {
			t.Errorf("New returned incorrect type kind=%d %v", i, tp)
		}
		if s.Active() != false {
			t.Errorf("Expected new instance to be inactive (%d)", i)
		}
	}
}

/*
Benchmarks
*/
var result interface{}

// BenchmarkNew runs a benchmark for the New() function
func BenchmarkNew(b *testing.B) {
	var s *Spinner
	for n := 0; n < b.N; n++ {
		s, _ = New()
	}
	result = s
}

// BenchmarkIfOne ...
func BenchmarkIfOne(b *testing.B) {
	var d int
	for n := 0; n < b.N; n++ {
		d = runewidth.StringWidth(" ") +
			runewidth.StringWidth(fmt.Sprintf(" %s ", " "))
	}
	result = d
	// fmt.Printf("One %s", result)
}

// BenchmarkIfTwo ...
func BenchmarkIfTwo(b *testing.B) {
	var d int
	for n := 0; n < b.N; n++ {
		d = runewidth.StringWidth(" " + fmt.Sprintf(" %s ", " "))
	}
	result = d
	// fmt.Printf("Two %s", result)
}

// func BenchmarkNewStartStop(b *testing.B) {
//    for n := 0; n < b.N; n++ {
//        s := New(CharSets[1], 1*time.Second)
//        s.Start()
//        s.Stop()
//    }
// }
