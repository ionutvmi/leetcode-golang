package main

import "math"

func reverse(x int) int {
	var result int32 = 0

	isNegative := false

	if x < 0 {
		isNegative = true
		x = -x
	}

	numLen := getDigitsLen(x)

	for x > 0 {
		digit := int32(x % 10)
		x /= 10

		numLen--
		base := math.Pow10(numLen)

		if base > math.MaxInt32 && digit > 0 {
			return 0
		}

		result += int32(base) * digit

		if result < 0 || (result < int32(base) && digit > 0) {
			return 0
		}
	}

	if isNegative {
		return -int(result)
	}

	return int(result)
}

func getDigitsLen(n int) int {

	if n == 0 {
		return 1
	}

	count := 0

	for n != 0 {
		n /= 10
		count++
	}

	return count
}
