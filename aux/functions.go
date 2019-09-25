package aux

// Bounds restricts f value into bounds of 0..1
func Bounds(f float32) float32 {
	if f < 0 {
		f = 0
	}
	if f > 1 {
		f = 1
	}
	return f
}
