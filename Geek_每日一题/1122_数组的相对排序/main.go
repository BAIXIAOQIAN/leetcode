package main

import "sort"

/*
 * @lc app=leetcode.cn id=1122 lang=golang
 *
 * [1122] 数组的相对排序
 */

// @lc code=start

func main() {
	arr1 := []int{}
	arr2 := []int{}
	res := relativeSortArray(arr1, arr2)
	println(res)
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	arr2m := make(map[int]int, 0)
	for _, v := range arr2 {
		arr2m[v] = 1
	}

	nums := make([]int, 0)
	for i := 0; i < len(arr1); i++ {
		if _, ok := arr2m[arr1[i]]; !ok {
			nums = append(nums, arr1[i])
		}
	}
	sort.Ints(nums)

	//nums -> arr2
	res := make([]int, 0)
	for _, v := range arr2 {
		for _, v2 := range arr1 {
			if v2 == v {
				res = append(res, v2)
			}
		}
	}

	res = append(res, nums...)
	return res
}

// @lc code=end
