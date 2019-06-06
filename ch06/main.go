package main

import "fmt"

//给定一个字符串，和行数，给出Z字形排列后的结果

func convert(s string, numRows int) string {
	if len(s) <= 1 || numRows <= 1 {
		return s
	}

	var str [][]rune
	str = make([][]rune, numRows)

	length := (len(s)/(numRows+numRows-2) + 1) * (numRows - 1)

	for i := 0; i < numRows; i++ {
		str[i] = make([]rune, length)
	}

	strs := []rune(s)

	i := 0
	for r := 0; r < length; r++ {
		if r == 0 || r%(numRows-1) == 0 {
			for l := 0; l < numRows; l++ {
				if i >= len(s) {
					break
				}
				str[l][r] = strs[i]
				i++
			}
		} else {
			if numRows == 2 {
				for l := numRows - 1; l > 0; l-- {
					if i >= len(s) {
						break
					}
					str[l][r] = strs[i]
					i++
				}
			} else {
				var l int
				if r < numRows {
					l = (numRows - 1) - r%numRows
				} else {
					l = (numRows - 1) - r%(numRows-1)
				}
				if i >= len(s) {
					break
				}
				str[l][r] = strs[i]
				i++
			}
		}
	}
	return twos2one(str)
}

//把二维字符数组转化成字符串
func twos2one(strs [][]rune) string {
	h, l := len(strs), len(strs[0])

	var ru []rune
	for i := 0; i < h; i++ {
		for j := 0; j < l; j++ {
			if strs[i][j] != 0 {
				fmt.Printf(string(strs[i][j]))
				ru = append(ru, strs[i][j])
			} else {
				fmt.Printf("_")
			}
		}
		println()
	}
	return string(ru)
}

//优化内存和运行时间
func goodfunc(s string, numRows int) string {
	if len(s) <= 1 || numRows <= 1 {
		return s
	}

	var result []rune

	str := []rune(s)
	var n int = len(s)
	var cycleLen int = 2*numRows - 2

	for i := 0; i < numRows; i++ {
		for j := 0; j+i < n; j += cycleLen { //每次加一个周期
			result = append(result, str[i+j])
			if i != 0 && i != numRows-1 && j+cycleLen-i < n { //除去第0行和最后一行
				result = append(result, str[j+cycleLen-i])
			}
		}
	}

	return string(result)
}
func main() {
	str := convert("PAYPALISHIRING", 4)
	str1 := goodfunc("PAYPALISHIRING", 4)
	println(str)
	println(str1)
}
