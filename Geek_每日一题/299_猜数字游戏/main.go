package main

import(
	"strconv"
)

func main(){
	res := getHint("1807","7810")
	println(res)
}

func getHint(secret string, guess string) string {
    m1 := make([]int,10)
    m2 := make([]int,10)

    aCount,bCount := 0,0
    for i := 0; i<len(secret); i++ {
        if secret[i] == guess[i] {
            aCount++
        }else {
            m1[secret[i]-'0']++
            m2[guess[i]-'0']++
        }
    }

    for i := 0; i<len(m1); i++ {
        bCount += min(m1[i],m2[i])
		println(i)
    }
    return strconv.Itoa(aCount)+"A"+strconv.Itoa(bCount)+"B"
}

func min(a,b int)int{
    if a < b {
        return a
    }
    return b
}