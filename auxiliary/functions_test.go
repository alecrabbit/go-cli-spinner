package auxiliary

import (
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