package spinner

import (
    "fmt"
    "sort"
    "time"

    "github.com/alecrabbit/go-cli-spinner/auxiliary"
    "github.com/alecrabbit/go-cli-spinner/color"
)

const (
    // Char represents char element
    Char = 1 + iota
    // Message represents message element
    Message
    // Progress represents progress element
    Progress
)

// Option type for functional options
type Option func(*Spinner) error

// ColorLevel sets color level support for spinner - TNoColor, TColor16, TColor256, TTrueColor
func ColorLevel(cl color.SupportLevel) Option {
    return func(s *Spinner) error {
        // TODO: check for correct value
        s.colorLevel = cl
        return nil
    }
}

// Order sets spinner elements order
func Order(o ...int) Option {
    return func(s *Spinner) error {
        u := auxiliary.Unique(o)
        if len(u) < 3 {
            return fmt.Errorf("spinner: order option should contain three unique values, given: %v", o)
        }
        c := make([]int, len(u))
        copy(c, u)
        sort.Ints(c)
        if !auxiliary.Equal(c, []int{1, 2, 3}) {
            return fmt.Errorf(
                "spinner: order option should contain all elements identifiers %v, given: %v",
                []int{Char, Message, Progress},
                o)
        }
        s.elementsOrder = u
        return nil
    }
}

// Interval sets interval between spinner refreshes
func Interval(ms time.Duration) Option {
    return func(s *Spinner) error {
        // TODO: check for correct value
        s.interval = ms * time.Millisecond
        return nil
    }
}

// Variant sets spinner variant
func Variant(v int) Option {
    return func(s *Spinner) error {
        if _, ok := NewCharSets[v]; !ok {
            return fmt.Errorf("spinner: unknown variant, %v", v)
        }
        s.interval = NewCharSets[v].interval * time.Millisecond
        s.charSettings.charSet = NewCharSets[v].chars
        return nil
    }
}

// MessageFormat sets spinner message format
func MessageFormat(f string) Option {
    return func(s *Spinner) error {
        // TODO: check for correct value
        s.messageSettings.format = f
        return nil
    }
}

// ProgressFormat sets spinner progress indicator format
func ProgressFormat(f string) Option {
    return func(s *Spinner) error {
        // TODO: check for correct value
        s.progressSettings.format = f
        return nil
    }
}

// ProgressIndicatorFormat sets spinner progress indicator format
func ProgressIndicatorFormat(f string) Option {
    return func(s *Spinner) error {
        // TODO: check for correct value
        s.progressSettings.auxFormat = f
        return nil
    }
}

// Format sets spinner format
func Format(f string) Option {
    return func(s *Spinner) error {
        // TODO: check for correct value
        s.charSettings.format = f
        return nil
    }
}

// Prefix sets spinner prefix
func Prefix(p string) Option {
    return func(s *Spinner) error {
        // TODO: check for correct value
        s.prefix = p
        s.prefixWidth = s.frameWidth(p)
        return nil
    }
}

// Reverse sets spinner to rotate in reverse
func Reverse() Option {
    return func(s *Spinner) error {
        s.reversed = true
        return nil
    }
}

// FinalMessage sets spinner's final message
func FinalMessage(m string) Option {
    return func(s *Spinner) error {
        // TODO: check for correct value
        s.finalMessage = m
        return nil
    }
}

// HideCursor sets spinner's final message
func HideCursor(h bool) Option {
    return func(s *Spinner) error {
        s.hideCursor = h
        return nil
    }
}

// // CharSet sets spinner's final message
// func CharSet(cs []string) Option {
//     return func(s *Spinner) error {
//         s.charSet = cs
//         return nil
//     }
// }
