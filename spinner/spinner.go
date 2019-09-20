package spinner

import (
    "container/ring"
    "fmt"
    "io"
    "os"
    "regexp"
    "sync"
    "time"

    "github.com/mattn/go-runewidth"
)

func init() {
    fmt.Println("Init")
}

type ColorLevel int

const (
    NoColor            = iota
    Color   ColorLevel = 1 << (4 * iota)
    Color256
    Truecolor
)

// Spinner struct to hold the provided options
type Spinner struct {
    Interval   time.Duration // Delay is the speed of the indicator
    frames     *ring.Ring    // chars holds the chosen character set
    active     bool          // active holds the state of the spinner
    FinalMSG   string        // string displayed after Stop() is called
    currentMSG string        // string displayed after Stop() is called
    progress   float32       // string displayed after Stop() is called
    colorLevel ColorLevel
    lock       *sync.RWMutex //
    Writer     io.Writer     // to make testing better, exported so users have access
    stop       chan bool     // stopChan is a channel used to stop the indicator
    HideCursor bool          // hideCursor determines if the cursor is visible
    r          *regexp.Regexp
    lastOutput string // last character(set) written
    // color      func(a ...interface{}) string // default color is white
    // enabled  bool          // active holds the state of the spinner
    // Prefix     string                        // Prefix is the text preppended to the indicator
    // Suffix     string                        // Suffix is the text appended to the indicator
}

// New provides a pointer to an instance of Spinner with the supplied options
func New(t int, d time.Duration) *Spinner {
    strings := CharSets[t]
    k := len(strings)
    s := Spinner{
        Interval:   d,
        frames:     ring.New(k),
        lock:       &sync.RWMutex{},
        Writer:     os.Stderr,
        colorLevel: Color256,
        FinalMSG:   "",
        currentMSG: "",
        stop:       make(chan bool),
        HideCursor: true,
        r:          regexp.MustCompile(`\x1b[[][^A-Za-z]*[A-Za-z]`),
    }
    for i := 0; i < k; i++ {
        s.frames.Value = strings[i]
        s.frames = s.frames.Next()
    }
    return &s
}

// IsActive will return whether or not the spinner is currently active
func (s *Spinner) IsActive() bool {
    return s.active
}

func (s *Spinner) getFrame() string {
    s.frames = s.frames.Next()
    return s.frames.Value.(string) + " " + s.currentMSG
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
        _, _ = fmt.Fprintf(s.Writer, "\033[?25l")
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
                s.lastOutput = s.getFrame()
                s.lastOutput += fmt.Sprintf("\x1b[%vD", runewidth.StringWidth(s.lastOutput))
                // s.lastOutput = strings.ReplaceAll(s.lastOutput, "\x1b", `\e`) + "\n"
                _, _ = fmt.Fprintf(s.Writer, s.lastOutput)
                s.lock.Unlock()
            }
        }
    }()
}

// Stop stops the indicator
func (s *Spinner) Stop() {
    s.lock.Lock()
    defer s.lock.Unlock()
    if s.active {
        s.erase()
        s.active = false
        if s.HideCursor {
            // show the cursor
            _, _ = fmt.Fprintf(s.Writer, "\033[?25h")
        }

        if s.FinalMSG != "" {
            _, _ = fmt.Fprintf(s.Writer, s.FinalMSG)
        }
        s.stop <- true
    }
}

func (s *Spinner) strip(in string) string {
    return s.r.ReplaceAllString(in, "")
}

func (s *Spinner) Erase() {
    s.lock.Lock()
    s.erase()
    s.lock.Unlock()
}

func (s *Spinner) erase() {
    // Note: external lock is needed
    if s.active {
        _, _ = fmt.Fprintf(s.Writer, fmt.Sprintf("\x1b[%vX", runewidth.StringWidth(s.strip(s.lastOutput))))
    }
}

func (s *Spinner) Last() {
    s.lock.Lock()
    _, _ = fmt.Fprintf(s.Writer, s.lastOutput)
    s.lock.Unlock()
}

func (s *Spinner) Message(m string) {
    s.lock.Lock()
    s.currentMSG = m
    s.lock.Unlock()
}
