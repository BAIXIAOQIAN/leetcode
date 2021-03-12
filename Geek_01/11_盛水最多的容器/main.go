package main

func maxArea(height []int) int {
	area, max := 0, 0

	left, right := 0, len(height)-1
	for left < right {
		if hl, hr, diff := height[left], height[right], right-left; hl > hr {
			area = hr * diff
			right--
		} else {
			area = hl * diff
			left++
		}
		if area > max {
			max = area
		}
	}
	return 0
}

//思路

//1. 暴力
//2. 双指针
