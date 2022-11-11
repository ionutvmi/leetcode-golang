package main

func maxArea(heightsList []int) int {
	left := 0
	right := len(heightsList) - 1

	maxArea := 0

	for left != right {
		leftH := heightsList[left]
		rightH := heightsList[right]
		containerHeight := min(leftH, rightH)

		area := containerHeight * (right - left)

		if area > maxArea {
			maxArea = area
		}

		if leftH < rightH {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
