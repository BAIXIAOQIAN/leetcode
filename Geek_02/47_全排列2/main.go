package main

import (
	"fmt"
	"sort"
)

func main() {
	res := permuteUnique2([]int{1, 1, 2})

	for _, v := range res {
		fmt.Printf("%+v\n", v)
	}
}

func permuteUnique2(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	visited := make([]bool, len(nums))

	var dfs func(path []int)
	dfs = func(path []int) {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if visited[i] || i > 0 && nums[i] == nums[i-1] && !visited[i-1] {
				continue
			}
			visited[i] = true
			path = append(path, nums[i])
			dfs(path)
			path = path[:len(path)-1]
			visited[i] = false
		}
	}

	dfs([]int{})
	return res
}
