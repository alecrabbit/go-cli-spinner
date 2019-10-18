package auxiliary

import (
	"regexp"

	"github.com/mattn/go-runewidth"
)

var regExp *regexp.Regexp // regExp instance

func init() {
	regExp = regexp.MustCompile(`\x1b[[][^A-Za-z]*[A-Za-z]`)
}

// StripANSI removes all ansi codes from in string
func StripANSI(in string) string {
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
	if i == nil {
		return i
	}
	k := make(map[int]bool)
	//noinspection GoPreferNilSlice
	l := []int{}
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

// Truncate sanitizes user names in an email
func Truncate(in string, w int, l ...interface{}) string {
	end := "…"
	if l != nil {
		if v, ok := l[0].(string);  ok {
			end = v
		}
	}
	if runewidth.StringWidth(StripANSI(in)) <= w {
		end = ""
	}
	result := in
	chars := 0
	for i := range in {
		if chars >= w {
			result = in[:i]
			break
		}
		chars++
	}
	return result + end
}