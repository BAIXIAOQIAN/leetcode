package main

//双指针
func merge(nums1 []int, m int, nums2 []int, n int) {
	mp, np, p := m-1, n-1, m+n-1
	for mp >= 0 && np >= 0 {
		if nums1[mp] > nums2[np] {
			nums1[p] = nums1[mp]
			mp--
		} else {
			nums1[p] = nums2[np]
			np--
		}
		p--
	}

	for np >= 0 {
		nums1[p] = nums2[np]
		p--
		np--
	}
}
