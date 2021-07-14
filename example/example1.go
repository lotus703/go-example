package example

import (
	"math"
)

func GrayCode(n int) []int {
	if n < 1 || n > 16 {
		return nil
	}
	lenNum := math.Pow(float64(n), 2)
	res := []int{}

	num := make([]int, n)
	generateGrayCode(int(lenNum), 0, &num, &res)
	return res
}

func generateGrayCode(n, step int, num *[]int, res *[]int) {
	if n == 0 {
		return
	}
	*res = append(*res, convertBinaryToDeximal(*num))

	if step%2 == 0 {
		(*num)[len(*num)-1] = flipDigit((*num)[len(*num)-1])
	} else {
		index := len(*num) - 1
		for ; index >= 0; index-- {
			if (*num)[index] == 1 {
				break
			}
		}
		if index == 0 {
			(*num)[len(*num)-1] = flipDigit((*num)[len(*num)-1])
		} else {
			(*num)[index-1] = flipDigit((*num)[index-1])
		}
	}
	generateGrayCode(n-1, step+1, num, res)
}

func convertBinaryToDeximal(num []int) int {
	res, rad := 0, 1
	for i := len(num) - 1; i >= 0; i-- {
		res += num[i] * rad
		rad *= 2
	}
	return res
}

func flipDigit(num int) int {
	if num == 0 {
		return 1
	}
	return 0
}
