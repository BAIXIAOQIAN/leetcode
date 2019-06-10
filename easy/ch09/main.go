package main

//判断一个整数是不是回文
func isPalindrome(x int) bool {
	var res int

	n := x
	for x > 0 {
		res = res*10 + x%10
		x = x / 10
	}
	if res == n {
		return true
	}
	return false
}

func main() {
	res := isPalindrome(-121)
	println(res)
}
