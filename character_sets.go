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
    BlockVertical: {
        120,
        []string{"▁", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃", "▁"},
    },
    BlockHorizontal: {
        120,
        []string{"▉", "▊", "▋", "▌", "▍", "▎", "▏", "▎", "▍", "▌", "▋", "▊", "▉"},
    },
    BouncingBlock: {
        120,
        []string{"▖", "▘", "▝", "▗"},
    },
    RotatingCircle: {
        120,
        []string{"◐", "◓", "◑", "◒"},
    },
    Snake: {
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
    FlyingLine: {
        120,
        []string{"|   ", " |  ", "  | ", "   |", "   |", "  | ", " |  ", "|   "},
    },
    Dots10: {
        120,
        []string{"⢄", "⢂", "⢁", "⡁", "⡈", "⡐", "⡠"},
    },
    Dots13: {
        120,
        []string{"⠁", "⠂", "⠄", "⡀", "⢀", "⠠", "⠐", "⠈"},
    },
    Dots14: {
        120,
        []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
    },
    Dots21: {
        120,
        []string{
            "⠁", "⠁", "⠉", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠤", "⠄", "⠄",
            "⠤", "⠠", "⠠", "⠤", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋", "⠉", "⠈", "⠈",
        },
    },
    Dots22: {
        120,
        []string{
            "⠈", "⠉", "⠋", "⠓", "⠒", "⠐", "⠐", "⠒", "⠖", "⠦", "⠤", "⠠",
            "⠠", "⠤", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋", "⠉", "⠈",
        },
    },
    Dots26: {
        120,
        []string{"⢹", "⢺", "⢼", "⣸", "⣇", "⡧", "⡗", "⡏"},
    },
    Blink: {
        200,
        []string{"▓", "▒", "░"},
    },
    Toggle: {
        250,
        []string{"■", "□"},
    },
    Dots23: {
        120,
        []string{
            "⠁", "⠉", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠤", "⠄",
            "⠄", "⠤", "⠴", "⠲", "⠒", "⠂", "⠂", "⠒", "⠚", "⠙", "⠉", "⠁",
        },
    },
    Dots24: {
        120,
        []string{".  ", ".. ", "...", " ..", "  .", "   "},
    },
    Dots25: {
        120,
        []string{"⠋", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋"},
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
    if l := len(c); l > MaxCharSetSize {
        return fmt.Errorf("given charset is too big: %v, max: %v", l, MaxCharSetSize)
    }
    var widths []int
    for _, c := range c {
        width := runewidth.StringWidth(c)
        widths = append(widths, width)
    }
    for _, w := range widths {
        if w != widths[0] {
            return fmt.Errorf("\nambiguous widths for char set:\n %v\n %v\n", c, widths)
        }
    }
    return nil
}
