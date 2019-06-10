package main

import (
	"fmt"
	"time"
)

/*
给定一个整数数组和一个整数，返回数组中相加之和等于整数的下标
*/

//复杂度o(n^2)
func twoSum(nums []int, target int) []int {
	var res []int

	timeStart := time.Now().UnixNano()
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				res = append(res, i)
				res = append(res, j)
			}
		}
	}
	timeEnd := time.Now().UnixNano()
	println(timeEnd - timeStart)
	return res
}

//优化时间复杂度
func twoSum2(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		_, prs := m[n]
		if prs {
			return []int{m[n], i}
		} else {
			m[target-n] = i
		}
	}
	return nil
}

func main() {
	nums := []int{2, 5, 7, 5, 8, 15, 5, 9, 1, 8}
	res := twoSum2(nums, 10)

	fmt.Printf("res:%+v\n", res)
}
