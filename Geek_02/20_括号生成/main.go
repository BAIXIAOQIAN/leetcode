package main

func generateParenthesis(n int) []string {
	var res []string
	//生成括号
	var gen func(l, r int, str string)

	gen = func(l, r int, str string) {
		if len(str) == 2*n {
			res = append(res, str)
			return
		}

		if l < n {
			gen(l+1, r, str+"(")
		}

		if r < l {
			gen(l, r+1, str+")")
		}
	}

	gen(0, 0, "")
	return res
}

func main() {
	generateParenthesis(3)
}
