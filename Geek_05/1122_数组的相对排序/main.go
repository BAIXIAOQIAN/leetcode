package main

func relativeSortArray(arr1 []int, arr2 []int) []int {

    arr1Map := make(map[int]int,0)
    for _, v := range arr1 {
        arr1Map[v]++
    }

    for v,k := range arr1Map {
        if k > 1 {
            for k > 1 {
                arr2 = insertArr2(v,arr2)
                k--
            }
        }
    }

    return arr2
}

func insertArr2(n int, arr2 []int) []int{
    var res []int
    for i := 0; i<len(arr2); i++ {
        if n <= arr2[i] {
            res = append(res,arr2[:i]...)
            res = append(res,n)
            res = append(res,arr2[i:]...)
        }
    }
    return res
}