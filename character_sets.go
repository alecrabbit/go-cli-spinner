package spinner

import (
	"fmt"
	"time"

	"github.com/mattn/go-runewidth"

	"github.com/alecrabbit/go-cli-spinner/color"
)

const (
	clockOneOClock = '\U0001F550'
	clockOneThirty = '\U0001F55C'
)

// maxCharSetSize maximum character set elements count
const maxCharSetSize = 60

// Declared spinner variants
const (
	BlockVertical int = iota
	// Arrows
	BouncingBlock
	Blink
	FlyingLine
	RotatingCircle
	Clock
	HalfClock
	HalfClock2
	Snake
	Snake2
	FlyingDots
	Dots10
	Dots13
	Dots14
	BlockHorizontal
	Toggle
	ToggleSmall
	Arrows01
	Arrows02
	Arrows03
	Arrows04
	Dots21
	Dots22
	Dots23
	Dots24
	Dots25
	Dots26
	Dev
	Dev2
	// Weather
	Simple
)

// Line is alias for Simple
const Line = Simple

type settings struct {
	interval time.Duration // interval between spinner refreshes
	chars    []string      //
	palette  *palette      //
}

// palette ...
type palette map[int]map[color.Level]int

// defaultPalette ...
var defaultPalette = palette{
	Char: {
		color.TTrueColor: color.C256Rainbow,
		color.TColor256:  color.C256Rainbow,
		color.TColor16:   color.CLightCyan,
	},
	Message: {
		color.TTrueColor: color.CDark,
		color.TColor256:  color.CDark,
		color.TColor16:   color.CDark,
	},
	Progress: {
		color.TTrueColor: color.C256YellowWhite,
		color.TColor256:  color.C256YellowWhite,
		color.TColor16:   color.CDark,
	},
}

