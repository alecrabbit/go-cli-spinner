package spinner

import (
    "fmt"

    "github.com/mattn/go-runewidth"
)

const (
    clockOneOClock = '\U0001F550'
    clockOneThirty = '\U0001F55C'
)

// Declare spinner types
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
    Dots21
    Dots22
    Dots23
    Dots24
    Dots25
    Dots26
    Arrows03
)

// CharSets contains the available character sets
var CharSets = map[int][]string{
    // ← 1
    // ↖ 2
    // ↑ 1
    // ↗ 2
    // → 1
    // ↘ 2
    // ↓ 1
    // ↙ 2
    // Arrows: {"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"}, // Ambiguous width

    BlockVertical:  {"▁", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃", "▁"},
    BouncingBlock:  {"▖", "▘", "▝", "▗"},
    RotatingCircle: {"◐", "◓", "◑", "◒"},
    Snake:          {"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"},
    Snake2:         {"⠏", "⠛", "⠹", "⢸", "⣰", "⣤", "⣆", "⡇",},
    Snake3:         {
        "⢀⠀",
        "⡀⠀",
        "⠄⠀",
        "⢂⠀",
        "⡂⠀",
        "⠅⠀",
        "⢃⠀",
        "⡃⠀",
        "⠍⠀",
        "⢋⠀",
        "⡋⠀",
        "⠍⠁",
        "⢋⠁",
        "⡋⠁",
        "⠍⠉",
        "⠋⠉",
        "⠋⠉",
        "⠉⠙",
        "⠉⠙",
        "⠉⠩",
        "⠈⢙",
        "⠈⡙",
        "⢈⠩",
        "⡀⢙",
        "⠄⡙",
        "⢂⠩",
        "⡂⢘",
        "⠅⡘",
        "⢃⠨",
        "⡃⢐",
        "⠍⡐",
        "⢋⠠",
        "⡋⢀",
        "⠍⡁",
        "⢋⠁",
        "⡋⠁",
        "⠍⠉",
        "⠋⠉",
        "⠋⠉",
        "⠉⠙",
        "⠉⠙",
        "⠉⠩",
        "⠈⢙",
        "⠈⡙",
        "⠈⠩",
        "⠀⢙",
        "⠀⡙",
        "⠀⠩",
        "⠀⢘",
        "⠀⡘",
        "⠀⠨",
        "⠀⢐",
        "⠀⡐",
        "⠀⠠",
        "⠀⢀",
        "⠀⡀",
    },
    BouncingBlock2: {
        "|   ",
        " |  ",
        "  | ",
        "   |",
        "   |",
        "  | ",
        " |  ",
        "|   ",
    },
    Dots13:          {"⠁", "⠂", "⠄", "⡀", "⢀", "⠠", "⠐", "⠈"},
    Dots14:          {"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
    BlockHorizontal: {"▉", "▊", "▋", "▌", "▍", "▎", "▏", "▎", "▍", "▌", "▋", "▊", "▉"},
    // ■ 1
    // □ 1
    // ▪ 2
    // ▫ 2
    // Toggle:          {"■", "□", "▪", "▫"}, // Ambiguous width
    Arrows01: {"←", "↑", "→", "↓"},
    Arrows02: {"⇐", "⇖", "⇑", "⇗", "⇒", "⇘", "⇓", "⇙"},
    Dots21:   {"⠁", "⠁", "⠉", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠤", "⠄", "⠄", "⠤", "⠠", "⠠", "⠤", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋", "⠉", "⠈", "⠈"},
    Dots22:   {"⠈", "⠉", "⠋", "⠓", "⠒", "⠐", "⠐", "⠒", "⠖", "⠦", "⠤", "⠠", "⠠", "⠤", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋", "⠉", "⠈"},
    Dots23:   {"⠁", "⠉", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠤", "⠄", "⠄", "⠤", "⠴", "⠲", "⠒", "⠂", "⠂", "⠒", "⠚", "⠙", "⠉", "⠁"},
    Dots24:   {"⠋", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋"},
    Dots25:   {".  ", ".. ", "...", " ..", "  .", "   "},
    Dots26:   {"⢹", "⢺", "⢼", "⣸", "⣇", "⡧", "⡗", "⡏",},
    Arrows03: {"▹▹▹▹▹", "▸▹▹▹▹", "▹▸▹▹▹", "▹▹▸▹▹", "▹▹▹▸▹", "▹▹▹▹▸"},
}

// type S struct {
// 	CharSet 	[]string
// 	Interval 	time.Duration
// }
// // Settings contains the spinner type specific settings
// var SettingsSets = map[int][]string{
// 	Arrows:  {"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"},
// 	BlockVertical:  {"▁", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃", "▁"},
// 	BouncingBlock:  {"▖", "▘", "▝", "▗"},
// 	RotatingCircle:  {"◐", "◓", "◑", "◒"},
// }

func init() {
    for i := rune(0); i < 12; i++ {
        CharSets[Clock] = append(CharSets[Clock], string([]rune{clockOneOClock + i}))
        CharSets[HalfClock] = append(CharSets[HalfClock], string([]rune{clockOneOClock + i}), string([]rune{clockOneThirty + i}))
    }
    // Check CharSets for width conformity
    for n, _ := range CharSets {
        var widths []int
        for _, c := range CharSets[n] {
            width := runewidth.StringWidth(c)
            widths = append(widths, width)
        }
        for _, e := range widths {
            if e != widths[0] {
                // fmt.Println(n, ":", CharSets[n])
                // fmt.Println(n, ":", widths)
                panic(fmt.Sprintf("\nAmbiguous widths for char set [%v]\n %v\n %v\n", n, CharSets[n], widths))
            }
        }
    }

}
