package auxiliary

import (
	"reflect"
	"testing"
)

func TestStripANSI(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"empty string",
			args{""},
			"",
		},
		{
			"string `text`",
			args{"\x1b[2mtext\x1b[0m"},
			"text",
		},
		{
			"string `text` 256 colors",
			args{"\x1b[38;5;214mtext\x1b[0m"},
			"text",
		},
		{
			"string `text` 256 colors and bg",
			args{"\x1b[38;5;214;48;5;161mtext\x1b[0m"},
			"text",
		},
		{
			"string `text` erase and mvBack",
			args{"text\x1b[1X\x1b[2D"},
			"text",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StripANSI(tt.args.in); got != tt.want {
				t.Errorf("StripANSI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBounds(t *testing.T) {
	type args struct {
		f float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			"with value `0`",
			args{0},
			0,
		},
		{
			"with value `-1`",
			args{-1},
			0,
		},
		{
			"with value `-10`",
			args{-10},
			0,
		},
		{
			"with value `2`",
			args{2},
			1,
		},
		{
			"with value `0.567`",
			args{0.567},
			0.567,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bounds(tt.args.f); got != tt.want {
				t.Errorf("Bounds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnique(t *testing.T) {
	type args struct {
		i []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"with value `[]`",
			args{[]int{}},
			[]int{},
		},
		{
			"with value `nil`",
			args{nil},
			nil,
		},
		{
			"with value `[1 2 3]`",
			args{[]int{1, 2, 3}},
			[]int{1, 2, 3},
		},
		{
			"with value `[1 2 2 3]`",
			args{[]int{1, 2, 2, 3}},
			[]int{1, 2, 3},
		},
		{
			"with value `[2 2]`",
			args{[]int{2, 2}},
			[]int{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// t.Run(fmt.Sprintf(tt.name, tt.args), func(t *testing.T) {
			if got := Unique(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEqual(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"with values `[], []`",
			args{[]int{}, []int{}},
			true,
		},
		{
			"with values `nil, []`",
			args{nil, []int{}},
			true,
		},
		{
			"with values `[], nil`",
			args{[]int{}, nil},
			true,
		},
		{
			"with values `[1], nil`",
			args{[]int{1}, nil},
			false,
		},
		{
			"with values `nil, [1]`",
			args{nil, []int{1}},
			false,
		},
		{
			"with values `[1], [1]`",
			args{[]int{1}, []int{1}},
			true,
		},
		{
			"with values `[1, 2], [1, 2]`",
			args{[]int{1, 2}, []int{1, 2}},
			true,
		},
		{
			"with values `[2, 1], [1, 2]`",
			args{[]int{2, 1}, []int{1, 2}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTruncate(t *testing.T) {
	type args struct {
		in string
		w  int
		l  interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"",
			args{"string", 4, ""},
			"stri",
		},
		{
			"",
			args{"string", 4, "..."},
			"stri...",
		},
		{
			"",
			args{"string", 4, nil},
			"stri…",
		},
		{
			"",
			args{"string", 4, "…"},
			"stri…",
		},
		{
			"",
			args{"string", 9, "…"},
			"string",
		},
		{
			"",
			args{"H㐀〾▓朗퐭텟şüöžåйкл¤〾▓朗", 4, nil},
			"H㐀〾▓…",
		},
		{
			"",
			args{"H㐀〾▓朗퐭텟şüöžåйкл¤〾▓朗", 7, ""},
			"H㐀〾▓朗퐭텟",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Truncate(tt.args.in, tt.args.w, tt.args.l); got != tt.want {
				t.Errorf("Truncate() = %v, want %v", got, tt.want)
			}
		})
	}
}