// CharSets contains the available character sets
var CharSets = map[int]settings{
	// Arrows: {
	//     120,
	//     []string{"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"}, // Ambiguous width, issue in runewidth
	// },
	// ← 1
	// ↖ 2
	// ↑ 1
	// ↗ 2
	// → 1
	// ↘ 2
	// ↓ 1
	// ↙ 2

	Arrows01: {
		120 * time.Millisecond,
		[]string{"←", "↑", "→", "↓"},
		&defaultPalette,
	},
	Arrows02: {
		120 * time.Millisecond,
		[]string{"↖", "↗", "↘", "↙"},
		&defaultPalette,
	},
	Arrows03: {
		120 * time.Millisecond,
		[]string{"⇐", "⇖", "⇑", "⇗", "⇒", "⇘", "⇓", "⇙"},
		&defaultPalette,
	},
	Arrows04: {
		120 * time.Millisecond,
		[]string{"▹▹▹▹▹", "▸▹▹▹▹", "▹▸▹▹▹", "▹▹▸▹▹", "▹▹▹▸▹", "▹▹▹▹▸"},
		&defaultPalette,
	},
	Simple: {
		120 * time.Millisecond,
		[]string{"|", "\\", "─", "/"},
		&defaultPalette,
	},
	Dev: { // Singe character used for dev purposes
		400 * time.Millisecond,
		[]string{"+"},
		&defaultPalette,
	},
	Dev2: { // Number characters used for dev purposes
		250 * time.Millisecond,
		[]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"},
		&defaultPalette,
	},
	BlockVertical: {
		120 * time.Millisecond,
		[]string{"▁", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃", "▁"},
		&defaultPalette,
	},
	BlockHorizontal: {
		120 * time.Millisecond,
		[]string{"▉", "▊", "▋", "▌", "▍", "▎", "▏", "▎", "▍", "▌", "▋", "▊", "▉"},
		&defaultPalette,
	},
	BouncingBlock: {
		120 * time.Millisecond,
		[]string{"▖", "▘", "▝", "▗"},
		&defaultPalette,
	},
	RotatingCircle: {
		120 * time.Millisecond,
		[]string{"◐", "◓", "◑", "◒"},
		&defaultPalette,
	},
	Snake: {
		150 * time.Millisecond,
		[]string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"},
		&defaultPalette,
	},
	Snake2: {
		120 * time.Millisecond,
		[]string{"⠏", "⠛", "⠹", "⢸", "⣰", "⣤", "⣆", "⡇"},
		&defaultPalette,
	},
	FlyingDots: {
		120 * time.Millisecond,
		[]string{
			"⢀⠀", "⡀⠀", "⠄⠀", "⢂⠀", "⡂⠀", "⠅⠀", "⢃⠀", "⡃⠀", "⠍⠀", "⢋⠀", "⡋⠀", "⠍⠁", "⢋⠁", "⡋⠁", "⠍⠉", "⠋⠉", "⠋⠉",
			"⠉⠙", "⠉⠙", "⠉⠩", "⠈⢙", "⠈⡙", "⢈⠩", "⡀⢙", "⠄⡙", "⢂⠩", "⡂⢘", "⠅⡘", "⢃⠨", "⡃⢐", "⠍⡐", "⢋⠠", "⡋⢀", "⠍⡁",
			"⢋⠁", "⡋⠁", "⠍⠉", "⠋⠉", "⠋⠉", "⠉⠙", "⠉⠙", "⠉⠩", "⠈⢙", "⠈⡙", "⠈⠩", "⠀⢙", "⠀⡙", "⠀⠩", "⠀⢘", "⠀⡘", "⠀⠨",
			"⠀⢐", "⠀⡐", "⠀⠠", "⠀⢀", "⠀⡀",
		},
		&defaultPalette,
	},
	FlyingLine: {
		120 * time.Millisecond,
		[]string{"|   ", " |  ", "  | ", "   |", "   |", "  | ", " |  ", "|   "},
		&defaultPalette,
	},
	Dots10: {
		120 * time.Millisecond,
		[]string{"⢄", "⢂", "⢁", "⡁", "⡈", "⡐", "⡠"},
		&defaultPalette,
	},
	Dots13: {
		120 * time.Millisecond,
		[]string{"⠁", "⠂", "⠄", "⡀", "⢀", "⠠", "⠐", "⠈"},
		&defaultPalette,
	},
	Dots14: {
		120 * time.Millisecond,
		[]string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
		&defaultPalette,
	},
	Dots21: {
		120 * time.Millisecond,
		[]string{
			"⠁", "⠁", "⠉", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠤", "⠄", "⠄",
			"⠤", "⠠", "⠠", "⠤", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋", "⠉", "⠈", "⠈",
		},
		&defaultPalette,
	},
	Dots22: {
		120 * time.Millisecond,
		[]string{
			"⠈", "⠉", "⠋", "⠓", "⠒", "⠐", "⠐", "⠒", "⠖", "⠦", "⠤", "⠠",
			"⠠", "⠤", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋", "⠉", "⠈",
		},
		&defaultPalette,
	},
	Dots26: {
		120 * time.Millisecond,
		[]string{"⢹", "⢺", "⢼", "⣸", "⣇", "⡧", "⡗", "⡏"},
		&defaultPalette,
	},
	Blink: {
		200 * time.Millisecond,
		[]string{"▓", "▒", "░"},
		&defaultPalette,
	},
	Toggle: {
		250 * time.Millisecond,
		[]string{"■", "□"},
		&defaultPalette,
	},
	Dots23: {
		120 * time.Millisecond,
		[]string{
			"⠁", "⠉", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠤", "⠄",
			"⠄", "⠤", "⠴", "⠲", "⠒", "⠂", "⠂", "⠒", "⠚", "⠙", "⠉", "⠁",
		},
		&defaultPalette,
	},
	Dots24: {
		120 * time.Millisecond,
		[]string{".  ", ".. ", "...", " ..", "  .", "   "},
		&defaultPalette,
	},
	Dots25: {
		120 * time.Millisecond,
		[]string{"⠋", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋"},
		&defaultPalette,
	},
	// ToggleSmall: { // Incorrect width 2 instead of 1 (runewidth issue)
	//     250,
	//     []string{"▪", "▫"},
	// },

	// ********
	// // Toggle:          {"■", "□", "▪", "▫"}, // Ambiguous width
	// // ■ 1
	// // □ 1
	// // ▪ 2
	// // ▫ 2
	// // Weather: { // Ambiguous width
	// //     "🌤 ", "🌤 ", "🌤 ", "🌤 ", "⛅️", "🌥 ", "☁️ ", "🌧 ", "🌨 ", "🌧 ", "🌨 ", "🌧 ", "🌨 ", "⛈ ",
	// //     "⛈ ", "🌨 ", "⛈ ", "🌧 ", "🌨 ", "☁️ ", "🌥 ", "⛅️", "🌤 ",
	// // },
}

func init() {
	fillCharSets()

	// Check CharSets for width conformity
	for n := range CharSets {
		err := checkCharSet(CharSets[n].chars)
		if err != nil {
			panic(err)
		}
	}
}

func fillCharSets() {
	var clockChars []string
	var halfClockChars []string
	var halfClockChars2 []string
	// Fill clocks char sets
	for i := rune(0); i < 12; i++ {
		clockChars = append(clockChars, string([]rune{clockOneOClock + i}))
		halfClockChars = append(halfClockChars, string([]rune{clockOneOClock + i}), string([]rune{clockOneThirty + i}))
	}
	halfClockChars2 = make([]string, len(clockChars))
	copy(halfClockChars2, clockChars)
	for i := rune(0); i < 12; i++ {
		halfClockChars2 = append(halfClockChars2, string([]rune{clockOneThirty + i}))
	}
	// Create clock sets
	CharSets[Clock] = settings{
		150 * time.Millisecond,
		clockChars,
		&defaultPalette,
	}
	CharSets[HalfClock] = settings{
		300 * time.Millisecond,
		halfClockChars,
		&defaultPalette,
	}
	CharSets[HalfClock2] = settings{
		150 * time.Millisecond,
		halfClockChars2,
		&defaultPalette,
	}
}

func checkCharSet(c []string) error {
	if l := len(c); l > maxCharSetSize {
		return fmt.Errorf("spinner: given charset is too big: %v, max: %v", l, maxCharSetSize)
	}
	var widths []int
	for _, c := range c {
		width := runewidth.StringWidth(c)
		widths = append(widths, width)
	}
	for _, w := range widths {
		if w != widths[0] {
			return fmt.Errorf("spinner: ambiguous widths for char set:\n %v\n %v", c, widths)
		}
	}
	return nil
}
