package spinner

import (
	"testing"
)

var moveBackSequences = map[int]string{
	0:   "",
	-10: "",
	1:   "\x1b[1D",
	3:   "\x1b[3D",
	10:  "\x1b[10D",
}

// TestMoveBackSequence ...
func TestMoveBackSequence(t *testing.T) {
	for w, r := range moveBackSequences {
		sequence := moveBackSequence(w)
		if sequence != r {
			t.Errorf("moveBackSequence(%v) returned incorrect value", w)
		}
	}
}

var eraseSequences = map[int]string{
	0:   "",
	-10: "",
	1:   "\x1b[1X",
	3:   "\x1b[3X",
	10:  "\x1b[10X",
}

// TestEraseSequence ...
func TestEraseSequence(t *testing.T) {
	for w, r := range eraseSequences {
		sequence := eraseSequence(w)
		if sequence != r {
			t.Errorf("eraseSequence(%v) returned incorrect value", w)
		}
	}
}

type testedString struct {
	expected string
	given    string
}

var replaceEscapesData = map[int]testedString{
	// expected, given
	0: {"", ""},
	1: {`\ex1b`, "\x1bx1b"},
}

// TestEraseSequence ...
func TestReplaceEscapes(t *testing.T) {
	for idx, r := range replaceEscapesData {
		result := replaceEscapes(r.given)
		if result != r.expected {
			t.Errorf("replaceEscapes() returned incorrect value %v on idx=%v", result, idx)
		}
	}
}
