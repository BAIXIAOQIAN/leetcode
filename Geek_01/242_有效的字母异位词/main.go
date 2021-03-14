package main

func isAnagram(s string, t string) bool {
	sr, tr := []rune(s), []rune(t)

	srMap := make(map[string]int)
	for _, v := range sr {
		srMap[string(v)]++
	}

	var length int

	trMap := make(map[string]int)
	for _, v := range tr {
		trMap[string(v)]++
	}

	for k, v := range srMap {
		if val, ok := trMap[k]; ok && val == v {
			length++
		}
	}

	return length == len(srMap) && len(s) == len(t)
}
