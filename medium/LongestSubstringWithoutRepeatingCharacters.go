package medium

func lengthOfLongestSubstring(s string) int {
	max_str := 0
	count := 0
	alpha := make(map[rune]struct{})

	for j := 0; j < len(s); j++ {
		for i := j; i < len(s); i++ {

			chr := rune(s[i])

			_, ok := alpha[chr]

			if ok {
				if max_str < count {
					max_str = count
				}
				alpha = make(map[rune]struct{})
				count = 0
				break
			} else {
				alpha[chr] = struct{}{}
				count++
			}
		}
	}
	if max_str < count {
		max_str = count
	}

	return max_str
}
