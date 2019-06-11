package main

//给定n个非负整数a1,a2,...,an,其中每个表示坐标(i,ai)处的点。绘制n条垂直线，使得线i的两个端点位于(i,ai)和(i,0)

//Time: O(n^2), Space: O(1)
func maxArea(height []int) int {
	if len(height) <= 1 {
		return 0
	}

	var res int
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			vector := minNum(height[i], height[j]) * (j - i)
			if res <= vector {
				res = vector
			}
		}
	}
	return res
}

// Two Pointers
// Time: O(n), Space: O(1)

func maxArea2(height []int) int {
	lo, hi := 0, len(height)-1
	ret := 0
	for lo < hi {
		area := (hi - lo) * minNum(height[lo], height[hi])
		if area > ret {
			ret = area
		}
		if height[lo] < height[hi] {
			lo++
		} else {
			hi--
		}
	}
	return ret
}

func minNum(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	res := maxArea2(height)

	println(res)
}
