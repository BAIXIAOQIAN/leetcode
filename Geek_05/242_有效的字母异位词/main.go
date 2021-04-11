package main

func isAnagram(s string, t string) bool {
    smap := make(map[rune]int,0)

    for _, v := range s {
        smap[v]++
    }

    count := 0
    for _, v := range t {
        if val, ok := smap[v]; ok && val > 0 {
            count++
            smap[v]--
        }else {
            return false
        }
    }

    if count == len(s) {
        return true
    }
    return false
}