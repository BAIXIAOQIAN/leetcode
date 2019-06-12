package main

//整数转罗马数字

//暴力解法，列出所有情况
func intToRoman(num int) string {
	var res []rune

	for num > 0 {
		if num >= 1000 {
			n := num / 1000
			num = num % 1000
			for n > 0 {
				res = append(res, 'M')
				n--
			}

		} else if num >= 100 {
			if num >= 900 {
				num = num - 900
				res = append(res, 'C')
				res = append(res, 'M')
			} else if num >= 500 {
				num = num - 500
				res = append(res, 'D')
			} else if num >= 400 {
				num = num - 400
				res = append(res, 'C')
				res = append(res, 'D')
			} else {
				n := num / 100
				num = num % 100
				for n > 0 {
					res = append(res, 'C')
					n--
				}
			}
		} else if num >= 10 {
			if num >= 90 {
				num = num - 90
				res = append(res, 'X')
				res = append(res, 'C')
			} else if num >= 50 {
				num = num - 50
				res = append(res, 'L')
			} else if num >= 40 {
				num = num - 40
				res = append(res, 'X')
				res = append(res, 'L')
			} else {
				n := num / 10
				num = num % 10
				for n > 0 {
					res = append(res, 'X')
					n--
				}
			}
		} else {
			if num >= 9 {
				num = num - 9
				res = append(res, 'I')
				res = append(res, 'X')
			} else if num >= 5 {
				num = num - 5
				res = append(res, 'V')
			} else if num == 4 {
				num = num - 4
				res = append(res, 'I')
				res = append(res, 'V')
			} else {
				for num > 0 {
					res = append(res, 'I')
					num--
				}
			}
		}
	}
	return string(res)
}

//优化暴力解法
func intToRoman2(num int) string {
	var res string
	numArray := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	numMap := make(map[int]string)
	numMap = map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}
	for num > 0 {
		for i := 0; i < len(numArray); i++ {
			if num-numArray[i] >= 0 {
				res += numMap[numArray[i]]
				num = num - numArray[i]
				break
			}
		}
	}
	return res
}
func main() {
	res := intToRoman(3997)
	println(res)
}
