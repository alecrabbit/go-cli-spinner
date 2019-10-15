package spinner

import (
    "fmt"
    "time"

    "github.com/mattn/go-runewidth"
)

const (
    clockOneOClock = '\U0001F550'
    clockOneThirty = '\U0001F55C'
)

const MaxCharSetSize = 60

// Declared spinner types
const (
    BlockVertical int = iota
    // Arrows
    BouncingBlock
    BouncingBlock2
    RotatingCircle
    Clock
    HalfClock
    HalfClock2
    Snake
    Snake2
    FlyingDots
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

type Settings struct {
    interval time.Duration // interval between spinner refreshes
    chars    []string      //
}

// CharSets contains the available character sets
var NewCharSets = map[int]Settings{
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
        120,
        []string{"←", "↑", "→", "↓"},
    },
    Arrows02: {
        120,
        []string{"↖", "↗", "↘", "↙"},
    },
    Arrows03: {
        120,
        []string{"⇐", "⇖", "⇑", "⇗", "⇒", "⇘", "⇓", "⇙"},
    },
    Arrows04: {
        120,
        []string{"▹▹▹▹▹", "▸▹▹▹▹", "▹▸▹▹▹", "▹▹▸▹▹", "▹▹▹▸▹", "▹▹▹▹▸"},
    },
    Simple: {
        120,
        []string{"|", "\\", "─", "/"},
    },
    Dev: { // Singe character used for dev purposes
        400,
        []string{"+"},
    },
    Dev2: { // Number characters used for dev purposes
        250,
        []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"},
    },
    BlockVertical: { // Number characters used for dev purposes
        120,
        []string{"▁", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃", "▁"},
    },
    BlockHorizontal: { // Number characters used for dev purposes
        120,
        []string{"▉", "▊", "▋", "▌", "▍", "▎", "▏", "▎", "▍", "▌", "▋", "▊", "▉"},
    },
    BouncingBlock: { // Number characters used for dev purposes
        120,
        []string{"▖", "▘", "▝", "▗"},
    },
    RotatingCircle: { // Number characters used for dev purposes
        120,
        []string{"◐", "◓", "◑", "◒"},
    },
    Snake: { // Number characters used for dev purposes
        150,
        []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"},
    },
    Snake2: {
        120,
        []string{"⠏", "⠛", "⠹", "⢸", "⣰", "⣤", "⣆", "⡇"},
    },
    FlyingDots: {
        120,
        []string{
                "⢀⠀", "⡀⠀", "⠄⠀", "⢂⠀", "⡂⠀", "⠅⠀", "⢃⠀", "⡃⠀", "⠍⠀", "⢋⠀", "⡋⠀", "⠍⠁", "⢋⠁", "⡋⠁", "⠍⠉", "⠋⠉", "⠋⠉",
                "⠉⠙", "⠉⠙", "⠉⠩", "⠈⢙", "⠈⡙", "⢈⠩", "⡀⢙", "⠄⡙", "⢂⠩", "⡂⢘", "⠅⡘", "⢃⠨", "⡃⢐", "⠍⡐", "⢋⠠", "⡋⢀", "⠍⡁",
                "⢋⠁", "⡋⠁", "⠍⠉", "⠋⠉", "⠋⠉", "⠉⠙", "⠉⠙", "⠉⠩", "⠈⢙", "⠈⡙", "⠈⠩", "⠀⢙", "⠀⡙", "⠀⠩", "⠀⢘", "⠀⡘", "⠀⠨",
                "⠀⢐", "⠀⡐", "⠀⠠", "⠀⢀", "⠀⡀",
            },
    },

    // : ,
    // BouncingBlock2: {
    //     "|   ", " |  ", "  | ", "   |", "   |", "  | ", " |  ", "|   "},
    // Dots13: {"⠁", "⠂", "⠄", "⡀", "⢀", "⠠", "⠐", "⠈"},
    // Dots14: {"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
    //
    // Dots21: {"⠁", "⠁", "⠉", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠤", "⠄", "⠄", "⠤", "⠠", "⠠", "⠤", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋", "⠉", "⠈", "⠈"},
    // Dots22: {"⠈", "⠉", "⠋", "⠓", "⠒", "⠐", "⠐", "⠒", "⠖", "⠦", "⠤", "⠠", "⠠", "⠤", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋", "⠉", "⠈"},
    // Dots23: {"⠁", "⠉", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠤", "⠄", "⠄", "⠤", "⠴", "⠲", "⠒", "⠂", "⠂", "⠒", "⠚", "⠙", "⠉", "⠁"},
    // Dots24: {"⠋", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋"},
    // Dots25: {".  ", ".. ", "...", " ..", "  .", "   "},
    // Dots26: {"⢹", "⢺", "⢼", "⣸", "⣇", "⡧", "⡗", "⡏"},
    //
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
    NewCharSets[Clock] = Settings{150, clockChars}
    NewCharSets[HalfClock] = Settings{300, halfClockChars}
    NewCharSets[HalfClock2] = Settings{150, halfClockChars2}
    // Check CharSets for width conformity
    checkCharSets()
}

func checkCharSets() {
    // Check NewCharSets for width conformity
    for n := range NewCharSets {
        err := checkCharSet(NewCharSets[n].chars)
        if err != nil {
            panic(err)
        }
    }
}

func checkCharSet(c []string) error {
    if len(c) > MaxCharSetSize {
        return fmt.Errorf("given charset is too big: %v, max: %v", len(c), MaxCharSetSize)
    }
    var widths []int
    for _, c := range c {
        width := runewidth.StringWidth(c)
        widths = append(widths, width)
    }
    for _, w := range widths {
        if w != widths[0] {
            return fmt.Errorf("\nAmbiguous widths for char set:\n %v\n %v\n", c, widths)
        }
    }
    return nil
}
