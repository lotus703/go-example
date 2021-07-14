package example

func FindPeakElement(nums []int) int {
	var index int
	peekElement := nums[0]
	for i, v := range nums {
		if peekElement < v {
			peekElement = nums[i]
			index = i
		}
	}
	return index
}
