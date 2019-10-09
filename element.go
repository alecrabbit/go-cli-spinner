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
    format         string     //
    spacer         string     //
    current        string     //
    currentWidth   int        //
    // previousWidth  int        //
    charSet        *ring.Ring //
    cFormat        *ring.Ring //
    reversed       bool       //
    emptyFormat    string     //
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
    }
    el.currentWidth = runewidth.StringWidth(fmt.Sprintf(el.format+el.spacer, el.current))
}

func newElement(c int, f, s string, cs ...interface{}) (*element, error) {
    el := element{
        format:         f, //
        spacer:         s, //
    }
    el.cFormat = createColorSet(color.Prototypes[c], el.format+el.spacer)
    if cs != nil {
        if v, ok := cs[0].([]string); ok {
            el.charSet = applyCharSet(v)
            if el.charSet != nil {
                el.currentWidth =
                    runewidth.StringWidth(el.charSet.Value.(string)) +
                        runewidth.StringWidth(el.spacer) +
                        runewidth.StringWidth(fmt.Sprintf(el.format, ""))
                // el.previousWidth = el.currentWidth
            }
        } else {
            return nil, errors.New("spinner.newElement: fourth param expected to be type of []string")
        }
    }
    return &el, nil
}

// Colorize char
func (el *element) colorized(s string) string {
    // Note: external lock
    if el.cFormat != nil {
        // rotate
        el.cFormat = el.cFormat.Next()
        // apply
        return fmt.Sprintf(el.cFormat.Value.(string), s)
    }
    return s
}
