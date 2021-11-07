package main

import "fmt"

func main() {
	fmt.Println("Expected: [1 2]")
	fmt.Println("Actual:", findStringAnagrams("ppqp", "pq"))
	fmt.Println("Expected: [2 3 4]")
	fmt.Println("Actual:", findStringAnagrams("abbcabc", "abc"))
}

//Question: string anagrams
//Time: O(N+M)
//Space: O(M)
func findStringAnagrams(str string, pattern string) []int {
	pMap := make(map[rune]int)
	for _, v := range pattern {
		_, found := pMap[v]
		if found {
			pMap[v]++
		} else {
			pMap[v] = 1
		}
	}

	answer := make([]int, 0)
	foundPattern := 0

	for i, v := range str {
		_, found := pMap[v]
		if found {
			pMap[v]--
			if pMap[v] == 0 {
				foundPattern++
			}
		}

		if i >= len(pattern) {
			_, found := pMap[rune(str[i-len(pattern)])]
			if found {
				if pMap[rune(str[i-len(pattern)])] == 0 {
					foundPattern--
				}
				pMap[rune(str[i-len(pattern)])]++
			}
		}

		if len(pattern) == foundPattern {
			answer = append(answer, i-len(pattern)+1)
		}
	}
	return answer
}
