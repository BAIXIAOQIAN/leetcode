package main

func numDecodings(s string) int {
    if s[0:1] == "0" {
        return 0
    }

    cur, pre := 1,1

    for i:=1; i < len(s); i++ {
        tmp := cur
        if s[i:i+1] == "0" {
            if s[i-1:i] == "1" || s[i-1:i] == "2" {
                cur = pre
            }else {
                return 0
            }
        }else if s[i-1:i] == "1" {
            cur = cur + pre
        }else if s[i-1:i] == "2" {
            if s[i:i+1] <= "6" {
                cur = cur + pre
            }
        } 
        pre = tmp
    }
    return cur
}