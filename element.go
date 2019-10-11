package spinner

import (
	"container/ring"
	"fmt"

	"github.com/mattn/go-runewidth"

	"github.com/alecrabbit/go-cli-spinner/color"
)

// element ...
type element struct {
	format       string     //
	spacer       string     //
	current      string     //
	currentWidth int        //
	charSet      *ring.Ring //
	colorFormat  *ring.Ring //
	reversed     bool       //
	emptyFormat  string     //
}

type elementSettings struct {
	colorizingSet int
	format        string
	spacer        string
	auxFormat     string
	charSet       []string
}

func (el *element) update() {
	if el.charSet != nil {
		if el.reversed {
			el.charSet = el.charSet.Prev()
		} else {
			el.charSet = el.charSet.Next()
		}
		el.current = el.charSet.Value.(string)
	}
}

func (el *element) setCurrent(s string) {
	el.current = s
	if s == "" {
		el.emptyFormat = ""
		el.currentWidth = 0
		return
	}
	el.currentWidth = runewidth.StringWidth(fmt.Sprintf(el.format+el.spacer, el.current))
}

func newElement(s *elementSettings) (*element, error) {
	el := element{
		format: s.format, //
		spacer: s.spacer, //
	}
	el.colorFormat = createColorSet(color.Prototypes[s.colorizingSet], el.format+el.spacer)
	if s.charSet != nil {
		el.charSet = applyCharSet(s.charSet)
		if el.charSet != nil {
			el.currentWidth =
				runewidth.StringWidth(el.charSet.Value.(string)) +
					runewidth.StringWidth(fmt.Sprintf(el.format, el.spacer))
		}
	}
	return &el, nil
}

// Colorize char
func (el *element) colorized() string {
	// Note: external lock
	if el.current == "" {
		return ""
	}
	if el.colorFormat != nil {
		// rotate
		el.colorFormat = el.colorFormat.Next()
		// apply
		return fmt.Sprintf(el.colorFormat.Value.(string), el.current)
	}
	return el.current
}
