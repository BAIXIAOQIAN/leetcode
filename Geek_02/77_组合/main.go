package main

import "fmt"

func main() {
	res := combine(4, 2)
	for _, v := range res {
		fmt.Println(v)
	}
}

func combine(n int, k int) [][]int {
	if n < k {
		return nil
	}

	nums := make([]int, 0)

	for n > 0 {
		nums = append(nums, n)
		n--
	}

	res := make([][]int, 0)

	var gen func(array []int, k int, r []int)

	gen = func(array []int, k int, r []int) {
		if k == 0 {
			tmp := make([]int, len(r))
			copy(tmp, r)
			res = append(res, tmp)
			return
		}

		k--
		for i, v := range array {
			if isExit(r, v) {
				continue
			}
			tmp := r
			r = append(r, v)
			gen(array[i:], k, r)
			r = tmp
		}
	}

	gen(nums, k, []int{})
	return res
}

func isExit(arr []int, n int) bool {
	for _, v := range arr {
		if v == n {
			return true
		}
	}
	return false
}
