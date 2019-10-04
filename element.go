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
    colorPrototype int        //
    current        string     //
    currentWidth   int        //
    previousWidth  int        //
    charSet        *ring.Ring //
    colorSet       *ring.Ring //
    reversed       bool       //
}

func (el *element) getCurrent() string {
    if el.charSet != nil {
        if el.reversed {
            el.charSet = el.charSet.Prev()
        } else {
            el.charSet = el.charSet.Next()
        }
        el.current = el.charSet.Value.(string)
    }
    return el.current
}

func newElement(c int, f, s string, cs ...interface{}) (*element, error) {
    el := element{
        format:         f, //
        spacer:         s, //
        colorPrototype: c, //
    }
    el.colorSet = createColorSet(color.Prototypes[el.colorPrototype], el.format)
    if cs != nil {
        if v, ok := cs[0].([]string); ok {
            el.charSet = applyCharSet(v)
            if el.charSet != nil {
                el.currentWidth =
                    runewidth.StringWidth(el.charSet.Value.(string)) +
                        runewidth.StringWidth(el.spacer) +
                        runewidth.StringWidth(fmt.Sprintf(el.format, ""))
                el.previousWidth = el.currentWidth
            }
        } else {
            return nil, errors.New("spinner.newElement: fourth param expected to be type of []string")
        }
    }
    return &el, nil
}

// Colorize char
func (el *element) colorized(c string) string {
    // Note: external lock
    // rotate
    el.colorSet = el.colorSet.Next()
    // apply
    return fmt.Sprintf(el.colorSet.Value.(string), c)
}
