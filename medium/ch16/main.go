package main

import (
	"fmt"
	"sort"
)

//最接近的三个数之和 时间复杂度o(n^3)
func threeSumClosest(nums []int, target int) int {
	var res = target

	//对nums进行排序
	sort.Ints(nums)

	resmap := make(map[int]int)
	//遍历排序后的数组
	for i := 0; i < len(nums)-2; i++ {
		println(i)
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				sum := nums[i] + nums[j] + nums[k]
				if sum > target {
					res = sum - target
				} else {
					res = target - sum
				}
				resmap[res] = sum
			}
		}
	}

	var min int = res

	for k, v := range resmap {
		fmt.Printf("k: %v, v: %v\n", k, v)
		if min > k {
			min = k
		}
	}
	return resmap[min]
}

func main() {
	nums := []int{0, 2, 1, -3}
	res := threeSumClosest(nums, 1)
	println(res)
}
