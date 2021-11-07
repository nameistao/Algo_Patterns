package main

import "fmt"

func main() {
	fmt.Println("Expected: abdec")
	fmt.Println("Actual:", findSubstring("aabdec", "abc"))
	fmt.Println("Expected: bca")
	fmt.Println("Actual:", findSubstring("abdbca", "abc"))
	fmt.Println("Expected: ")
	fmt.Println("Actual:", findSubstring("adcad", "abc"))
}

//Question: smallet window containing substring
//Time: O(N+M)
//Space: O(M)
func findSubstring(str string, pattern string) string {
	pMap := make(map[rune]int)
	for _, v := range pattern {
		_, found := pMap[v]
		if found {
			pMap[v]++
		} else {
			pMap[v] = 1
		}
	}

	foundPattern, start, answerLen, answer := 0, 0, int(^uint(0)>>1), ""

	for i, v := range str {
		_, found := pMap[v]
		if found {
			pMap[v]--
			if pMap[v] == 0 {
				foundPattern++
			}
		}

		for len(pMap) == foundPattern {
			if i-start+1 < answerLen {
				answerLen = i - start + 1
				answer = str[start : i+1]
			}
			_, found := pMap[rune(str[start])]
			if found {
				if pMap[rune(str[start])] == 0 {
					foundPattern--
				}
				pMap[rune(str[start])]++
			}
			start++
		}

	}

	return answer
}
