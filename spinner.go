// Package spinner implements a colorful console spinner
package spinner

import (
    "container/ring"
    "errors"
    "fmt"
    "io"
    "regexp"
    "sync"
    "time"

    "github.com/mattn/go-colorable"
    "github.com/mattn/go-runewidth"

    "github.com/alecrabbit/go-cli-spinner/aux"
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
                el.currentWidth = charSetWidth(el.charSet)
                el.previousWidth = el.currentWidth
            }
        } else {
            return nil, errors.New("spinner.newElement: fourth param expected to be type of []string")
        }
    }
    return &el, nil
}

func charSetWidth(charSet *ring.Ring) int {
    return runewidth.StringWidth(charSet.Value.(string))
}

// Spinner struct representing spinner instance
type Spinner struct {
    char                   *element           //
    message                *element           //
    progress               *element           //
    formatMessage          string             // message format
    formatChars            string             // frames format
    formatProgress         string             // progress format
    l                      *sync.RWMutex      // lock
    charSet                *ring.Ring         // charSet holds chosen character set
    charColorSet           *ring.Ring         // charColorSet holds chosen colorizeChar set
    messageColorSet        *ring.Ring         // messageColorSet holds chosen colorizeChar set
    progressColorSet       *ring.Ring         // progressColorSet holds chosen colorizeChar set
    charColorPrototype     int                //
    messageColorPrototype  int                //
    progressColorPrototype int                //
    active                 bool               // active holds the state of the spinner
    currentChar            string             // current spinner symbol
    currentMessage         string             // current message
    currentProgress        string             // current progress string
    colorLevel             color.SupportLevel // writeCurrentFrame color level
    stop                   chan bool          // stop, channel to stop the spinner
    regExp             *regexp.Regexp         // regExp instance
    outputFormat       string                 // output format string e.g"%s %s %s"
    currentFrame       string                 // current frame string to write to output
    currentFrameWidth  int                    // width of currentFrame string
    previousFrameWidth int                    // previous width of currentFrame string
    interval           time.Duration          // interval between spinner refreshes
    finalMessage       string                 // spinner final message, displayed after Stop()
    reversed           bool                   // flag, spin in the opposite direction
    Writer             io.Writer              // to make testing better, exported so users have access
    hideCursor         bool                   // flag, hide cursor
    prefix             string                 // spinner prefix
}

// New provides a pointer to an instance of Spinner
func New(options ...Option) (*Spinner, error) {
    var err error
    s := Spinner{
        regExp:                 regexp.MustCompile(`\x1b[[][^A-Za-z]*[A-Za-z]`),
        interval:               120 * time.Millisecond,
        l:                      &sync.RWMutex{},
        colorLevel:             color.TColor256,
        charColorPrototype:     color.CDark,
        messageColorPrototype:  color.C256Rainbow,
        progressColorPrototype: color.CDark,
        formatMessage:          "%s ",
        formatChars:            "%s ",
        formatProgress:         "%s ",
        currentMessage:         "",
        currentProgress:        "",
        outputFormat:           "%s%s%s",
        stop:                   make(chan bool),
        finalMessage:           "",
        hideCursor:             true,
        Writer:                 colorable.NewColorableStderr(),
    }
    s.char, err = newElement(color.CDark, "%s", " ", CharSets[Snake2])
    if err != nil {
        return nil, err
    }
    s.message, err = newElement(color.CDark, "%s", " ")
    if err != nil {
        return nil, err
    }
    s.progress, err = newElement(color.CDark, "%s", " ")
    if err != nil {
        return nil, err
    }
    // Initialize default characters colorizing set
    s.charColorSet = createColorSet(color.Prototypes[s.charColorPrototype], s.formatChars)
    s.messageColorSet = createColorSet(color.Prototypes[s.messageColorPrototype], s.formatMessage)
    s.progressColorSet = createColorSet(color.Prototypes[s.progressColorPrototype], s.formatProgress)

    // Initialize default characters set
    s.charSet = applyCharSet(CharSets[Snake2])

    // Process provided options
    for _, option := range options {
        err := option(&s)
        if err != nil {
            return nil, err
        }
    }
    // Get first frame to correctly initialize output format
    s.updateCurrentFrame()
    // Initialize output format
    // s.updateOutputFormat()

    return &s, nil
}

// Active returns true if spinner is currently active
func (s *Spinner) Active() bool {
    return s.active
}

// Get writeCurrentFrame spinner frame
func (s *Spinner) getCurrentChar() string {
    // Note: external lock
    // Rotate Ring
    if s.reversed {
        s.charSet = s.charSet.Prev() // Backward
    } else {
        s.charSet = s.charSet.Next() // Forward
    }
    // writeCurrentFrame frame
    return s.charSet.Value.(string)
}

