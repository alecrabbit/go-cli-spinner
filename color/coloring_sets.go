package color

import (
<<<<<<< develop
    "fmt"
=======
	"fmt"
>>>>>>> change colorng model
)

// Names for colorizing sets
const (
<<<<<<< develop
    CNoColor = iota
    CDefault
    CDark
    CBlink
    CRedBoldItalic
    C256Rainbow
=======
    C256Rainbow = iota
>>>>>>> change colorng model
    C256YellowWhite
    C256RSingle
)

<<<<<<< develop
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
=======
// Sets contains colorizing sets
var Sets = map[int][]int{ // TODO: rename this
    C256Rainbow: {
        196, 196, 202, 202, 208, 208,
        214, 214, 220, 220, 226, 226,
        190, 190, 154, 154, 118, 118,
        82, 82, 46, 46, 47, 47,
        48, 48, 49, 49, 50, 50,
        51, 51, 45, 45, 39, 39,
        33, 33, 27, 27, 56, 56,
        57, 57, 93, 93, 129, 129,
        165, 165, 201, 201, 200, 200,
        199, 199, 198, 198, 197, 197},
    C256YellowWhite: {
        226, 227, 228, 229, 229, 230,
        230, 230, 231, 231, 231, 231,
        230, 230, 230, 229, 229, 228,
        227, 226},
}

// Prototypes contains colorizing sets
var Prototypes = map[int]StylePrototype{ // TODO: rename this
    C256RSingle: StylePrototype{
>>>>>>> change colorng model
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
<<<<<<< develop
}

func multiply(c [][]int, factor int) [][]int {
    r := make([][]int, len(c)*factor)
    for i, _ := range r {
        r[i] = c[i/factor]
    }
    return r
=======
>>>>>>> change colorng model
}
