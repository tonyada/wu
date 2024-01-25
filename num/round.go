package num

import "math"

// Get Round 官方四舍五入
func Round(x float64) int {
	return int(math.Round(x))
}

// Get round up RoundUp(230.55) = 231
func RoundUp(x float64) int {
	return int(math.Ceil(x))
}

// Get round down RoundDown(230.55) = 230
func RoundDown(x float64) int {
	return int(math.Floor(x))
}
