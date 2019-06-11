package main

import (
	"fmt"
	"sort"
)

//怎么优化内存的使用的，目前的时间复杂度是o(n*n/2*n) = o(n^3)
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	if len(s) == 1 {
		return 1
	}

	var long []int
	length, str := len(s), []rune(s)

	var index int = 0
	for i := index; i < length-1; i++ {
		var sum int = 1
		for j := i + 1; j < length; j++ {
			if !exitStr(str[i:j], str[j]) {
				sum++
			} else {
				long = append(long, sum)
				index, sum = j, 1
			}
		}
	}

	sort.Ints(long)
	return long[len(long)-1]
}

func exitStr(s []rune, c rune) bool {
	for _, v := range s {
		if c == v {
			return true
		}
	}
	return false
}

func main() {
	leg := lengthOfLongestSubstring2("abcabcbb")

	fmt.Println(leg)
}

//优化内存 滑动窗口
func lengthOfLongestSubstring2(s string) int {
	freq, l, r := [256]int{0}, 0, -1
	var res int = 0

	str := []rune(s)
	for l < len(s) {
		if r+1 < len(s) && freq[str[r+1]] == 0 {
			fmt.Println(str[r+1])
			r++
			freq[str[r]]++
		} else {
			freq[str[l]]--
			l++
		}
		res = max(res, r-l+1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
