package main

func permute(nums []int) [][]int {
	var res [][]int
	visited := make(map[int]bool)

	var dfs func(path []int)
	dfs = func(path []int) {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for _, v := range nums {
			if visited[v] {
				continue
			}

			visited[v] = true
			path = append(path, v)
			dfs(path)
			path = path[:len(path)-1]
			visited[v] = false
		}
	}
	dfs([]int{})
	return res
}
