package spinner

import (
	"fmt"

	"github.com/mattn/go-runewidth"
)

const (
	clockOneOClock = '\U0001F550'
	clockOneThirty = '\U0001F55C'
)

// Declared spinner types
const (
	BlockVertical int = iota
	// Arrows
	BouncingBlock
	BouncingBlock2
	RotatingCircle
	Clock
	HalfClock
	Snake
	Snake2
	Snake3
	Dots13
	Dots14
	BlockHorizontal
	// Toggle
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

// CharSets contains the available character sets
var CharSets = map[int][]string{
	// Arrows: {"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"}, // Ambiguous width
	// ← 1
	// ↖ 2
	// ↑ 1
	// ↗ 2
	// → 1
	// ↘ 2
	// ↓ 1
	// ↙ 2

	Arrows01:       {"←", "↑", "→", "↓"},
	Arrows02:       {"↖", "↗", "↘", "↙"},
	Arrows03:       {"⇐", "⇖", "⇑", "⇗", "⇒", "⇘", "⇓", "⇙"},
	Arrows04:       {"▹▹▹▹▹", "▸▹▹▹▹", "▹▸▹▹▹", "▹▹▸▹▹", "▹▹▹▸▹", "▹▹▹▹▸"},
	Dev:            {"+"},                                              // Singe character used for dev purposes
	Dev2:           {"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}, // Number characters used for dev purposes
	BlockVertical:  {"▁", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃", "▁"},
	BouncingBlock:  {"▖", "▘", "▝", "▗"},
	RotatingCircle: {"◐", "◓", "◑", "◒"},
	Snake:          {"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"},
	Snake2:         {"⠏", "⠛", "⠹", "⢸", "⣰", "⣤", "⣆", "⡇"},
	Snake3: {
		"⢀⠀", "⡀⠀", "⠄⠀", "⢂⠀", "⡂⠀", "⠅⠀", "⢃⠀", "⡃⠀", "⠍⠀", "⢋⠀", "⡋⠀", "⠍⠁", "⢋⠁", "⡋⠁", "⠍⠉", "⠋⠉", "⠋⠉",
		"⠉⠙", "⠉⠙", "⠉⠩", "⠈⢙", "⠈⡙", "⢈⠩", "⡀⢙", "⠄⡙", "⢂⠩", "⡂⢘", "⠅⡘", "⢃⠨", "⡃⢐", "⠍⡐", "⢋⠠", "⡋⢀", "⠍⡁",
		"⢋⠁", "⡋⠁", "⠍⠉", "⠋⠉", "⠋⠉", "⠉⠙", "⠉⠙", "⠉⠩", "⠈⢙", "⠈⡙", "⠈⠩", "⠀⢙", "⠀⡙", "⠀⠩", "⠀⢘", "⠀⡘", "⠀⠨",
		"⠀⢐", "⠀⡐", "⠀⠠", "⠀⢀", "⠀⡀",
	},
	BlockHorizontal: {"▉", "▊", "▋", "▌", "▍", "▎", "▏", "▎", "▍", "▌", "▋", "▊", "▉"},
	BouncingBlock2: {
		"|   ", " |  ", "  | ", "   |", "   |", "  | ", " |  ", "|   "},
	Dots13: {"⠁", "⠂", "⠄", "⡀", "⢀", "⠠", "⠐", "⠈"},
	Dots14: {"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},

	Dots21: {"⠁", "⠁", "⠉", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠤", "⠄", "⠄", "⠤", "⠠", "⠠", "⠤", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋", "⠉", "⠈", "⠈"},
	Dots22: {"⠈", "⠉", "⠋", "⠓", "⠒", "⠐", "⠐", "⠒", "⠖", "⠦", "⠤", "⠠", "⠠", "⠤", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋", "⠉", "⠈"},
	Dots23: {"⠁", "⠉", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠤", "⠄", "⠄", "⠤", "⠴", "⠲", "⠒", "⠂", "⠂", "⠒", "⠚", "⠙", "⠉", "⠁"},
	Dots24: {"⠋", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋"},
	Dots25: {".  ", ".. ", "...", " ..", "  .", "   "},
	Dots26: {"⢹", "⢺", "⢼", "⣸", "⣇", "⡧", "⡗", "⡏"},

	// Toggle:          {"■", "□", "▪", "▫"}, // Ambiguous width
	// ■ 1
	// □ 1
	// ▪ 2
	// ▫ 2
	// Weather: { // Ambiguous width
	//     "🌤 ", "🌤 ", "🌤 ", "🌤 ", "⛅️", "🌥 ", "☁️ ", "🌧 ", "🌨 ", "🌧 ", "🌨 ", "🌧 ", "🌨 ", "⛈ ",
	//     "⛈ ", "🌨 ", "⛈ ", "🌧 ", "🌨 ", "☁️ ", "🌥 ", "⛅️", "🌤 ",
	// },
	Simple: {"|", "\\", "─", "/"},
}

func init() {
	// Fill clocks char sets
	for i := rune(0); i < 12; i++ {
		CharSets[Clock] = append(CharSets[Clock], string([]rune{clockOneOClock + i}))
		CharSets[HalfClock] = append(CharSets[HalfClock], string([]rune{clockOneOClock + i}), string([]rune{clockOneThirty + i}))
	}
	checkCharSets()
}

func checkCharSets() {
	// Check CharSets for width conformity
	for n := range CharSets {
		var widths []int
		for _, c := range CharSets[n] {
			width := runewidth.StringWidth(c)
			widths = append(widths, width)
		}
		for _, w := range widths {
			if w != widths[0] {
				panic(fmt.Sprintf("\nAmbiguous widths for char set [%v]\n %v\n %v\n", n, CharSets[n], widths))
			}
		}
	}
}
