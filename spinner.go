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

// type juggler struct {
// 	Format   string
// 	Spacer   string
// 	charColorSet *ring.Ring // charColorSet holds chosen colorize set
// }

// Spinner struct representing spinner instance
type Spinner struct {
    lock               *sync.RWMutex     // lock
    charSet            *ring.Ring        // charSet holds chosen character set
    charColorSet       *ring.Ring        // charColorSet holds chosen colorize set
    messageColorSet    *ring.Ring        // messageColorSet holds chosen colorize set
    progressColorSet   *ring.Ring        // progressColorSet holds chosen colorize set
    active             bool              // active holds the state of the spinner
    currentChar        string            // current spinner symbol
    currentMessage     string            // current message
    currentProgress    string            // current progress string
    colorLevel         ColorSupportLevel // writeCurrentFrame color level
    stop               chan bool         // stop, channel to stop the spinner
    regExp             *regexp.Regexp    // regExp instance
    outputFormat       string            // output format string e.g"%s %s %s"
    currentFrame       string            // writeCurrentFrame string to write to output
    currentFrameWidth  int
    previousFrameWidth int
    interval           time.Duration // interval between spinner refreshes
    FinalMessage       string        // spinner final message, displayed after Stop()
    Reversed           bool          // flag, spin in the opposite direction
    Writer             io.Writer     // to make testing better, exported so users have access
    HideCursor         bool          // flag, hide cursor
    FormatMessage      string
    FormatFrames       string
    FormatProgress     string
    Prefix             string
}

// New provides a pointer to an instance of Spinner
func New(options ...Option) (*Spinner, error) {
    s := Spinner{
        interval:       120 * time.Millisecond,
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
    s.charSet = applyCharSet(CharSets[Line])
    s.charColorSet = applyColorSet(ColorSet{Set256: ColorSets[C256Rainbow]})
    s.updateOutputFormat()

    // Process options
    for _, option := range options {
        err := option(&s)
        if err != nil {
            return nil, err
        }
    }
    return &s, nil
}

func applyColorSet(cs ColorSet) (r *ring.Ring) {
    u := len(cs.Set256)
    r = ring.New(u)
    for i := 0; i < u; i++ {
        r.Value = fmt.Sprintf("\x1b[38;5;%vm%s\x1b[0m", cs.Set256[i], "%s")
        r = r.Next()
    }
    return
}

func applyCharSet(charSet []string) (r *ring.Ring) {
    u := len(charSet)
    r = ring.New(u)
    for i := 0; i < u; i++ {
        r.Value = charSet[i]
        r = r.Next()
    }
    return
}

// IsActive returns true if spinner is currently active
func (s *Spinner) IsActive() bool {
    return s.active
}

type ColorSet struct {
    Set256 []int
}

// Colors ...
func (s *Spinner) Colors(cs ...ColorSet) {
    for i, c := range cs {
        switch i {
        case 0:
            // fmt.Println("Spinner ", c)
            s.charColorSet = applyColorSet(c)
        case 1:
            fmt.Println("Message ", c)
            s.messageColorSet = applyColorSet(c)
        case 2:
            fmt.Println("Progress", c)
            s.progressColorSet = applyColorSet(c)
        }
    }
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
    // Note: external lock
    if s.colorLevel >= Color16 {
        // Rotate Ring
        s.charColorSet = s.charColorSet.Next()
        // apply charColorSet format
        return fmt.Sprintf(s.charColorSet.Value.(string), in)
    }
    return in
}

// Start will start the spinner
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
    ticker := time.NewTicker(s.interval)
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
    preFrame := s.Prefix + fmt.Sprintf(s.outputFormat, s.currentChar, s.currentMessage, s.currentProgress)
    s.currentFrameWidth = s.frameWidth(preFrame)
    s.currentFrame = preFrame + s.eraseSequence(s.previousFrameWidth-s.currentFrameWidth) + s.moveBackSequence(s.currentFrameWidth)
}

// Write writeCurrentFrame to output
func (s *Spinner) writeCurrentFrame() {
    // Note: external lock
    _, _ = fmt.Fprint(s.Writer, s.currentFrame)
}

func (s *Spinner) moveBackSequence(w int) string {
    return fmt.Sprintf("\x1b[%vD", w)
}

func (s *Spinner) updateOutputFormat() {
    // Note: external lock
    s.outputFormat = fmt.Sprintf("%s%s%s",
        s.prepFmt(s.currentChar, s.FormatFrames),
        s.prepFmt(s.currentMessage, s.FormatMessage),
        s.prepFmt(s.currentProgress, s.FormatProgress),
    )
}

func (s *Spinner) prepFmt(f string, format string) string {
    if f == "" {
        return "%s"
    }
    return format
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

// erase writes erasing sequence to output
func (s *Spinner) erase() {
    // Note: external lock
    if s.active {
        _, _ = fmt.Fprint(s.Writer, s.eraseSequence(s.currentFrameWidth))
    }
}

// frameWidth gets frame f width
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
    return fmt.Sprintf("\x1b[%vX", w)
}

// Message sets writeCurrentFrame spinner message
func (s *Spinner) Message(m string) {
    s.lock.Lock()
    defer s.lock.Unlock()
    s.currentMessage = m
    s.updateOutputFormat()
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
    s.updateOutputFormat()
}