// Start will start the spinner
func (s *Spinner) Start() {
    s.l.Lock()
    if s.active {
        s.l.Unlock()
        return
    }
    if s.hideCursor {
        // hide the cursor
        _, _ = fmt.Fprint(s.Writer, "\033[?25l")
    }

    s.active = true
    s.l.Unlock()
    ticker := time.NewTicker(s.interval)
    go s.spin(ticker)
}

func (s *Spinner) spin(ticker *time.Ticker) {
    for {
        select {
        case <-s.stop:
            return
        case <-ticker.C:
            s.l.Lock()
            s.updateCurrentFrame()
            s.assembleCurrentFrame()
            s.writeCurrentFrame()
            s.l.Unlock()
        }
    }
}

func (s *Spinner) updateCurrentFrame() {
    // Note: external lock
    s.currentChar = s.colorizeChar(s.getCurrentChar())
}

func (s *Spinner) assembleCurrentFrame() {
    // Note: external lock
    s.previousFrameWidth = s.currentFrameWidth
    f := s.prefix + fmt.Sprintf(s.outputFormat, s.currentChar, s.currentMessage, s.currentProgress)
    s.currentFrameWidth = s.frameWidth(f)
    s.currentFrame = f + eraseSequence(s.previousFrameWidth-s.currentFrameWidth) + moveBackSequence(s.currentFrameWidth)
}

// Write writeCurrentFrame to spinner writer
func (s *Spinner) writeCurrentFrame() {
    // Note: external lock
    _, _ = fmt.Fprint(s.Writer, s.currentFrame)
}

func (s *Spinner) refineFormat(f string, format string) string {
    if s.strip(f) == "" {
        return "%s"
    }
    return format
}

// Stop stops the spinner
func (s *Spinner) Stop() {
    s.l.Lock()
    defer s.l.Unlock()
    if s.active {
        s.erase()
        s.active = false
        if s.hideCursor {
            // show the cursor
            _, _ = fmt.Fprint(s.Writer, "\033[?25h")
        }
        if s.finalMessage != "" {
            _, _ = fmt.Fprint(s.Writer, s.finalMessage)
        }
        s.stop <- true
    }
}

// remove all ansi codes from string
func (s *Spinner) strip(in string) string {
    return s.regExp.ReplaceAllString(in, "")
}

// Erase erases spinner output
func (s *Spinner) Erase() {
    s.l.Lock()
    s.erase()
    s.l.Unlock()
}

// erase writes erasing sequence to output
func (s *Spinner) erase() {
    // Note: external lock
    if s.active {
        _, _ = fmt.Fprint(s.Writer, eraseSequence(s.currentFrameWidth))
    }
}

// Current writes spinner current frame to output represented by spinner writer
func (s *Spinner) Current() {
    s.l.Lock()
    s.writeCurrentFrame()
    s.l.Unlock()
}

// Message sets spinner message
func (s *Spinner) Message(m string) {
    s.l.Lock()
    defer s.l.Unlock()
    if m != "" {
        s.currentMessage = s.colorizeMessage(fmt.Sprintf(s.formatMessage, m))
    } else {
        s.currentMessage = ""
    }
}

// Progress sets spinner progress value 0..1 → 0%..100%
func (s *Spinner) Progress(p float32) {
    p = aux.Bounds(p)
    var r string
    switch {
    case p > 0:
        r = fmt.Sprintf("%.0f%%", p*float32(100))
    default:
        r = ""
    }
    s.l.Lock()
    defer s.l.Unlock()
    if r != "" {
        s.currentProgress = s.colorizeProgress(fmt.Sprintf(s.formatProgress, r))
    } else {
        s.currentProgress = ""
    }
}

// frameWidth gets frame width
func (s *Spinner) frameWidth(f string) int {
    return runewidth.StringWidth(s.strip(f))
}

// Colorize char
func (s *Spinner) colorizeChar(c string) string {
    // Note: external lock
    if s.colorLevel > color.TNoColor {
        // Rotate Ring
        s.charColorSet = s.charColorSet.Next()
        // apply charColorSet format
        return fmt.Sprintf(s.charColorSet.Value.(string), c)
    }
    return c
}

// Colorize message
func (s *Spinner) colorizeMessage(m string) string {
    if s.colorLevel > color.TNoColor && s.messageColorSet != nil {
        s.messageColorSet = s.messageColorSet.Next()
        // apply charColorSet format
        return fmt.Sprintf(s.messageColorSet.Value.(string), m)
        // return fmt.Sprintf("\x1b[2m%s\x1b[0m", m) // Dark
    }
    return m
}

// Colorize progress
func (s *Spinner) colorizeProgress(p string) string {
    if s.colorLevel > color.TNoColor && s.progressColorSet != nil {
        s.progressColorSet = s.progressColorSet.Next()
        // apply charColorSet format
        return fmt.Sprintf(s.progressColorSet.Value.(string), p)
        // return fmt.Sprintf("\x1b[2m%s\x1b[0m", p) // Dark
    }
    return p
}
