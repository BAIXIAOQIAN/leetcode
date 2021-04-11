package main

func main(){
	res := solveNQueens(8)
	for _, v := range res {
		for _,v1 := range v {
			println("queen: ",v1)
		}
		println()
	}
}

func solveNQueens(n int) [][]string {
    res := make([][]string,0)
    arrs := make([][]int,n)

    var dfs func(nums []int,row,cols,pie,na int)
    dfs = func(nums []int,row,cols,pie,na int){
        if row >= n {
            tmp := make([]int,n)
            copy(tmp,nums)
            arrs = append(arrs,tmp)
            return
        }

        bits := (^(cols | pie | na) & ((1 << n) - 1 )) //得到当前所有的空位

        for bits > 0 {
            p := bits & -bits //取到最低位的1
            bits = bits & (bits-1) //表示在p位置上放上皇后
			tmp := p &((1 << n) - 1 )
            dfs(append(nums,tmp),row+1,cols^p,(pie^p)<<1,(na^p)>>1)
        }
    }

	dfs([]int{},0,0,0,0)

    for _, v := range arrs {
       res = append(res,generalNqueen(v,n))
    }
    return res
}

func generalNqueen(nums []int, n int)[]string{
    var res []string
    for _, v := range nums{
        str := ""
        tmp := n
        for tmp > 0 {
            if v&1 == 0 {
                str = str + "."
            }else {
                str = str + "Q"
            }
			v = v >> 1
			tmp--
        }
        res = append(res,str)
    }
    return res
}