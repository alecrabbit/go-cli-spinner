package spinner

import (
    "time"

    "github.com/alecrabbit/go-cli-spinner/color"
)

// Option ...
type Option func(*Spinner) error

// ColorLevel sets color level support for spinner - TNoColor, TColor16, TColor256, TTrueColor
func ColorLevel(cl color.SupportLevel) Option {
    return func(s *Spinner) error {
        // TODO: check for correct value
        s.colorLevel = cl
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

