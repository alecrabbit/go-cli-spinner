package spinner

import (
    "container/ring"
    "errors"
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

func newElement(c int, f, s string, cs ...interface{}) (*element, error) {
    el := element{
        format:         f, //
        spacer:         s, //
    }
    el.colorFormat = createColorSet(color.Prototypes[c], el.format+el.spacer)
    if cs != nil {
        if v, ok := cs[0].([]string); ok {
            el.charSet = applyCharSet(v)
            if el.charSet != nil {
                el.currentWidth =
                    runewidth.StringWidth(el.charSet.Value.(string)) +
                        runewidth.StringWidth(fmt.Sprintf(el.format, el.spacer))
            }
        } else {
            return nil, errors.New("spinner.newElement: fourth param expected to be type of []string")
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
