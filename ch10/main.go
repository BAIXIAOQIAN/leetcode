package main

//实现一个简单的正则匹配函数
func isMatch(s string, p string) bool {
	if s == p {
		return true
	}
	return false
}

func main() {
	res := isMatch("aa", "a")
	println(res)
}
