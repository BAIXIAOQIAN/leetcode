package main

import "fmt"

//实现一个简单的正则匹配函数
/*
给定两个字符串,判断后一个字符串是否匹配前一个字符串
'.' 匹配任何单个字符。
'*'匹配前面元素的零个或多个。
*/

//递归
func isMatchh(s string, p string) bool {
	return false
}

//动态规划
func isMatch(s string, p string) bool {
	//初始化一个二维数组，用来记录每次比较的结果
	a := make([][]bool, len(p)+1)

	//初始化第一列的值
	for i := 0; i <= len(p); i++ {
		a[i] = make([]bool, len(s)+1)
		a[i][0] = i == 0 || (p[i-1] == '*' && a[i-2][0])
	}

	for i := 1; i <= len(p); i++ {
		for j := 1; j <= len(s); j++ {
			pc, sc := p[i-1], s[j-1]
			fmt.Printf("pc:%v,sc:%v \n", string(pc), string(sc))
			a[i][j] = (matches(pc, sc) && a[i-1][j-1]) || (pc == '*' && ((matches(p[i-2], sc) && a[i][j-1]) || a[i-2][j]))
			println(a[i][j])
		}
	}
	return a[len(p)][len(s)]
}

func matches(pc, sc byte) bool {
	return pc == sc || pc == '.'
}

func main() {
	res := isMatch("aaaa", "a*")
	println(res)
}
