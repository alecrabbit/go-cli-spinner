package spinner

import (
    "reflect"
    "testing"
    "time"
)

// TestNew verifies that the returned instance is of the proper type
func TestNew(t *testing.T) {
    for i := 0; i < len(CharSets); i++ {
        s := New(i, 100*time.Millisecond)
        if reflect.TypeOf(s).String() != "*spinner.Spinner" {
            t.Errorf("New returned incorrect type kind=%d", i)
        }
        if s.IsActive() != false {
	        t.Errorf("Expected new instance to be inactive (%d)", i)
        }
    }
}

/*
Benchmarks
*/

// BenchmarkNew runs a benchmark for the New() function
func BenchmarkNew(b *testing.B) {
    for n := 0; n < b.N; n++ {
        New(Arrows, 1*time.Second)
    }
}

// func BenchmarkNewStartStop(b *testing.B) {
//	for n := 0; n < b.N; n++ {
//		s := New(CharSets[1], 1*time.Second)
//		s.Start()
//		s.Stop()
//	}
// }
