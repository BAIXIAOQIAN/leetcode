package main

//罗马数字转整数
func roman2int(s string) int {
	if len(s) <= 0 {
		return 0
	}

	var res int
	numArray := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	numMap := make(map[string]int)
	numMap = map[string]int{
		"M":  1000,
		"CM": 900,
		"D":  500,
		"CD": 400,
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}

	str := []rune(s)
	for i := 0; i < len(s); i++ {
		for _, v := range numArray {
			if string(str[i]) == v {
				res += numMap[v]
			}
			if i < len(s)-1 {
				if string(str[i])+string(str[i+1]) == v {
					index := string(str[i]) + string(str[i+1])
					res += numMap[index]
					i++
				}
			}
		}
	}
	return res
}

func main() {
	res := roman2int("XI")
	println(res)
}
