package main

func merge(intervals [][]int) [][]int {
    //排序
    sort.Slice(intervals,func(i,j int)bool{
        return intervals[i][0] < intervals[j][0]
    })

    res := make([][]int,0)
    cs := make([]int,2)

    //合并
    for _, v := range intervals {
        if len(res) == 0 || cs[1] < v[0] {
            res = append(res,v)
            cs = v
        }else if cs[1] < v[1]{
            cs[1] = v[1]
        }
    }
    return res
}
