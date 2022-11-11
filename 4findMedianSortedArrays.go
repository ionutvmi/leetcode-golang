package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	n1Size := len(nums1)
	n2Size := len(nums2)
	allSize := n1Size + n2Size

	allNums := []int{}

	n1Index := 0
	n2Index := 0

	for n1Index < n1Size && n2Index < n2Size {

		if nums1[n1Index] < nums2[n2Index] {
			allNums = append(allNums, nums1[n1Index])
			n1Index++
		} else {
			allNums = append(allNums, nums2[n2Index])
			n2Index++
		}

	}

	for n1Index < n1Size {
		allNums = append(allNums, nums1[n1Index])
		n1Index++
	}

	for n2Index < n2Size {
		allNums = append(allNums, nums2[n2Index])
		n2Index++
	}

	if allSize%2 == 0 {
		middle := allSize / 2
		return float64(allNums[middle-1]+allNums[middle]) / 2
	}

	return float64(allNums[allSize/2])
}
