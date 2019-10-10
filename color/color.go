package color

// SupportLevel represents color support level
type SupportLevel int

const (
	// TNoColor represents terminal no color support
	TNoColor SupportLevel = iota
	// TColor16 represents terminal 16 color level support
	TColor16 SupportLevel = 1 << (4 * iota)
	// TColor256 represents terminal 256 color level support
	TColor256
	// TTrueColor represents terminal true color level support
	TTrueColor
)

// PrototypeHandler represents a function to process ANSIStyles from StylePrototype
type PrototypeHandler func([][]int) []string

// StylePrototype represents a struct to contain ansi styling
type StylePrototype struct {
	Level      SupportLevel
	ANSIStyles [][]int
	Handler    PrototypeHandler
}
