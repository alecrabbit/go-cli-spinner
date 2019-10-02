package spinner

import (
	"time"

	"github.com/alecrabbit/go-cli-spinner/color"
)

// Deprecated
// type juggler struct {
// 	Format   string
// 	Spacer   string
// 	charColorSet *ring.Ring // charColorSet holds chosen colorizeChar set
// }

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

// Interval sets interval between spinner refreshes
func Interval(ms time.Duration) Option {
	return func(s *Spinner) error {
		// TODO: check for correct value
		s.interval = ms * time.Millisecond
		return nil
	}
}

// MessageFormat sets spinner message format
func MessageFormat(f string) Option {
	return func(s *Spinner) error {
		// TODO: check for correct value
		s.formatMessage = f
		return nil
	}
}

// ProgressFormat sets spinner progress indicator format
func ProgressFormat(f string) Option {
	return func(s *Spinner) error {
		// TODO: check for correct value
		s.formatProgress = f
		return nil
	}
}

// Format sets spinner format
func Format(f string) Option {
	return func(s *Spinner) error {
		// TODO: check for correct value
		s.formatChars = f
		return nil
	}
}
