package main

func main(){

}


//1.baoli
func countSubstrings(s string) int {
    str := []byte(s)
    var res int

    for i := 0; i<len(str);i++ {
        for j := i; j<len(str);j++ {
            if ishw(str[i:j+1]){
                res++
            }
        }
    }
    return res
}

func ishw(s []byte)bool {
    l,r := 0,len(s)-1

    for l <= r {
        if s[l] != s[r] {
            return false
        }
        l++
        r--
    }
    return true
}

//2.动态规划
func countSubstrings2(s string) int {
    dp := make([][]bool,len(s))
    for i := 0; i<len(s); i++ {
        dp[i] = make([]bool,len(s))
    }

    var res int
    for i := 0; i<len(s); i++ {
        for j := 0; j<=i;j++ {
            if i== j {
                dp[j][i] = true
                res++
            }else if i-j ==1 && s[i] == s[j] {
                dp[j][i] = true
                res++
            }else if i-j > 1 && s[i] == s[j] && dp[j+1][i-1] {
                dp[j][i] = true
                res++
            }
        }
    }
    return res
} 