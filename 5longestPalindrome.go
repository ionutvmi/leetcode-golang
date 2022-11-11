package main

import "strings"

func longestPalindrome(s string) string {

	chars := strings.Split(s, "")
	result := ""

	for index := range chars {
		oddResult := expandPalindrome(chars, index, index)
		evenResult := expandPalindrome(chars, index, index+1)

		if len(result) < len(oddResult) {
			result = oddResult
		}

		if len(result) < len(evenResult) {
			result = evenResult
		}
	}

	return result
}

func expandPalindrome(chars []string, left int, right int) string {
	size := len(chars)

	for left >= 0 && right < size && chars[left] == chars[right] {
		left--
		right++
	}

	return strings.Join(chars[left+1:right], "")
}
