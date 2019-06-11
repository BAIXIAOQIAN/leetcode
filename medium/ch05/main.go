package main

//给定一个字符串，找出最长的回文字符串

//我的方法，运行时间太长
func getlongstring(s string) string {

	if len(s) <= 1 {
		return s
	}

	var result []string
	str := []rune(s)

	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j <= len(s); j++ {
			if huiwenstring(str[i:j]) {
				result = append(result, string(str[i:j]))
			}
		}
	}
	return maxstring(result)
}

//判断一个字符串是否是回文字符串
func huiwenstring(s []rune) bool {
	str := []rune(s)

	l, r := 0, len(s)-1
	for l <= r {
		if str[l] == str[r] {
			l++
			r--
		} else {
			return false
		}
	}
	return true
}

//获取一个字符数组中的最长字符串
func maxstring(s []string) string {
	var max int = 0
	var index int
	for i, str := range s {
		if len(str) >= max {
			max = len(str)
			index = i
		}
	}
	return s[index]
}

//优化运行时间
func timequickGetstring(s string) string {
	if len(s) <= 1 {
		return s
	}

	var n int = len(s)
	dp := make([]bool, n)
	var maxPal string = ""

	str := []rune(s)

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= i; j-- {
			dp[j] = str[i] == str[j] && (j-i < 2 || dp[j-1])

			if dp[j] && j-i+1 > len(maxPal) {
				maxPal = string(str[i : j+1])
			}
		}
	}
	return maxPal
}

func main() {
	str := timequickGetstring("ababa")
	println(str)
}
