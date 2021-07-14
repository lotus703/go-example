package example

func FindLength(nums1 []int, nums2 []int) int {
	count := 0
	subIndex := 0
	for i := 0; i < len(nums1); i++ {
		for j := subIndex; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				count++
				subIndex++
			}
		}
	}
	return count
}
