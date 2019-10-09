// Package spinner implements a colorful console spinner
package spinner

import (
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

// Spinner struct representing spinner instance
type Spinner struct {
    // jugglers           map[int]*juggler
    elements           map[int]*element   //
    char               *element           //
    message            *element           //
    progress           *element           //
    l                  *sync.RWMutex      // lock
    active             bool               // active holds the state of the spinner
    colorLevel         color.SupportLevel // writeCurrentFrame color level
    stop               chan bool          // stop, channel to stop the spinner
    regExp             *regexp.Regexp     // regExp instance
    outputFormat       string             // output format string e.g"%s %s %s"
    currentFrame       string             // current frame string to write to output
    currentFrameWidth  int                // width of currentFrame string
    previousFrameWidth int                // previous width of currentFrame string
    interval           time.Duration      // interval between spinner refreshes
    finalMessage       string             // spinner final message, displayed after Stop()
    reversed           bool               // flag, spin in the opposite direction
    hideCursor         bool               // flag, hide cursor
    prefix             string             // spinner prefix
    Writer             io.Writer          // to make testing better, exported so users have access
    prefixWidth        int
    charSettings       *elementSettings
    messageSettings    *elementSettings
    progressSettings   *elementSettings
}

// New provides a pointer to an instance of Spinner
func New(options ...Option) (*Spinner, error) {
    var err error
    s := Spinner{
        regExp:       regexp.MustCompile(`\x1b[[][^A-Za-z]*[A-Za-z]`),
        interval:     120 * time.Millisecond,
        l:            &sync.RWMutex{},
        colorLevel:   color.TColor256,
        outputFormat: "%s%s%s",
        stop:         make(chan bool),
        finalMessage: "",
        hideCursor:   true,
        Writer:       colorable.NewColorableStderr(),
    }
    s.charSettings = &elementSettings{
        colorizingSet: color.C256Rainbow,
        format:        "%s",
        spacer:        " ",
        charSet:       CharSets[Snake2],
    }
    s.messageSettings = &elementSettings{
        colorizingSet: color.C256YellowWhite,
        format:        "%s",
        spacer:        " ",
    }
    s.progressSettings = &elementSettings{
        colorizingSet: color.C256YellowWhite,
        format:        "%s",
        spacer:        " ",
    }
    // Process provided options
    for _, option := range options {
        err := option(&s)
        if err != nil {
            return nil, err
        }
    }
    s.char, err = newElement(
        s.charSettings.colorizingSet,
        s.charSettings.format,
        s.charSettings.spacer,
        s.charSettings.charSet,
    )
    if err != nil {
        return nil, err
    }

    s.message, err = newElement(
        s.messageSettings.colorizingSet,
        s.messageSettings.format,
        s.messageSettings.spacer,
        s.messageSettings.charSet,
    )
    if err != nil {
        return nil, err
    }

    s.progress, err = newElement(
        s.progressSettings.colorizingSet,
        s.progressSettings.format,
        s.progressSettings.spacer,
        s.progressSettings.charSet,
    )
    if err != nil {
        return nil, err
    }

    return &s, nil
}

// Active returns true if spinner is currently active
func (s *Spinner) Active() bool {
    return s.active
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
    go s.spin()
}

func (s *Spinner) spin() {
    ticker := time.NewTicker(s.interval)
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
    s.char.update()
}

func (s *Spinner) assembleCurrentFrame() {
    // Note: external lock
    s.previousFrameWidth = s.currentFrameWidth
    f := s.prefix + fmt.Sprintf(
        s.outputFormat,
        s.char.colorized(s.char.current),
        // s.char.current,
        s.message.colorized(s.message.current),
        s.progress.colorized(s.progress.current),
    )
    s.currentFrameWidth = s.prefixWidth + s.char.currentWidth + s.message.currentWidth + s.progress.currentWidth
    s.currentFrame = f + eraseSequence(s.previousFrameWidth-s.currentFrameWidth) + moveBackSequence(s.currentFrameWidth)
}

// Write writeCurrentFrame to spinner writer
func (s *Spinner) writeCurrentFrame() {
    // Note: external lock
    _, _ = fmt.Fprint(s.Writer, s.currentFrame)
    // _, _ = fmt.Fprint(s.Writer, replaceEscapes(s.currentFrame) + "\n")
}

// Stop stops the spinner
func (s *Spinner) Stop() {
    s.l.Lock()
    defer s.l.Unlock()
    if s.active {
        s.erase()
        s.active = false
        s.stop <- true
        if s.finalMessage != "" {
            _, _ = fmt.Fprint(s.Writer, s.finalMessage)
        }
        if s.hideCursor {
            // show the cursor
            _, _ = fmt.Fprint(s.Writer, "\033[?25h")
        }
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
    s.message.setCurrent(m)

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
    s.progress.setCurrent(r)
}

// frameWidth gets frame width
func (s *Spinner) frameWidth(f string) int {
    return runewidth.StringWidth(s.strip(f))
}
