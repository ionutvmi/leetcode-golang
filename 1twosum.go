package main

import "sort"

func twoSum_bruteForce(nums []int, target int) []int {

	for i, n := range nums { // O(n)
		for j, m := range nums { // O(n)
			if i == j {
				continue
			}

			if n+m == target {
				return []int{i, j}
			}
		}
	}
	// total O(n^2)

	return []int{}
}

func twoSum_binarySearch(nums []int, target int) []int {
	numsCount := len(nums)
	sort.Ints(nums) // O(n log n)

	for i, firstNumber := range nums { // O(n)
		secondNumber := target - firstNumber

		j := sort.SearchInts(nums, secondNumber) // O(log n)

		if i == j {
			continue
		}

		if j < numsCount && nums[j] == secondNumber {
			return []int{i, j}
		}
	}

	// total O(n log n)

	return []int{}
}

func twoSum(nums []int, target int) []int {
	numsMap := map[int]int{}

	for i, n := range nums { // O(n)
		numsMap[n] = i
	}

	for i, firstNumber := range nums { // O(n)
		secondNumber := target - firstNumber

		// do we have the second number?
		j, ok := numsMap[secondNumber] // O(1)

		if ok && i != j {
			return []int{i, j}
		}
	}

	// total O(n)

	return []int{}
}
