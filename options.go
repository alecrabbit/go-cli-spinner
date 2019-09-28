package spinner

import (
    "time"
)

// Option ...
type Option func(*Spinner) error

// ColorLevel sets color level support for spinner - NoColor, Color16, Color256, TrueColor
func ColorLevel(cl ColorSupportLevel) Option {
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

