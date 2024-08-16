func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charMap := make(map[rune]int)

	for _, char := range s {
		charMap[char]++
	}
	for _, char := range t {
		charMap[char]--

		if charMap[char] < 0 {
			return false
		}

	}
	return true
}