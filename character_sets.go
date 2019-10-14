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

// // CharSets contains the available character sets
// var CharSets = map[int][]string{
//     // Arrows: {"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"}, // Ambiguous width
//     // ← 1
//     // ↖ 2
//     // ↑ 1
//     // ↗ 2
//     // → 1
//     // ↘ 2
//     // ↓ 1
//     // ↙ 2
//
//     Arrows01:       {"←", "↑", "→", "↓"},
//     Arrows02:       {"↖", "↗", "↘", "↙"},
//     Arrows03:       {"⇐", "⇖", "⇑", "⇗", "⇒", "⇘", "⇓", "⇙"},
//     Arrows04:       {"▹▹▹▹▹", "▸▹▹▹▹", "▹▸▹▹▹", "▹▹▸▹▹", "▹▹▹▸▹", "▹▹▹▹▸"},
//     Dev:            {"+"},                                              // Singe character used for dev purposes
//     Dev2:           {"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}, // Number characters used for dev purposes
//     BlockVertical:  {"▁", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃", "▁"},
//     BouncingBlock:  {"▖", "▘", "▝", "▗"},
//     RotatingCircle: {"◐", "◓", "◑", "◒"},
//     Snake:          {"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"},
//     Snake2:         {"⠏", "⠛", "⠹", "⢸", "⣰", "⣤", "⣆", "⡇"},
//     Snake3: {
//         "⢀⠀", "⡀⠀", "⠄⠀", "⢂⠀", "⡂⠀", "⠅⠀", "⢃⠀", "⡃⠀", "⠍⠀", "⢋⠀", "⡋⠀", "⠍⠁", "⢋⠁", "⡋⠁", "⠍⠉", "⠋⠉", "⠋⠉",
//         "⠉⠙", "⠉⠙", "⠉⠩", "⠈⢙", "⠈⡙", "⢈⠩", "⡀⢙", "⠄⡙", "⢂⠩", "⡂⢘", "⠅⡘", "⢃⠨", "⡃⢐", "⠍⡐", "⢋⠠", "⡋⢀", "⠍⡁",
//         "⢋⠁", "⡋⠁", "⠍⠉", "⠋⠉", "⠋⠉", "⠉⠙", "⠉⠙", "⠉⠩", "⠈⢙", "⠈⡙", "⠈⠩", "⠀⢙", "⠀⡙", "⠀⠩", "⠀⢘", "⠀⡘", "⠀⠨",
//         "⠀⢐", "⠀⡐", "⠀⠠", "⠀⢀", "⠀⡀",
//     },
//     BlockHorizontal: {"▉", "▊", "▋", "▌", "▍", "▎", "▏", "▎", "▍", "▌", "▋", "▊", "▉"},
//     BouncingBlock2: {
//         "|   ", " |  ", "  | ", "   |", "   |", "  | ", " |  ", "|   "},
//     Dots13: {"⠁", "⠂", "⠄", "⡀", "⢀", "⠠", "⠐", "⠈"},
//     Dots14: {"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
//
//     Dots21: {"⠁", "⠁", "⠉", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠤", "⠄", "⠄", "⠤", "⠠", "⠠", "⠤", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋", "⠉", "⠈", "⠈"},
//     Dots22: {"⠈", "⠉", "⠋", "⠓", "⠒", "⠐", "⠐", "⠒", "⠖", "⠦", "⠤", "⠠", "⠠", "⠤", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋", "⠉", "⠈"},
//     Dots23: {"⠁", "⠉", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠤", "⠄", "⠄", "⠤", "⠴", "⠲", "⠒", "⠂", "⠂", "⠒", "⠚", "⠙", "⠉", "⠁"},
//     Dots24: {"⠋", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋"},
//     Dots25: {".  ", ".. ", "...", " ..", "  .", "   "},
//     Dots26: {"⢹", "⢺", "⢼", "⣸", "⣇", "⡧", "⡗", "⡏"},
//
//     // Toggle:          {"■", "□", "▪", "▫"}, // Ambiguous width
//     // ■ 1
//     // □ 1
//     // ▪ 2
//     // ▫ 2
//     // Weather: { // Ambiguous width
//     //     "🌤 ", "🌤 ", "🌤 ", "🌤 ", "⛅️", "🌥 ", "☁️ ", "🌧 ", "🌨 ", "🌧 ", "🌨 ", "🌧 ", "🌨 ", "⛈ ",
//     //     "⛈ ", "🌨 ", "⛈ ", "🌧 ", "🌨 ", "☁️ ", "🌥 ", "⛅️", "🌤 ",
//     // },
//     Simple: {"|", "\\", "─", "/"},
// }

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
    Snake2: {
        120,
        []string{"⠏", "⠛", "⠹", "⢸", "⣰", "⣤", "⣆", "⡇"},
    },

    // Dev:            {"+"},                                              // Singe character used for dev purposes
    // Dev2:           {"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}, // Number characters used for dev purposes
    // BlockVertical:  {"▁", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃", "▁"},
    // BouncingBlock:  {"▖", "▘", "▝", "▗"},
    // RotatingCircle: {"◐", "◓", "◑", "◒"},
    // Snake:          {"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"},
    // Snake3: {
    //     "⢀⠀", "⡀⠀", "⠄⠀", "⢂⠀", "⡂⠀", "⠅⠀", "⢃⠀", "⡃⠀", "⠍⠀", "⢋⠀", "⡋⠀", "⠍⠁", "⢋⠁", "⡋⠁", "⠍⠉", "⠋⠉", "⠋⠉",
    //     "⠉⠙", "⠉⠙", "⠉⠩", "⠈⢙", "⠈⡙", "⢈⠩", "⡀⢙", "⠄⡙", "⢂⠩", "⡂⢘", "⠅⡘", "⢃⠨", "⡃⢐", "⠍⡐", "⢋⠠", "⡋⢀", "⠍⡁",
    //     "⢋⠁", "⡋⠁", "⠍⠉", "⠋⠉", "⠋⠉", "⠉⠙", "⠉⠙", "⠉⠩", "⠈⢙", "⠈⡙", "⠈⠩", "⠀⢙", "⠀⡙", "⠀⠩", "⠀⢘", "⠀⡘", "⠀⠨",
    //     "⠀⢐", "⠀⡐", "⠀⠠", "⠀⢀", "⠀⡀",
    // },
    // BlockHorizontal: {"▉", "▊", "▋", "▌", "▍", "▎", "▏", "▎", "▍", "▌", "▋", "▊", "▉"},
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
    // Simple: {"|", "\\", "─", "/"},
}

