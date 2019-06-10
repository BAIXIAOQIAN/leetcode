package main

import (
	"math"
	"strconv"
)

//字符串转数字
/*1.剔除空的字符
  2.找到第一个字符
  3.判断第一个字符是不是有效"+"-"或数字不等于0
  4.判断第一个字符后的字符是否有效
*/
func myAtoi(str string) int {
	var res int = 0

	str1 := delete(str)
	arry := []rune(str1)

	if len(arry) <= 1 {
		if len(arry) == 1 && string(arry[0]) >= "0" && string(arry[0]) <= "9" {
			result, _ := strconv.Atoi(str)
			return result
		} else {
			return 0
		}
	}

	if string(arry[0]) != "+" && string(arry[0]) != "-" {
		if string(arry[0]) < "0" || string(arry[0]) > "9" {
			return 0
		} else {
			for _, v := range arry {
				num, _ := strconv.Atoi(string(v))
				if string(v) < "0" || string(v) > "9" {
					break
				}
				res = res*10 + num
			}
		}
	} else {
		if string(arry[1]) <= "0" || string(arry[1]) > "9" {
			return 0
		} else {
			for _, v := range arry[1:] {
				num, _ := strconv.Atoi(string(v))
				if string(v) < "0" || string(v) > "9" {
					break
				}
				res = res*10 + num
			}
		}
		if string(arry[0]) == "-" {
			res = -1 * res
		}
	}
	if res > 2147483647 {
		res = 2147483647
	} else if res < -2147483648 {
		res = -2147483648
	}
	return res
}

//去除字符串中的空格
func delete(s string) string {
	str := []rune(s)
	var res []rune
	for _, v := range str {
		if string(v) != " " {
			res = append(res, v)
		}
	}
	return string(res)
}

//参考别人的做法
func getSign(s string) int {
	var (
		ans   int64
		start bool
		sign  = 1
	)

	for _, v := range s {
		if '0' <= v && v <= '9' {
			start = true
			ans = ans*10 + int64(v-'0')
			if ans >= math.MaxInt32+1 {
				break
			}
		} else if !start && v == ' ' {
			continue
		} else if !start && v == '+' {
			sign = 1
			start = true
		} else if !start && v == '-' {
			sign = -1
			start = true
		} else {
			break
		}
	}

	ans *= int64(sign)

	if ans > math.MaxInt32 {
		ans = math.MaxInt32
	}

	if ans < math.MinInt32 {
		ans = math.MinInt32
	}

	return int(ans)
}
func main() {
	num := getSign("+ 88")
	println(num)
}
