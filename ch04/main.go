package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var sums []int

	for i := 0; i < len(nums2); i++ {

	}
	sums = sum(nums1, nums2)

	sort.Ints(sums)

	var result float64
	if len(sums)%2 != 0 {
		index := (len(sums) - 1) / 2
		result = float64(sums[index])
	} else if len(sums) == 1 {
		result = float64(sums[0])
	} else {
		index := (len(sums) - 1) / 2
		result = (float64(sums[index]) + float64(sums[index+1])) / 2
	}
	return result
}

//把两个小数组合成一个大数组
func sum(nums1 []int, nums2 []int) []int {
	var sum []int
	sum = make([]int, len(nums1)+len(nums2))
	for i := 0; i < len(nums1); i++ {
		sum[i] = nums1[i]
	}

	for i := 0; i < len(nums2); i++ {
		sum[len(nums1)+i] = nums2[i]
	}

	fmt.Println(sum)
	return sum
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}

	result := findMedianSortedArrays(nums1, nums2)
	println(result)
}
