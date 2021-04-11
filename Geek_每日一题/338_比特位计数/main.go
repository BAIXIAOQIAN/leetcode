package main

func main(){
	res := countBits(2)
	println(res)
}

func countBits(num int) []int {
    res := make([]int,0)

    tmp := 0
    for num-tmp >= 0 {
        if tmp == 0 {
            res = append(res,0)
        }else {
            var count int
			println(tmp)
            for tmp > 0 {
                tmp = tmp&(tmp-1)
                count++
            }
            res = append(res,count)
        }
        tmp++
    }
    return res
}