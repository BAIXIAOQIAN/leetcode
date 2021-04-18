package main

func main(){
	b,e := "a","c"
	w := []string{"a","b","c","lot","log","cog"}
	println(ladderLength(b,e,w))
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
    n := len(beginWord)
    wordMap := make(map[string]bool)

    for _, v := range wordList{
        wordMap[v] = true
    }

    var resCount []int

    var dfs func(beginWord, endWord string,count int, wordMap map[string]bool)

    dfs = func(beginWord, endWord string,count int,wordMap map[string]bool){
		println(count,beginWord,endWord)
        if beginWord == endWord {
            resCount = append(resCount,count)
            return
        }

        for i := 0 ; i<n; i++ {
            tmp := beginWord
            for j := 'a'; j <= 'z'; j++ {
                var nextWord string
                 if i == n-1 {
                    nextWord = beginWord[0:i] + string(j) 
                }else {
                    nextWord = beginWord[0:i] + string(j) + beginWord[i+1:n]
                }
                if wordMap[nextWord]{
					tmpMap := make(map[string]bool)
					for k, v := range wordMap{
						tmpMap[k] = v
					}
					tmpMap[nextWord] = false
                    dfs(nextWord,endWord,count+1,tmpMap)
                }
            }
            beginWord = tmp
        }
	}
    dfs(beginWord,endWord,0,wordMap)
    min := 0
    for i := 0; i<len(resCount); i++ {
		println(resCount[i])
        if min < resCount[i] {
            min = resCount[i]
        }
    }
    return min
}