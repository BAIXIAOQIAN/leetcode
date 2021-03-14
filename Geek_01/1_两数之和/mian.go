package main

func twoSum(nums []int, target int) []int {
	var res []int

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return res
}

func twoSum2(nums []int, target int) []int {
	nMap := make(map[int]int)

	for i, v := range nums {
		nMap[v] = i
	}

	for i, v := range nums {
		if j, ok := nMap[target-v]; ok {
			return []int{i, j}
		}
	}
	return nil
}
