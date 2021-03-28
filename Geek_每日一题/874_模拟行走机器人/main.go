package main

func main(){
	arr := make([][]int,0)
	arr = append(arr,[]int{2,4})

	a := []int{4,-1,4,-2,4}
	res := robotSim(a,arr)
	println(res)
}	

func robotSim(commands []int, obstacles [][]int) int {
    dx := []int{0,1,0,-1}
    dy := []int{1,0,-1,0}

    obMap := make(map[[2]int]bool)

    for _, v :=range obstacles {
        if len(v) > 1 {
            obMap[[2]int{v[0],v[1]}] = true
        }
    }

    var res,dir int
	var x, y int

    for _, v := range commands {
        if v == -1 {
            dir = (dir+1)%4
        }else if v == -2 {
            dir = (dir+3)%4
        }else {
            for i := 0; i<v;i++ {
                if obMap[[2]int{x+dx[dir],y+dy[dir]}] {
                    break
                }
                x = x + dx[dir]
                y = y + dy[dir]
            }
			println(x,y,dx[dir],dy[dir])
            if x*x + y*y > res {
                res = x*x + y*y
            }
        }
    }
    return res
}