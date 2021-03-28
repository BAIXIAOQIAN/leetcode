package main

func main(){
	text1 := "abcde"
	text2 := "ace"

	res := longestCommonSubsequence(text1,text2)
	println(res)
}

func longestCommonSubsequence(text1 string, text2 string) int {
    n := len(text1)
    m := len(text2)
	println(n,m)

    dp := make([][]int,n)
    for i := 0; i< n; i++ {
        dp[i] = make([]int,m)
    }

    for i := 0; i<n;i++{
        for j := 0; j<m; j++ {
			if i ==0 || j == 0 {
				if text1[i] == text2[j] {
					dp[i][j] = 1
				}else {
					if i > 0 {
						dp[i][j] = dp[i-1][j]
					}else {
						dp[i][j] = dp[i][j-1]
					}
				}
			}else {
				if text1[i] == text2[j] {
					dp[i][j] = max(dp[i-1][j],dp[i][j-1]) + 1
				}else {
					dp[i][j] = max(dp[i-1][j],dp[i][j-1])
				}
			}
        }
    }

    return dp[n-1][m-1]
}

func max(a,b int)int {
    if a>b {
        return a
    }
    return b
}