package main

func firstUniqChar(s string) int {
    smap := make(map[byte]int)
    for i := 0; i< len(s);i++{
        smap[s[i]]++
    }

    for i := 0; i< len(s); i++{
        if smap[s[i]] == 1{
            return i
        }
    }
    return -1
}