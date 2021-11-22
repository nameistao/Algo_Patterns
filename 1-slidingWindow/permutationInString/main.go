package main

//Question: permutation in a string
//Time: O(N+M)
//Space: O(M)
func findPermutation(str string, pattern string) bool {
	pMap := make(map[rune]int)
	for _, v := range pattern {
		_, found := pMap[v]
		if found {
			pMap[v]++
		} else {
			pMap[v] = 1
		}
	}

	foundLetter := 0

	for i, v := range str {
		_, found := pMap[v]

		if found {
			pMap[v]--
			if pMap[v] == 0 {
				foundLetter++
			}
		}

		if i >= len(pattern) {
			_, found := pMap[rune(str[i-len(pattern)])]
			if found {
				if pMap[rune(str[i-len(pattern)])] == 0 {
					foundLetter--
				}
				pMap[rune(str[i-len(pattern)])]++
			}
		}

		if len(pMap) == foundLetter {
			return true
		}
	}

	return false
}
