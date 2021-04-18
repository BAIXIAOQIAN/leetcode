package main

func main(){
	s := "a-bC-dEf-ghIj"

	res := reverseOnlyLetters(s)
	println(res)
}

func reverseOnlyLetters(S string) string {
    l,r := 0,len(S)-1

    bs := []byte(S)

    for l < r {
		println(bs[l],bs[r],'a','Z',string([]byte{97}))
        if bs[l] < 'a' || bs[l] > 'z' {
            l++
            continue
        }

        if bs[r] < 'a' || bs[r] > 'z' {
            r--
            continue
        }

        bs[l],bs[r] = bs[r],bs[l]
        l++
        r--
    }
    return string(bs)
}