package spinner

import (
	"container/ring"
	"fmt"
	"strings"

	"github.com/alecrabbit/go-cli-spinner/color"
)

// moveBackSequence returns string containing ANSI move cursor back sequence
func moveBackSequence(w int) string {
	if w <= 0 {
		return ""
	}
	return fmt.Sprintf("\x1b[%vD", w)
}

// eraseSequence returns string containing ANSI erase sequence
func eraseSequence(w int) string {
	if w < 1 {
		return ""
	}
	return fmt.Sprintf("\x1b[%vX", w)
}

// replace all "\x1b" to `\e`
func replaceEscapes(in string) string {
	return strings.ReplaceAll(in, "\x1b", `\e`)
}

func createColorSet(p color.StylePrototype, format string) (r *ring.Ring) {
	xs := p.Handler(p.ANSIStyles)
	u := len(xs)
	r = ring.New(u)
	for i := 0; i < u; i++ {
		r.Value = fmt.Sprintf(xs[i], format)
		r = r.Next()
	}
	return
}

func applyCharSet(charSet []string) (r *ring.Ring) {
	u := len(charSet)
	r = ring.New(u)
	for i := 0; i < u; i++ {
		r.Value = charSet[i]
		r = r.Next()
	}
	return
}
