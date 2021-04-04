package main

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	res := ladderLength(beginWord, endWord, wordList)
	println(res)
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordMap := make(map[string]bool, len(wordList))
	for _, w := range wordList {
		wordMap[w] = true
	}

	//起始词加入queue
	queue := []string{beginWord}
	level := 1
	for len(queue) != 0 {
		println(queue)
		ls := len(queue)
		for i := 0; i < ls; i++ {
			word := queue[0]
			queue = queue[1:]

			if word == endWord {
				return level
			}

			for c := 0; c < len(word); c++ {
				for j := 'a'; j <= 'z'; j++ {
					newWord := word[:c] + string(j) + word[c+1:]
					if wordMap[newWord] {
						queue = append(queue, newWord)
						wordMap[newWord] = false
					}
				}
			}
		}
		level++
	}
	return 0
}
