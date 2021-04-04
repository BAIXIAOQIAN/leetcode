package main

import "fmt"

func main(){
	input := [][]byte{
		{'X','O','X','O','X','O'},
		{'O','X','O','X','O','X'},
		{'X','O','X','O','X','O'},
		{'O','X','O','X','O','X'}}
	solve(input)

	for _, v := range input {
		fmt.Printf("%+v\n",v)
	}
}

func solve(board [][]byte)  {
    m := len(board)
    n := len(board[0])

    if m < 3 || n <3 {
        return
    }

    var dfs func(i,j int)
    dfs = func(i,j int){
        if i < 0 || j < 0 || i >=m || j >= n || board[i][j] != 'O'{
            return
        }

        board[i][j] = 'A'
        
        dfs(i-1,j)
        dfs(i+1,j)
        dfs(i,j-1)
        dfs(i,j+1)
    }

    //判断与边界上的O连通的O,并修改为A
    for i := 0; i < n; i++ {
        if board[0][i] == 'O'{
            dfs(0,i)
			fmt.Printf("0,%v\n",i)
        }

        if board[m-1][i] == 'O'{
            dfs(m-1,i)
			fmt.Printf("%v,%v\n",m-1,i)
        }
    }

    for i := 0; i < m; i++ {
        if board[i][0] == 'O'{
            dfs(i,0)
        }

        if board[i][n-1] == 'O'{
            dfs(i,n-1)
        }
    }

    for i := 0; i<m ;i++ {
        for j :=0 ;j<n; j++ {
            if board[i][j] == 'A' {
                board[i][j] = 'O'
            }else if board[i][j] == 'O' {
                board[i][j] = 'X'
            }
        }
    }
}