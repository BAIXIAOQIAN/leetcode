package main

func isIsomorphic(s string, t string) bool {
    h1 := make(map[byte]int)
    h2 := make(map[byte]int)

    for i:= 0; i<len(s); i++ {
        h1[s[i]] = i
        h2[t[i]] = i
    }

    for i := 0; i<len(s); i++ {
        if h1[s[i]] != h2[t[i]] {
            return false
        }
    }
    return true
}