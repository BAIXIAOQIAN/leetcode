package main

func moveZeroes(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}
}

//思路
//1.把非零元素拷贝到一个新的数组
//2.删除零元素,并记录下个数,然后在末尾补零
//3.双指针,一个指向0,另一个指向非零元素,交换