func init() {
    var clockChars []string
    var halfClockChars []string
    // Fill clocks char sets
    for i := rune(0); i < 12; i++ {
        clockChars = append(clockChars, string([]rune{clockOneOClock + i}))
        halfClockChars = append(halfClockChars, string([]rune{clockOneOClock + i}), string([]rune{clockOneThirty + i}))
    }
    // Create clock sets
    NewCharSets[Clock] = Settings{120, clockChars}
    NewCharSets[HalfClock] = Settings{120, halfClockChars}

    checkCharSets()
}

func checkCharSets() {
    // // Check CharSets for width conformity
    // for n := range CharSets {
    //     var widths []int
    //     for _, c := range CharSets[n] {
    //         width := runewidth.StringWidth(c)
    //         widths = append(widths, width)
    //     }
    //     for _, w := range widths {
    //         if w != widths[0] {
    //             panic(fmt.Sprintf("\nAmbiguous widths for char set [%v]\n %v\n %v\n", n, CharSets[n], widths))
    //         }
    //     }
    // }
    // Check NewCharSets for width conformity
    for n := range NewCharSets {
        checkCharSet(n)
    }
}

func checkCharSet(n int) {
    var widths []int
    for _, c := range NewCharSets[n].chars {
        width := runewidth.StringWidth(c)
        widths = append(widths, width)
    }
    for _, w := range widths {
        if w != widths[0] {
            panic(fmt.Sprintf(
                "\nAmbiguous widths for char set [%v]\n %v\n %v\n", n, NewCharSets[n].chars, widths))
        }
    }
}
