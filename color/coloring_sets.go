package color

import (
	"fmt"
)

// Names for colorizing sets
const (
	CNoColor = iota
	CDefault
	CDark
	CBlink
	CRedBoldItalic
	C256Rainbow
	C256YellowWhite
	C256RSingle
)

func init() {
	Prototypes[CDefault] = Prototypes[CNoColor]
}

// Prototypes contains colorizing sets
var Prototypes = map[int]StylePrototype{
	CNoColor: {
		TNoColor,
		[][]int{},
		func(a [][]int) []string {
			return []string{"%s"}
		},
	},
	CDark: {
		TColor16,
		[][]int{},
		func(a [][]int) []string {
			return []string{"\x1b[2m%s\x1b[0m"}
		},
	},
	CBlink: {
		TColor16,
		[][]int{},
		func(a [][]int) []string {
			return []string{"\x1b[5m%s\x1b[0m"}
		},
	},
	CRedBoldItalic: {
		TColor16,
		[][]int{},
		func(a [][]int) []string {
			return []string{"\x1b[31;1;3m%s\x1b[0m"}
		},
	},
	C256Rainbow: {
		TColor256,
		[][]int{
			{196},
			{202},
			{208},
			{214},
			{220},
			{226},
			{190},
			{154},
			{118},
			{82},
			{46},
			{47},
			{48},
			{49},
			{50},
			{51},
			{45},
			{39},
			{33},
			{27},
			{56},
			{57},
			{93},
			{129},
			{165},
			{201},
			{200},
			{199},
			{198},
			{197},
		},
		func(a [][]int) []string {
			a = multiply(a, 3)
			r := make([]string, len(a))
			for i, v := range a {
				r[i] = fmt.Sprintf("\x1b[38;5;%vm%s\x1b[0m", v[0], "%s")
			}
			return r
		},
	},
	C256YellowWhite: {
		TColor256,
		[][]int{
			{226},
			{227},
			{228},
			{229},
			{229},
			{230},
			{230},
			{230},
			{231},
			{231},
			{231},
			{231},
			{230},
			{230},
			{230},
			{229},
			{229},
			{228},
			{227},
			{226},
		},
		func(a [][]int) []string {
			r := make([]string, len(a))
			for i, v := range a {
				r[i] = fmt.Sprintf("\x1b[38;5;%vm%s\x1b[0m", v[0], "%s")
			}
			return r
		},
	},
	C256RSingle: {
		TColor256,
		[][]int{
			{196, 232, 3},
			{202, 232, 3},
			{208, 232, 3},
			{214, 232, 3},
			{220, 232, 3},
			{226, 232, 3},
			{190, 232, 3},
			{154, 232, 3},
			{118, 232, 3},
			{82, 232, 3},
			{46, 232, 3},
			{47, 232, 3},
			{48, 232, 3},
			{49, 232, 3},
			{50, 232, 3},
			{51, 232, 3},
			{45, 232, 3},
			{39, 232, 3},
			{33, 232, 3},
			{27, 232, 3},
			{56, 232, 3},
			{57, 232, 3},
			{93, 232, 3},
			{129, 232, 3},
			{165, 232, 3},
			{201, 232, 3},
			{200, 232, 3},
			{199, 232, 3},
			{198, 232, 3},
			{197, 232, 3},
		},
		func(a [][]int) []string {
			r := make([]string, len(a))
			for i, v := range a {
				r[i] = fmt.Sprintf("\x1b[38;5;%v;48;5;%v;%vm%s\x1b[0m", v[0], v[1], v[2], "%s")
			}
			return r
		},
	},
}

func multiply(c [][]int, factor int) [][]int {
	r := make([][]int, len(c)*factor)
	for i := range r {
		r[i] = c[i/factor]
	}
	return r
}
