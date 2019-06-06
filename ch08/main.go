package main

import "strconv"

//字符串转数字
func myAtoi(str string) int {
	arry := []rune(str)

	var exti bool
	var exti1 bool

	var res int = 0
	for _, v := range arry {
		if string(v) == "-" {
			exti = true
		}

		if string(v) == "+" {
			exti1 = true
		}

		if string(v) >= "0" && string(v) <= "9" {
			num, _ := strconv.Atoi(string(v))
			res = res*10 + num
		}

		if string(v) != "-" && string(v) != "+" && (string(v) < "0" || string(v) > "9") && string(v) != " " {
			break
		}
	}

	if exti && exti1 {
		return 0
	} else if exti {
		res = -1 * res
	}

	if res > 2147483647 {
		res = 2147483647
	} else if res < -2147483648 {
		res = -2147483648
	}
	return res
}

func main() {
	num := myAtoi("+ -987")
	println(num)
}
