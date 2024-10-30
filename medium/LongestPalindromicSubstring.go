package medium

func IsPalindrome(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func LongestPalindrome(s string) string {
	var maxLen, maxIndexJ int
	var maxIndexI int = -1

	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if maxLen < (j - i) {
				if s[i] == s[j] {
					if s[i+1] == s[j-1] {
						if IsPalindrome(s, i, j) {
							maxLen = (j - i)
							maxIndexI = i
							maxIndexJ = j
						}
					}
				}
			}
		}
	}

	if maxIndexI == -1 {
		return string(s[0])
	}

	// Парадоксально, но это убыстряет код на 10мс
	var result []rune
	for i := maxIndexI; i <= maxIndexJ; i++ {
		result = append(result, rune(s[i]))
	}

	return s[maxIndexI : maxIndexJ+1]
}
