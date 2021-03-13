func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	n := len(nums)

	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target := -1 * nums[i]
		k := n - 1
		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			for j < k && nums[j]+nums[k] > target {
				k--
			}

			if j == k {
				break
			}

			if nums[j]+nums[k] == target {
				res = append(res, []int{nums[i], nums[j], nums[k]})
			}
		}
	}
	return res
}