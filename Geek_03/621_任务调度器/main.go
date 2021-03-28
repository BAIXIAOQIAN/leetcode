package main

func main(){
	input := []byte{'A','A','A','B','B','B','C','C','C'}
	n := 2

	res := leastInterval(input,n)
	println(res)
}

func leastInterval(tasks []byte, n int) int {
    if n == 0 {
        return len(tasks)
    }

    tMap := make(map[byte]int)
    for _, v := range tasks {
        tMap[v]++
    }

    var res int
    cMap := make(map[int]int)
    for _, v := range tMap {
        cMap[v]++
        if v > res {
            res = v
        }
    }

    count := cMap[res]
    k := res

    time := (k-1)*(n+1)+count
    if time < len(tasks) {
        return len(tasks)
    }
    return  time
}