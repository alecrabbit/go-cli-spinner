package spinner

import (
	"reflect"
	"runtime"
	"testing"

	"github.com/alecrabbit/go-cli-spinner/aux"
)

// TestNew verifies that the returned instance is of the proper type
func TestNew(t *testing.T) {
	for i := 0; i < len(CharSets); i++ {
		s, _ := New()
		tp := reflect.TypeOf(s).String()
		if tp != "*spinner.Spinner" {
			t.Errorf("New returned incorrect type kind=%d %v", i, tp)
		}
		if s.IsActive() != false {
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
		if runtime.GOOS != aux.WINDOWS {
			d = 0
		}
	}
	result = d
}

// BenchmarkIfTwo ...
func BenchmarkIfTwo(b *testing.B) {
	var d int
	notWindows := runtime.GOOS != aux.WINDOWS
	for n := 0; n < b.N; n++ {
		if notWindows {
			d = 0
		}
	}
	result = d
}

// func BenchmarkNewStartStop(b *testing.B) {
//	for n := 0; n < b.N; n++ {
//		s := New(CharSets[1], 1*time.Second)
//		s.Start()
//		s.Stop()
//	}
// }
