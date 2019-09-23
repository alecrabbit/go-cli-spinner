// Contains auxiliary constants, functions and variables
package aux

func Bounds(f float32) float32 {
    if f < 0 {
        f = 0
    }
    if f > 1 {
        f = 1
    }
    return f
}
