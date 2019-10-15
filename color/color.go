package color

// Level represents color support level
type Level int

const (
    // TNoColor represents terminal no color support
    TNoColor Level = iota
    // TColor16 represents terminal 16 color level support
    TColor16 Level = 1 << (4 * iota)
    // TColor256 represents terminal 256 color level support
    TColor256
    // TTrueColor represents terminal true color level support
    TTrueColor
)

// PrototypeHandler represents a function to process ANSIStyles from StylePrototype
type PrototypeHandler func([][]int) []string

// StylePrototype represents a struct to contain ansi styling
type StylePrototype struct {
    Level      Level
    ANSIStyles [][]int
    Handler    PrototypeHandler
}

var SupportedLevels = map[Level]bool{
    TNoColor: true,
    TColor16: true,
    TColor256: true,
    TTrueColor: false,
}