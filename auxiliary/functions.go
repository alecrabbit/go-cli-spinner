package auxiliary

import (
	"regexp"
)

var regExp *regexp.Regexp // regExp instance

func init() {
	regExp = regexp.MustCompile(`\x1b[[][^A-Za-z]*[A-Za-z]`)
}

// remove all ansi codes from string
func Strip(in string) string {
	return regExp.ReplaceAllString(in, "")
}

// Bounds restricts f value into bounds of 0..1
func Bounds(f float32) float32 {
	if f < 0 {
		f = 0
	}
	if f > 1 {
		f = 1
	}
	return f
}

// Unique returns slice with unique elements only
func Unique(i []int) []int {
	k := make(map[int]bool)
	var l []int
	for _, value := range i {
		if _, e := k[value]; !e {
			k[value] = true
			l = append(l, value)
		}
	}
	return l
}

// Equal compares two int slices and returns true if a and b are same length contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
