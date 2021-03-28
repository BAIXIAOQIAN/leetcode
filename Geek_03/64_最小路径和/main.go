package main

func main(){
	input := [][]int{{1,3,1},{1,5,1},{4,2,1}}
	res := minPathSum(input)
	println(res)
}

func minPathSum(grid [][]int) int {
    n := len(grid)
    m := len(grid[0])

    for j := 1; j <m; j++{
        grid[0][j] = grid[0][j] + grid[0][j-1]
    }

    for i := 1; i<n; i++ {
        grid[i][0] = grid[i][0] + grid[i-1][0]
    }

    for i := 1; i <n;i++{
        for j:=1;j<m;j++{
            grid[i][j] = min(grid[i-1][j],grid[i][j-1]) + grid[i][j]
        }
    }
    return grid[n-1][m-1]
}

func min(a,b int)int{
    if a<b{
        return a
    }
    return b
}