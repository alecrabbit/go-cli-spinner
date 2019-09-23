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
    // Color16 represents 16 color level support
    Color16 ColorLevel = 1 << (4 * iota)
    // Color256 represents 256 color level support
    Color256
    // TrueColor represents true color level support
    TrueColor
)

// Spinner struct representing spinner instance
type Spinner struct {
    Interval       time.Duration // interval between spinner refreshes
    frames         *ring.Ring    // frames holds chosen character set
    colors         *ring.Ring    // colors holds chosen colorize set
    active         bool          // active holds the state of the spinner
    FinalMessage   string        // spinner final message, displayed after Stop()
    currentMessage string        // string
    FormatMessage  string
    FormatFrames   string
    FormatProgress string
    outputFormat   string
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
        FormatMessage:  "%s ",
        FormatFrames:   "%s ",
        FormatProgress: "%s ",
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
    s.updateOutputFormat()
    // Override os specific settings
    s = platformOverrides(s)

    return &s
}

// Method to override settings
func platformOverrides(s Spinner) Spinner {
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
    // Note: external lock
    // Rotate Ring
    if s.Reversed {
        s.frames = s.frames.Prev() // Backward
    } else {
        s.frames = s.frames.Next() // Forward
    }
    frame := s.frames.Value.(string)
    // format := s.createFormat(frame)
    return fmt.Sprintf(s.outputFormat, s.colorize(frame), s.currentMessage, s.progress)
    // return fmt.Sprintf(format, s.colorize(frame), s.currentMessage, s.progress)
}

func (s *Spinner) createFormat(frame string) string {
    // Note: external lock
    var format string
    if frame != "" {
        format += s.FormatFrames
    }
    if s.currentMessage != "" {
        format += s.FormatMessage
    }
    if s.progress != "" {
        format += s.FormatProgress
    }
    return format
}

// Colorize in string
func (s *Spinner) colorize(in string) string {
    // Note: external lock
    if s.colorLevel >= Color256 {
        // Rotate Ring
        s.colors = s.colors.Next()
        // Sprintf accordignly to colors format
        return fmt.Sprintf(s.colors.Value.(string), in)
    }
    return in
}

// Start will start the indicator
func (s *Spinner) Start() {
    s.lock.Lock()
    if s.active {
        // If already active - do nothing
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

// Get current frame and assign it to lastOutput
func (s *Spinner) updateLastOutput() {
    // Note: external lock
    frame := s.getFrame()
    // Add move cursor back ansi sequence
    s.lastOutput = frame + fmt.Sprintf("\x1b[%vD", s.frameWidth(frame))
}

func (s *Spinner) updateOutputFormat() {
    // Note: external lock
    s.outputFormat = fmt.Sprintf("%s%s%s", s.FormatFrames, s.FormatMessage, s.FormatProgress)
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

// Write erasing sequence to output
func (s *Spinner) erase() {
    // Note: external lock
    if s.active {
        _, _ = fmt.Fprint(s.Writer, fmt.Sprintf("\x1b[%vX", s.frameWidth(s.lastOutput)))
    }
}

// Get frame width
func (s *Spinner) frameWidth(f string) int {
    return runewidth.StringWidth(s.strip(f))
}

// Last prints out last spinner output
func (s *Spinner) Last() {
    s.lock.Lock()
    s.last()
    s.lock.Unlock()
}

func (s *Spinner) last() {
    // Note: external lock
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
        // s.FormatProgress = ""
    }
    s.lock.Unlock()
}
