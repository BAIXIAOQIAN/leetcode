package main

const (
	max = 2^31 -1
)

func main() {
	res := test(32)
	println(res)
}

func test(dividend int ) int {
	return dividend >> 2
}
//两数相除,要求不能使用乘法、除法和mod运算，如果除法结果溢出，则返回2^31-1
func divide(dividend int, divisor int) int {
	/*思路
	  1.
	 */
	var res int
	for dividend > divisor {
		dividend = dividend - divisor
		res++
	}
	return res
}
