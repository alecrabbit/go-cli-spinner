// Package spinner implements a colorful console spinner
package spinner

import (
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/mattn/go-colorable"
	"github.com/mattn/go-runewidth"

	"github.com/alecrabbit/go-cli-spinner/auxiliary"
	"github.com/alecrabbit/go-cli-spinner/color"
)

// Spinner struct representing spinner instance
type Spinner struct {
	elements           map[int]*element         //
	elementsSettings   map[int]*elementSettings //
	elementsOrder      []int                    //
	char               *element                 //
	message            *element                 //
	progress           *element                 //
	charSettings       *elementSettings         //
	messageSettings    *elementSettings         //
	progressSettings   *elementSettings         //
	l                  *sync.RWMutex            // lock
	active             bool                     // flag, spinner is active
	colorLevel         color.Level              // holds color level
	stop               chan bool                // channel to stop the spinner
	outputFormat       string                   // output format string
	currentFrame       string                   // current frame string to write to output
	currentFrameWidth  int                      // width of currentFrame string
	previousFrameWidth int                      // previous width of currentFrame string
	interval           time.Duration            // interval between spinner refreshes
	finalMessage       string                   // spinner final message, displayed by calling Stop()
	reversed           bool                     // flag, spin in the opposite direction
	hideCursor         bool                     // flag, hide cursor
	prefix             string                   // spinner prefix
	prefixWidth        int                      // width of prefix string
	Writer             io.Writer                //
	maxMessageWidth    int                      //
	messageEllipsis    string                   //
	palette            *palette                 //
}

// New provides a pointer to an instance of Spinner
func New(options ...Option) (*Spinner, error) {
	charSet := CharSets[Snake2]
	s := Spinner{
		interval:        charSet.interval,
		palette:         charSet.palette,
		l:               &sync.RWMutex{},
		colorLevel:      color.TColor256,
		outputFormat:    "%s%s%s%s",
		stop:            make(chan bool),
		finalMessage:    "",
		hideCursor:      true,
		Writer:          colorable.NewColorableStderr(),
		elementsOrder:   []int{Char, Progress, Message},
		maxMessageWidth: 50,
		messageEllipsis: "…",
	}
	// Default settings for spinner elements
	s.charSettings = &elementSettings{
		colorizingSet: (*s.palette)[Char][s.colorLevel],
		format:        "%s",
		spacer:        " ",
		charSet:       charSet.chars,
	}
	s.messageSettings = &elementSettings{
		colorizingSet: (*s.palette)[Message][s.colorLevel],
		format:        "%s",
		spacer:        " ",
	}
	s.progressSettings = &elementSettings{
		colorizingSet: (*s.palette)[Progress][s.colorLevel],
		format:        "%s",
		auxFormat:     "%.0f%%",
		spacer:        " ",
	}
	s.elementsSettings = map[int]*elementSettings{
		Char:     s.charSettings,
		Message:  s.messageSettings,
		Progress: s.progressSettings,
	}
	// Process provided options
	for _, option := range options {
		if err := option(&s); err != nil {
			return nil, err
		}
	}

	// for _, entry := range s.elementsSettings {
	// 	if color.Prototypes[entry.colorizingSet].Level > s.colorLevel {
	// 		entry.colorizingSet = color.CNoColor
	// 	}
	// }
	// Create spinner elements
	if err := s.createElements(); err != nil {
		return nil, err
	}
	// Check interval
	if err := checkInterval(s.interval); err != nil {
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
	s.l.Lock()
	defer s.l.Unlock()
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
	m = auxiliary.Truncate(m, s.maxMessageWidth, s.messageEllipsis)
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
	return runewidth.StringWidth(auxiliary.StripANSI(f))
}

// write string by Writer
func (s *Spinner) write(v string) {
	// Note: external lock

	// Suppressed returns
	_, _ = fmt.Fprint(s.Writer, v)
}
