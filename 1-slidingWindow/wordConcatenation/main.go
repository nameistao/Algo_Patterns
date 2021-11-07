package main

import "fmt"

func main() {
	fmt.Println("Expected: [0 3]")
	fmt.Println("Actual:", findWordConcatenation("catfoxcat", []string{"cat", "fox"}))
	fmt.Println("Expected: [3]")
	fmt.Println("Actual:", findWordConcatenation("catcatfoxfox", []string{"cat", "fox"}))
}

func findWordConcatenation(str string, words []string) []int {
	m := make(map[string]int)
	for _,v := range words {
		_, found := m[v]
		if found {
			m[v]++
		} else {
			m[v] = 1
		}
	}

	foundCount, answer, wordLen, wordCount:= 0, make([]int, 0), len(words[0]), len(words)

	for i:=0; i<len(str) - wordLen * wordCount; i++ {
		_, found := m[str[i:i+wordLen]]
		if found {
			m[str[i:i+wordLen]]--
			if m[str[i:i+wordLen]] == 0 {
				foundCount++
			}
		}

		if i >= len(words)*wordLen {
			_, found := m[str[start: start+wordLen]]
			if found {
				if m[str[start: start+wordLen]] == 0 {
					foundCount--
				}
				m[str[start: start+wordLen]]++
			}
			start += len(words[0])
		}

		if foundCount == len(m) {
			answer = append(answer, start)
		}
	}

	return answer
}