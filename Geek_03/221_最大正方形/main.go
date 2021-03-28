package main

func main(){
	input := [][]byte{
		{'0','1'},
		{'1','0'},
	}
	res := maximalSquare(input)
	println(res)
}

func maximalSquare(matrix [][]byte) int {
    m,n,res := len(matrix),len(matrix[0]),0

    dp := make([][]int,m)
    for i:=0;i<m;i++{
        dp[i] = make([]int,n)
    }

    for i := 0; i< n; i++ {
        if matrix[0][i] == '1' {
             dp[0][i] = 1
        }
        if res < dp[0][i] {
            res = dp[0][i]
        }
    }

    for j := 0; j<m; j++ {
        if matrix[j][0] == '1' {
            dp[j][0] = 1
        }
        if res < dp[j][0] {
            res = dp[j][0]
        }
    }

    for i:=1;i<m;i++{
        for j := 1; j<n;j++{
            if matrix[i][j] == '1' {
                dp[i][j] = min(min(dp[i-1][j-1],dp[i][j-1]),dp[i-1][j]) +1
                if res < dp[i][j] {
                    res = dp[i][j]
                }
            }
        }
    }
    return res*res
}

func min(a,b int) int{
    if a < b {
        return a
    }
    return b
}