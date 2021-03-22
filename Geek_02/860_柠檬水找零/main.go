package main

func main() {
	lemonadeChange([]int{})
}

func lemonadeChange(bills []int) bool {
	five, ten := 0, 0

	for _, v := range bills {
		if v == 5 {
			five++
			continue
		}

		if v == 10 {
			if five < 1 {
				return false
			}
			five--
			ten++
			continue
		}

		if v == 20 {
			if ten > 0 && five > 0 {
				ten--
				five--
			} else if five >= 3 {
				five = five - 3
			} else {
				return false
			}
		}
	}
	return true
}
