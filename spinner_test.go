package spinner

import (
	"reflect"
	"testing"

	"github.com/alecrabbit/go-cli-spinner/color"
)

// TestNew verifies that the returned instance is of the proper type
func TestNew(t *testing.T) {
	for i, cs := range CharSets {
		s, _ := New(
			Variant(i),
			Interval(cs.interval),
			CharSet(cs.chars),
			HideCursor(true),
			Reverse(),
			ColorLevel(color.TNoColor),
			Order(Char, Progress, Message),
			ProgressFormat("%6s"),
			ProgressIndicatorFormat("%.1f%%"),
			MessageFormat("(%s)"),
			Format("-%s -"),
			Prefix("\x1b[38;5;161m>>\x1b[0m"),
			FinalMessage("\x1b[38;5;34mDone!\x1b[0m\n"),
		)
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
// var result interface{}
//
// // BenchmarkNew runs a benchmark for the New() function
// func BenchmarkNew(b *testing.B) {
// 	var s *Spinner
// 	for n := 0; n < b.N; n++ {
// 		s, _ = New()
// 	}
// 	result = s
// }
//
// // BenchmarkIfOne ...
// func BenchmarkIfOne(b *testing.B) {
// 	var d int
// 	for n := 0; n < b.N; n++ {
// 		d = runewidth.StringWidth(" ") +
// 			runewidth.StringWidth(fmt.Sprintf(" %s ", " "))
// 	}
// 	result = d
// 	// fmt.Printf("One %s", result)
// }
//
// // BenchmarkIfTwo ...
// func BenchmarkIfTwo(b *testing.B) {
// 	var d int
// 	for n := 0; n < b.N; n++ {
// 		d = runewidth.StringWidth(" " + fmt.Sprintf(" %s ", " "))
// 	}
// 	result = d
// 	// fmt.Printf("Two %s", result)
// }
