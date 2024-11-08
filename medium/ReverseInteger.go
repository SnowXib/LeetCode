package medium

import (
	"math"
)

func Reverse(x int) int {
	result_array := make([]int, 0, 16)
	var result_digit int
	var plus_var int = 10
	var minus bool

	if x < 0 {
		minus = true
		x *= -1
	}

	// if x < -1<<31 || x > 1<<31-1 {
	// return 0
	// }

	for i := 10; i <= x*10; i = plus_var {
		remains := x % i
		remains /= (i / 10)
		result_array = append(result_array, remains)
		plus_var = i * 10
	}
	for i, v := range result_array {
		upper := int(math.Pow(10, float64(len(result_array)-i-1))) * v
		result_digit += upper
	}

	// if result_digit < -1<<31 || result_digit > 1<<31-1 {
	// return 0
	// }

	if minus {
		result_digit *= -1
	}

	return result_digit
}
