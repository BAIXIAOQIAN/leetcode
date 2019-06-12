package main

import (
	"fmt"
	"sort"
)

//给定一个数组，找出所有满足三个数相加等于０的三元组
func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j, k := i+1, len(nums)-1
		for j < k {
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			if k < len(nums)-1 && nums[k] == nums[k+1] {
				k--
				continue
			}
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return res
}

//暴力解法
func threeSums(nums []int) [][]int {
	var res [][]int

	if len(nums) < 3 {
		return res
	}
	return res
}
func main() {
	nums := []int{-1, 0, -2, 1, 2}
	res := threeSum(nums)

	for _, v := range res {
		fmt.Printf("%+v \n", v)
	}
}
