package main

//给定一个正数，返回翻转的数字

func reverse(x int) int {
	s := x
	if s < 10 && s > -10 {
		return s
	}

	if s < 0 {
		s = -1 * s
	}
	var result int
	var sum []int

	for s > 0 {
		n := s % 10
		s = s / 10
		sum = append(sum, n)
	}

	for i := 0; i < len(sum); i++ {
		result += sum[i] * cum(len(sum)-i-1)
	}

	if x < 0 {
		result = -1 * result
	}

	if result > 2147483647 || result < -2147483648 {
		return 0
	}
	return result
}

func cum(n int) int {
	var result int = 1
	for i := 0; i < n; i++ {
		result = result * 10
	}
	return result
}

//最快的算法
func fasereverse(x int) int {
	var res int = 0

	for x != 0 {
		res = res*10 + x%10
		x = x / 10
	}

	if res > 2147483647 || res < -2147483648 {
		return 0
	}
	return res
}

func main() {
	result := reverse(-321)
	println(result)
}
