package main

func main(){
	s := " -42"

	count := myAtoi(s)
	println(count)
}

func myAtoi(s string) int {
    fh := 1
    min,max := -1 * (1 << 31), (1<<31) -1

    start := false
    digits := false
    count := 0
    for i := 0; i< len(s); i++{
        if s[i] == ' ' {
            if start {
                break
            }
            continue
        }

        if s[i] == '-' && !digits{
            if start {
                break
            }
            digits = true
            fh = -1
            continue
        }

        if s[i] == '+' && !digits{
            if start {
                break
            }
            digits = true
            continue
        }

        if s[i] >= '0' && s[i] <= '9' {
            start = true
			if count >= max {
				count = max
				break
			}
            count = count*10 + int(s[i]-'0')
        }else {
            break
        }
	}

	println(min)
	count = fh * count
	if count <= min {
		count = min
	}

    return count
}