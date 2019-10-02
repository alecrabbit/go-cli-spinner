package spinner

import (
	"container/ring"
	"fmt"
	"strings"

	"github.com/alecrabbit/go-cli-spinner/color"
)

// moveBackSequence returns string containing ANSI move cursor back sequence
func moveBackSequence(w int) string {
	return fmt.Sprintf("\x1b[%vD", w)
}

// eraseSequence returns string containing ANSI erase sequence
func eraseSequence(w int) string {
	if w < 0 {
		return ""
	}
	return fmt.Sprintf("\x1b[%vX", w)
}

// replace all "\x1b" to `\e`
func replaceEscapes(in string) string {
	return strings.ReplaceAll(in, "\x1b", `\e`)
}

func applyColorSetOld(cs color.Set) (r *ring.Ring) {
	u := len(cs.Set256)
	r = ring.New(u)
	for i := 0; i < u; i++ {
		r.Value = fmt.Sprintf("\x1b[38;5;%vm%s\x1b[0m", cs.Set256[i], "%s")
		r = r.Next()
	}
	return
}

<<<<<<< develop
func applyColorSet(p color.StylePrototype) (r *ring.Ring) {
	xs := p.Handler(p.ANSIStyles)
=======
func applyColorSet(xs []string) (r *ring.Ring) {
>>>>>>> change colorng model
	u := len(xs)
	r = ring.New(u)
	for i := 0; i < u; i++ {
		r.Value = xs[i]
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
