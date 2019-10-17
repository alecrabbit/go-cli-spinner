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

const (
	// maxPrefixWidth spinner's max prefix width
	maxPrefixWidth = 10
	// minInterval
	minInterval = 50 * time.Millisecond
	// maxInterval
	maxInterval = 5 * time.Second
)

// Option type for functional options
type Option func(*Spinner) error

// ColorLevel sets color level support for spinner - TNoColor, TColor16, TColor256, TTrueColor
func ColorLevel(cl color.Level) Option {
	return func(s *Spinner) error {
		supported, ok := color.SupportedLevels[cl]
		if !ok {
			return fmt.Errorf("spinner: unknown color level: %v", cl)
		}
		if !supported {
			return fmt.Errorf("spinner: color level %v is not supported", cl)
		}
		s.colorLevel = cl
		return nil
	}
}

// Order sets spinner elements order
func Order(o ...int) Option {
	return func(s *Spinner) error {
		u := auxiliary.Unique(o)
		if len(u) != 3 {
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

// Variant sets spinner variant
func Variant(v int) Option {
	return func(s *Spinner) error {
		if _, ok := CharSets[v]; !ok {
			return fmt.Errorf("spinner: unknown variant, %v", v)
		}
		s.interval = CharSets[v].interval * time.Millisecond
		s.charSettings.charSet = CharSets[v].chars
		return nil
	}
}

// CharSet sets spinner char set
func CharSet(c []string) Option {
	return func(s *Spinner) error {
		err := checkCharSet(c)
		if err != nil {
			return err
		}
		s.charSettings.charSet = c
		return nil
	}
}

// Interval sets interval between spinner refreshes
func Interval(ms time.Duration) Option {
	return func(s *Spinner) error {
		interval := ms * time.Millisecond
		if interval < minInterval {
			return fmt.Errorf("spinner: interval is too small - %v, min=%v", interval, minInterval)
		}
		if interval > maxInterval {
			return fmt.Errorf("spinner: interval is too small - %v, min=%v", interval, minInterval)
		}
		s.interval = interval
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
		width := s.frameWidth(p)
		if width > maxPrefixWidth {
			return fmt.Errorf("spinner: prefix too long - %v", width)
		}
		s.prefix = p
		s.prefixWidth = width
		return nil
	}
}

// Reverse sets spinner's flag to rotate in reverse
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

// HideCursor sets spinner's hideCursor flag
func HideCursor(h bool) Option {
	return func(s *Spinner) error {
		s.hideCursor = h
		return nil
	}
}
