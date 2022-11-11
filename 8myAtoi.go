package main

import (
	"math"
	"strings"
)

func myAtoi(s string) int {

	chars, isPositive := GetDigits(strings.Split(s, ""))
	size := len(chars)

	if size == 0 {
		return 0
	}

	if size > 10 {
		if isPositive {
			return math.MaxInt32
		} else {
			return math.MinInt32
		}
	}

	var result int64 = 0

	for index, c := range chars {
		digit := int64(c[0] - '0')
		pos := size - index - 1
		result += int64(math.Pow10(pos) * float64(digit))
	}

	if result > math.MaxInt32 {
		if isPositive {
			return math.MaxInt32
		} else {
			return math.MinInt32
		}
	}

	if !isPositive {
		return -int(result)
	}

	return int(result)
}

func GetDigits(chars []string) ([]string, bool) {
	numberHasStarted := false
	result := []string{}
	isPositive := true
	size := len(chars)

	for index, c := range chars {
		if numberHasStarted && !isDigit(c) {
			break
		}

		if c == " " {
			continue
		}

		if c == "-" && index < size-1 {
			isPositive = false
			numberHasStarted = true
			continue
		}

		if c == "+" && index < size-1 {
			isPositive = true
			numberHasStarted = true
			continue
		}

		if len(result) == 0 && c == "0" {
			numberHasStarted = true
			continue
		}

		if !isDigit(c) {
			break
		}

		numberHasStarted = true
		result = append(result, c)
	}

	return result, isPositive
}

func isDigit(c string) bool {
	return strings.Contains("0123456789", c)
}
