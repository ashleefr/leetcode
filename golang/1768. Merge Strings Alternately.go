func mergeAlternately(word1 string, word2 string) string {
	var res string

	if len(word1) < len(word2) {
		for i := 0; i < len(word1); i++ {
			res += string(word1[i])
			res += string(word2[i])
		}

		for i := len(word1); i < len(word2); i++ {
			res += string(word2[i])
		}
	} else {
		for i := 0; i < len(word2); i++ {
			res += string(word1[i])
			res += string(word2[i])
		}

		for i := len(word2); i < len(word1); i++ {
			res += string(word1[i])
		}
	}

	return res
}