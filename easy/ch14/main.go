package main

//返回多个字符串的最长公共子前缀
func longestCommonPrefix(strs []string) string {
	var res string

	if len(strs) <= 1 {
		if len(strs) == 1 {
			return strs[0]
		}
		return ""
	}

	strTmp := []byte(strs[0])

	for j, v := range strTmp {
		var equal bool
		for i := 1; i < len(strs); i++ {
			if len(strs[i]) > j {
				if strs[i][j] == v {
					equal = true
				} else {
					equal = false
					break
				}
			} else {
				equal = false
				break
			}
		}
		if equal {
			res += string(v)
		} else {
			break
		}
	}
	return res
}
func main() {
	strs := []string{"c", "acc", "ccc"}
	res := longestCommonPrefix(strs)
	println(res)
}
