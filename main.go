package main

import (
	"demo/example"
	"fmt"
)

func main() {
	//example1
	result1 := example.GrayCode(0)
	fmt.Println(result1)
	//exmaple2
	nums1 := []int{1, 2, 1, 1, 5, 6, 4}
	result2 := example.FindPeakElement(nums1)
	fmt.Println(result2)
	//example3
	nums2 := []int{1, 2, 3, 2, 1}
	nums3 := []int{3, 2, 1, 4, 7}
	result3 := example.FindLength(nums2, nums3)
	fmt.Println(result3)
}
