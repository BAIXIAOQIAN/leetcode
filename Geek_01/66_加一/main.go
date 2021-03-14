package main

func plusOne(digits []int) []int {
	n := len(digits)
	digits[n-1] = digits[n-1] + 1

	for i := n - 1; i >= 0; i-- {
		if i != 0 && digits[i] >= 10 {
			digits[i] = 10 - digits[i]
			digits[i-1] = digits[i-1] + 1
		}
	}

	if digits[0] >= 10 {
		digits[0] = 10 - digits[0]
		digits = append([]int{1}, digits[:]...)
	}
	return digits
}
