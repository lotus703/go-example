package example

import "demo/validate"

func FindPeakElement(nums []int) int {
	var index int
	peekElement := nums[0]
	if !validate.Validate(nums) {
		return -1
	}
	for i, v := range nums {
		if nums[i] == nums[i+1] {
			return -1
		}
		if peekElement < v {
			peekElement = nums[i]
			index = i
		}
	}
	return index
}
