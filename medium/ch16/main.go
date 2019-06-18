package main

import "sort"

//最接近的三个数之和
func threeSumClosest(nums []int, target int) int {
	var res = target

	//对nums进行排序
	sort.Ints(nums)

	resmap := make(map[int]int)
	//遍历排序后的数组
	for i := 0; i < len(nums); i++ {
		j, k := i+1, len(nums)-1
		for j < k {
			if j > i+1 {
				j++
				continue
			}
			if k < len(nums)-1 {
				k--
				continue
			}
			sum := nums[i] + nums[j] + nums[k]

			if sum > target {
				res = sum - target
			} else {
				res = target - sum
			}
			resmap[res] = sum
		}
	}

	var min int = 0
	var resvalue int

	for k, v := range resmap {
		if k > min {
			min = k
			resvalue = v
		}
	}
	return resvalue
}

func main() {
	nums := []int{-1,2,1,-4}
	res := threeSumClosest(nums, 9)
	println(res)
}
