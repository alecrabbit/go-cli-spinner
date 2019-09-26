// Package spinner implements a colorful console spinner
package spinner

import (
    "container/ring"
    "fmt"
    "io"
    "regexp"
    "strings"
    // "strconv"
    "sync"
    "time"

    "github.com/mattn/go-colorable"
    "github.com/mattn/go-runewidth"

    "github.com/alecrabbit/go-cli-spinner/aux"
)

func init() {
    // Initialize here
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
    Interval           time.Duration // interval between spinner refreshes
    charSet            *ring.Ring    // charSet holds chosen character set
    colorSet           *ring.Ring    // colorSet holds chosen colorize set
    active             bool          // active holds the state of the spinner
    FinalMessage       string        // spinner final message, displayed after Stop()
    currentMessage     string        // string
    currentChar        string
    currentProgress    string         // string
    Reversed           bool           // flag, spin in the opposite direction
    colorLevel         ColorLevel     // writeCurrentFrame color level
    lock               *sync.RWMutex  // lock
    Writer             io.Writer      // to make testing better, exported so users have access
    stop               chan bool      // stop, channel to stop the spinner
    HideCursor         bool           // flag, hide cursor
    regExp             *regexp.Regexp // regExp instance
    currentFrame       string         // writeCurrentFrame string written to output
    currentFrameWidth  int
    previousFrameWidth int
    FormatMessage      string
    FormatFrames       string
    FormatProgress     string
    outputFormat       string
}

// New provides a pointer to an instance of Spinner
func New(t int, d time.Duration) *Spinner {
    charSet := CharSets[t]
    colors := aux.ColorSets[aux.C256Rainbow]
    k := len(charSet)
    u := len(colors)
    s := Spinner{
        Interval:       d,
        charSet:        ring.New(k),
        colorSet:       ring.New(u),
        lock:           &sync.RWMutex{},
        Writer:         colorable.NewColorableStderr(),
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
    // Initialize charSet
    for i := 0; i < k; i++ {
        s.charSet.Value = charSet[i]
        s.charSet = s.charSet.Next()
    }
    // Initialize colorSet
    for i := 0; i < u; i++ {
        s.colorSet.Value = fmt.Sprintf("\x1b[38;5;%vm%s\x1b[0m", colors[i], "%s")
        s.colorSet = s.colorSet.Next()
    }
    s.updateOutputFormat()

    return &s
}

// IsActive returns true if spinner is currently active
func (s *Spinner) IsActive() bool {
    return s.active
}

// Get writeCurrentFrame spinner frame
func (s *Spinner) getCurrentChar() string {
    // Note: external lock
    // Rotate Ring
    if s.Reversed {
        s.charSet = s.charSet.Prev() // Backward
    } else {
        s.charSet = s.charSet.Next() // Forward
    }
    // writeCurrentFrame frame
    return s.charSet.Value.(string)
}

// Colorize in string
func (s *Spinner) colorize(in string) string {
    // TODO: use colorizing callback here?
    // Note: external lock
    if s.colorLevel >= Color16 {
        // Rotate Ring
        s.colorSet = s.colorSet.Next()
        // apply colorSet format
        return fmt.Sprintf(s.colorSet.Value.(string), in)
    }
    return in
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
                s.updateCurrentFrame()
                s.assembleCurrentFrame()
                s.writeCurrentFrame()
                s.lock.Unlock()
            }
        }
    }()
}

func (s *Spinner) updateCurrentFrame() {
    // Note: external lock
    s.currentChar = s.colorize(s.getCurrentChar())
}

func (s *Spinner) assembleCurrentFrame() {
    // Note: external lock
    s.previousFrameWidth = s.currentFrameWidth

    sp := fmt.Sprintf(s.outputFormat, s.currentChar, s.currentMessage, s.currentProgress)
    s.currentFrameWidth = s.frameWidth(sp)

    // Add erase and move cursor back ansi sequences
    // debugStr := " [" +strconv.Itoa(s.previousFrameWidth-s.currentFrameWidth) + "] " + strconv.Itoa(s.previousFrameWidth) + " " + strconv.Itoa(s.currentFrameWidth) + " "
    debugStr := ""
    s.currentFrame = sp + s.eraseSequence(s.previousFrameWidth-s.currentFrameWidth) + s.moveBackSequence(s.currentFrameWidth) + debugStr
}

// Write writeCurrentFrame to output
func (s *Spinner) writeCurrentFrame() {
    // Note: external lock
    _, _ = fmt.Fprint(s.Writer, s.currentFrame)
    // _, _ = fmt.Fprint(s.Writer, s.debugReplace(s.currentFrame)+"\n")
}

func (s *Spinner) moveBackSequence(w int) string {
    return fmt.Sprintf("\x1b[%vD", w)
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

// remove all ansi codes from string
func (s *Spinner) strip(in string) string {
    return s.regExp.ReplaceAllString(in, "")
}

// replace all "\x1b" to `\e`
func (s *Spinner) debugReplace(in string) string {
    return strings.ReplaceAll(in, "\x1b", `\e`)
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
        _, _ = fmt.Fprint(s.Writer, s.eraseSequence(s.currentFrameWidth))
        // _, _ = fmt.Fprint(s.Writer, s.eraseSequence(s.currentFrameWidth)+s.moveBackSequence(s.currentFrameWidth))
    }
}

// Get frame width
func (s *Spinner) frameWidth(f string) int {
    return runewidth.StringWidth(s.strip(f))
}

// Current prints out writeCurrentFrame frame
func (s *Spinner) Current() {
    s.lock.Lock()
    s.writeCurrentFrame()
    s.lock.Unlock()
}

func (s *Spinner) eraseSequence(w int) string {
    if w < 0 {
        return ""
    }
    // return strings.Repeat("-", w)
    return fmt.Sprintf("\x1b[%vX", w)
}

// Message sets writeCurrentFrame spinner message
func (s *Spinner) Message(m string) {
    s.lock.Lock()
    defer s.lock.Unlock()
    s.currentMessage = m
    // s.assembleCurrentFrame()
}

// Progress sets writeCurrentFrame spinner currentProgress value 0..1
func (s *Spinner) Progress(p float32) {
    p = aux.Bounds(p)
    s.lock.Lock()
    defer s.lock.Unlock()
    switch {
    case p > 0:
        s.currentProgress = fmt.Sprintf("%.0f%%", p*float32(100))
    default:
        s.currentProgress = ""
    }
    // s.assembleCurrentFrame()
}
