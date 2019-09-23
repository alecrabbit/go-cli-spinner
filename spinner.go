// Package spinner implements a colorful console spinner
package spinner

import (
    "container/ring"
    "fmt"
    "io"
    "os"
    "regexp"
    "runtime"
    "sync"
    "time"

    "github.com/mattn/go-runewidth"

    "github.com/alecrabbit/go-cli-spinner/aux"
)

func init() {
    // Initialize here
    // fmt.Println("Init")
}

// ColorLevel represents color support level
type ColorLevel int

const (
    // NoColor no color support
    NoColor ColorLevel = iota
    // Color represents 16 color level support
    Color ColorLevel = 1 << (4 * iota)
    // Color256 represents 256 color level support
    Color256
    // TrueColor represents true color level support
    TrueColor
)

// Spinner struct representing spinner instance
type Spinner struct {
    Interval       time.Duration  // interval between spinner refreshes
    frames         *ring.Ring     // frames holds chosen character set
    colors         *ring.Ring     // colors holds chosen colorize set
    active         bool           // active holds the state of the spinner
    FinalMessage   string         // spinner final message, displayed after Stop()
    currentMessage string         // string
    progress       string         // string
    Reversed       bool           // flag, spin in the opposite direction
    colorLevel     ColorLevel     // current color level
    lock           *sync.RWMutex  // lock
    Writer         io.Writer      // to make testing better, exported so users have access
    stop           chan bool      // stop, channel to stop the spinner
    HideCursor     bool           // flag, hide cursor
    regExp         *regexp.Regexp // regExp instance
    lastOutput     string         // last string written to output
    // color      func(a ...interface{}) string // default color is white
    // enabled  bool          // active holds the state of the spinner
    // Prefix     string                        // Prefix is the text prepended to the indicator
    // Suffix     string                        // Suffix is the text appended to the indicator
}

// New provides a pointer to an instance of Spinner
func New(t int, d time.Duration) *Spinner {
    strings := CharSets[t]
    colors := aux.ColorsSets[aux.C256Rainbow]
    k := len(strings)
    u := len(colors)
    s := Spinner{
        Interval:       d,
        frames:         ring.New(k),
        colors:         ring.New(u),
        lock:           &sync.RWMutex{},
        Writer:         os.Stderr,
        colorLevel:     Color256,
        FinalMessage:   "",
        currentMessage: "",
        stop:           make(chan bool),
        HideCursor:     true,
        regExp:         regexp.MustCompile(`\x1b[[][^A-Za-z]*[A-Za-z]`),
    }
    // Initialize frames
    for i := 0; i < k; i++ {
        s.frames.Value = strings[i]
        s.frames = s.frames.Next()
    }
    // Initialize colors
    for i := 0; i < u; i++ {
        s.colors.Value = fmt.Sprintf("\x1b[38;5;%vm%s\x1b[0m", colors[i], "%s")
        s.colors = s.colors.Next()
    }

    // Override os specific settings
    s = specificSettings(s)

    return &s
}

func specificSettings(s Spinner) Spinner {
    // if s.Writer == os.Stderr {
    // 	s.colorLevel = NoColor
    // }
    if runtime.GOOS == aux.WINDOWS {
        s.HideCursor = false
    }
    return s
}

// IsActive returns true if spinner is currently active
func (s *Spinner) IsActive() bool {
    return s.active
}

// Get current frame
func (s *Spinner) getFrame() string {
    // Note: external lock is needed
    // Rotate Ring
    if s.Reversed {
        s.frames = s.frames.Prev() // Backward
    } else {
        s.frames = s.frames.Next() // Forward
    }
    return fmt.Sprintf("%s %s %s", s.colorize(s.frames.Value.(string)), s.currentMessage, s.progress)
}

// Colorize in string
func (s *Spinner) colorize(in string) string {
    // Rotate Ring
    s.colors = s.colors.Next()
    // Sprintf accordignly to colors format
    return fmt.Sprintf(s.colors.Value.(string), in)
}

// Start will start the indicator
func (s *Spinner) Start() {
    s.lock.Lock()
    if s.active {
        s.lock.Unlock()
        return
    }
    if s.HideCursor {
        // hide the cursor
        _, _ = fmt.Fprint(s.Writer, "\033[?25l")
    }

    s.active = true
    s.lock.Unlock()
    ticker := time.NewTicker(s.Interval)
    go func() {
        for {
            select {
            case <-s.stop:
                return
            case <-ticker.C:
                s.lock.Lock()
                s.updateLastOutput()
                s.last()
                s.lock.Unlock()
            }
        }
    }()
}

func (s *Spinner) updateLastOutput() {
    // Note: external lock is needed
    frame := s.getFrame()
    // Add move cursor back ansi sequence
    frame += fmt.Sprintf("\x1b[%vD", runewidth.StringWidth(frame))
    s.lastOutput = frame
}

// Stop stops the spinner
func (s *Spinner) Stop() {
    s.lock.Lock()
    defer s.lock.Unlock()
    if s.active {
        s.erase()
        s.active = false
        if s.HideCursor {
            // show the cursor
            _, _ = fmt.Fprint(s.Writer, "\033[?25h")
        }

        if s.FinalMessage != "" {
            _, _ = fmt.Fprint(s.Writer, s.FinalMessage)
        }
        s.stop <- true
    }
}

// remove all ansi codes from in string
func (s *Spinner) strip(in string) string {
    return s.regExp.ReplaceAllString(in, "")
}

// Erase erases spinner output
func (s *Spinner) Erase() {
    s.lock.Lock()
    s.erase()
    s.lock.Unlock()
}

func (s *Spinner) erase() {
    // Note: external lock is needed
    if s.active {
        _, err := fmt.Fprint(s.Writer, fmt.Sprintf("\x1b[%vX", runewidth.StringWidth(s.strip(s.lastOutput))))
        if err != nil {
            panic(err)
        }
    }
}

// Last prints out last spinner output
func (s *Spinner) Last() {
    s.lock.Lock()
    s.last()
    s.lock.Unlock()
}

func (s *Spinner) last() {
    // Note: external lock is needed
    _, _ = fmt.Fprint(s.Writer, s.lastOutput)
}

// Message sets current spinner message
func (s *Spinner) Message(m string) {
    s.lock.Lock()
    s.erase()
    s.currentMessage = m
    s.updateLastOutput()
    s.last()
    s.lock.Unlock()
}

// Progress sets current spinner progress value 0..1
func (s *Spinner) Progress(p float32) {
    p = aux.Bounds(p)
    r := fmt.Sprintf("%.0f%%", p*float32(100))
    s.lock.Lock()
    if p > 0 {
        s.progress = r
    } else {
        s.progress = ""
    }
    s.lock.Unlock()
}

