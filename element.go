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
    cFormat        *ring.Ring //
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

func (el *element) setCurrent(s string) {
    el.current = s
    el.currentWidth = runewidth.StringWidth(el.current) +
        runewidth.StringWidth(fmt.Sprintf(el.format + el.spacer, ""))

}

func newElement(c int, f, s string, cs ...interface{}) (*element, error) {
    el := element{
        format:         f, //
        spacer:         s, //
        colorPrototype: c, //
    }
    el.cFormat = createColorSet(color.Prototypes[el.colorPrototype], el.format + el.spacer)
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
    fmt.Println(el.currentWidth)

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
