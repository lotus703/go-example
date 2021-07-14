package example

import "demo/validate"

func FindLength(nums1 []int, nums2 []int) int {
	count := 0
	subIndex := 0
	if !validate.Validate(nums1) || !validate.Validate(nums2) {
		return -1
	}
	for i := 0; i < len(nums1); i++ {
		if nums1[i] > 100 || nums1[1] < 0 {
			return -1
		}
		for j := subIndex; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				if nums2[i] > 100 || nums2[i] < 0 {
					return -1
				}
				count++
				subIndex++
			}
		}
	}
	return count
}
