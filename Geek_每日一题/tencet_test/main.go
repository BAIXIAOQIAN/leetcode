package main

import "strconv"

//将一个长度最多为30位的十进制数转换为二进制
func main() {
	num := "4387197563"
	res := num2Dec(num)
	println(res)
}

func num2Dec(numStr string) string {
	var res string

	numArry := []int{}
	for i := 0; i < len(numStr); i++ {
		n, _ := strconv.Atoi(string(numStr[i]))
		numArry[i] = n
	}
	for len(numArry) != 0 {
		//%2
		var tmpArray []int
		for i := 0; i < len(numArry); i++ {
			t := numArry[i] / 2
			y := numArry[i] % 2
			if i < len(numArry)-1 {
				numArry[i] = numArry[i] + y
			}
			tmpArray = append(tmpArray, t)
			if i == len(numArry)-1 {
				st := strconv.Itoa(y)
				res = st + res
			}
		}
		numArry = tmpArray
	}

	return res
}

/*
10
5 0
2 10
1 010
0 1010
*/

/*
5 6 3
2 8 1 1
*/
