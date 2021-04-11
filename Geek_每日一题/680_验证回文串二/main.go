package main

func main() {
	s := "abcafasd"
	println(validPalindrome2(s))
}

//递归
func validPalindrome(s string) bool {
	l, r := 0, len(s)-1

	var validStr func(l, r int, isdel bool) bool

	validStr = func(l, r int, isdel bool) bool {
		if l == r || l > r {
			return true
		}

		if s[l] == s[r] {
			return validStr(l+1, r-1, isdel)
		}

		if (s[l+1] == s[r] || s[l] == s[r-1]) && !isdel {
			return validStr(l+2, r-1, true) || validStr(l+1, r-2, true)
		}
		return false
	}

	return validStr(l, r, false)
}

func validPalindrome2(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		if s[l] != s[r] {
			return isPire(s[l:r]) || isPire(s[l+1:r+1])
		}
		l++
		r--
	}
	return true
}

func isPire(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
