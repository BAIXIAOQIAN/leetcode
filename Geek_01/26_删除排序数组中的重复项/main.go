package main

//快慢指针
func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	left, right := 1, 1
	for right < len(nums) {
		if nums[right] != nums[right-1] {
			nums[left] = nums[right]
			left++
		}
		right++
	}
	return left
}
