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

	"github.com/alecrabbit/go-cli-spinner/auxiliary"
	"github.com/alecrabbit/go-cli-spinner/color"
)

// Spinner struct representing spinner instance
type Spinner struct {
	elements           map[int]*element //
	elementsOrder      []int            //
	char               *element         //
	message            *element         //
	progress           *element         //
	l                  *sync.RWMutex    // lock
	active             bool             // active holds the state of the spinner
	colorLevel         color.Level      // writeCurrentFrame color level
	stop               chan bool        // stop, channel to stop the spinner
	regExp             *regexp.Regexp   // regExp instance
	outputFormat       string           // output format string e.g"%s %s %s"
	currentFrame       string           // current frame string to write to output
	currentFrameWidth  int              // width of currentFrame string
	previousFrameWidth int              // previous width of currentFrame string
	interval           time.Duration    // interval between spinner refreshes
	finalMessage       string           // spinner final message, displayed after Stop()
	reversed           bool             // flag, spin in the opposite direction
	hideCursor         bool             // flag, hide cursor
	prefix             string           // spinner prefix
	Writer             io.Writer        // to make testing better, exported so users have access
	prefixWidth        int              //
	charSettings       *elementSettings //
	messageSettings    *elementSettings //
	progressSettings   *elementSettings //
}

// New provides a pointer to an instance of Spinner
func New(options ...Option) (*Spinner, error) {
	var err error
	s := Spinner{
		regExp:        regexp.MustCompile(`\x1b[[][^A-Za-z]*[A-Za-z]`), // TODO move to auxiliary
		interval:      NewCharSets[Snake2].interval * time.Millisecond,
		l:             &sync.RWMutex{},
		colorLevel:    color.TColor256,
		outputFormat:  "%s%s%s%s",
		stop:          make(chan bool),
		finalMessage:  "",
		hideCursor:    true,
		Writer:        colorable.NewColorableStderr(),
		elementsOrder: []int{Char, Progress, Message},
	}
	// Default settings for spinner elements
	s.charSettings = &elementSettings{
		colorizingSet: color.C256Rainbow,
		format:        "%s",
		spacer:        " ",
		charSet:       NewCharSets[Snake2].chars,
	}
	s.messageSettings = &elementSettings{
		colorizingSet: color.CDark,
		format:        "%s",
		spacer:        " ",
	}
	s.progressSettings = &elementSettings{
		colorizingSet: color.C256YellowWhite,
		format:        "%s",
		auxFormat:     "%.0f%%",
		spacer:        " ",
	}
	// Process provided options
	for _, option := range options {
		err := option(&s)
		if err != nil {
			return nil, err
		}
	}
	// Create spinner elements
	err = s.createElements()
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (s *Spinner) createElements() error {
	var err error

	s.char, err = newElement(s.charSettings)
	if err != nil {
		return err
	}
	s.message, err = newElement(s.messageSettings)
	if err != nil {
		return err
	}
	s.progress, err = newElement(s.progressSettings)
	if err != nil {
		return err
	}
	s.elements = map[int]*element{
		Char:     s.char,
		Message:  s.message,
		Progress: s.progress,
	}
	return nil
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
		s.write("\033[?25l")
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
			s.write(s.currentFrame)
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
	first := s.elements[s.elementsOrder[0]]
	second := s.elements[s.elementsOrder[1]]
	third := s.elements[s.elementsOrder[2]]
	f := fmt.Sprintf(
		s.outputFormat,
		s.prefix,
		first.colorized(),
		second.colorized(),
		third.colorized(),
	)
	s.currentFrameWidth = s.prefixWidth + s.char.currentWidth + s.message.currentWidth + s.progress.currentWidth
	s.currentFrame = f + eraseSequence(s.previousFrameWidth-s.currentFrameWidth) + moveBackSequence(s.currentFrameWidth)
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
			s.write(s.finalMessage)
		}
		if s.hideCursor {
			// show the cursor
			s.write("\033[?25h")
		}
	}
}

// // remove all ansi codes from string
// func (s *Spinner) strip(in string) string { // TODO move to auxiliary
//     return s.regExp.ReplaceAllString(in, "")
// }
//
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
		s.write(eraseSequence(s.currentFrameWidth))
	}
}

// Current writes spinner current frame to output represented by spinner writer
func (s *Spinner) Current() {
	s.l.Lock()
	s.write(s.currentFrame)
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
	p = auxiliary.Bounds(p)
	var r string
	switch {
	case p > 0:
		r = fmt.Sprintf(s.progressSettings.auxFormat, p*float32(100))
	default:
		r = ""
	}
	s.l.Lock()
	defer s.l.Unlock()
	s.progress.setCurrent(r)
}

// frameWidth gets frame width
func (s *Spinner) frameWidth(f string) int {
	return runewidth.StringWidth(auxiliary.Strip(f))
}

// write string by Writer
func (s *Spinner) write(v string) {
	// Suppressed returns
	_, _ = fmt.Fprint(s.Writer, v)
}
