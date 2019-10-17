package spinner

import (
	"bytes"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/alecrabbit/go-cli-spinner/color"
)

// syncBuffer ...
type syncBuffer struct {
	sync.Mutex
	bytes.Buffer
}

// Write ...
func (b *syncBuffer) Write(data []byte) (int, error) {
	b.Lock()
	defer b.Unlock()
	return b.Buffer.Write(data)
}

// TestNewOk verifies that the returned instance is of the proper type
func TestNewOk(t *testing.T) {
	for i, cs := range CharSets {
		s, err := New(
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
		if err != nil {
			t.Errorf("Unexpected error (%v) on set #%v", err, i)
			return
		}
		s.Writer = &syncBuffer{}
		tp := reflect.TypeOf(s).String()
		if tp != "*spinner.Spinner" {
			t.Errorf("New returned incorrect type kind=%d %v", i, tp)
			return
		}
		if s.Active() != false {
			t.Errorf("Expected new instance to be inactive (%d)", i)
		}

	}
}

// TestNewOk verifies that the returned instance is of the proper type
func TestDefaultRun(t *testing.T) {
	s, err := New()
	if err != nil {
		t.Errorf("Unexpected error (%v)", err)
		return
	}
	buffer := &syncBuffer{}
	s.Writer = buffer
	s.interval = 1 * time.Millisecond
	if s.Active() != false {
		t.Errorf("Expected spinner to be inactive")
	}
	s.Start()
	if s.Active() != true {
		t.Errorf("Expected spinner to be active")
	}
	s.Message("Message")
	s.Progress(0.1)
	time.Sleep(200 * time.Millisecond)
	s.Stop()
	time.Sleep(200 * time.Millisecond)
	if s.Active() != false {
		t.Errorf("Expected spinner to be inactive")
	}
}

// TestSets verifies that set can be used
func TestSets(t *testing.T) {
	for idx, sp := range color.Prototypes {
		r := sp.Handler(sp.ANSIStyles)
		tp := reflect.TypeOf(r).String()
		expected := "[]string"
		if tp != expected {
			t.Errorf("returned incorrect type for idx %v, expected: %v, given: %v", idx, expected, tp)
		}
	}
}

func TestNew(t *testing.T) {
	type args struct {
		option Option
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Prefix too long",
			args{Prefix("12345678901")},
			true,
		},
		{
			"Unknown color level",
			args{ColorLevel(13)},
			true,
		},
		{
			"Unsupported color level",
			args{ColorLevel(color.TTrueColor)},
			true,
		},
		{
			"Order three not unique",
			args{Order(Char, Char, Message)},
			true,
		},
		{
			"Order two not unique",
			args{Order(Char, Char)},
			true,
		},
		{
			"Order five not unique",
			args{Order(Char, Char, Message, Progress, 4)},
			true,
		},
		{
			"Order three wrong",
			args{Order(Message, Progress, 4)},
			true,
		},
		{
			"Unknown variant",
			args{Variant(12323)},
			true,
		},
		{
			"CharSet is too big",
			args{CharSet(returnBigCharSet(maxCharSetSize))},
			true,
		},
		{
			"Interval is too small",
			args{Interval(10 * time.Millisecond)},
			true,
		},
		{
			"Interval is too big",
			args{Interval(10 * time.Second)},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.option)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tp := reflect.TypeOf(got).String()
			if tp != "*spinner.Spinner" {
				t.Errorf("returns incorrect type kind %v", tp)
			}
		})
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

