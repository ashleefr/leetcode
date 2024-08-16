func strStr(haystack string, needle string) int {
	res := -1

	for i, char := range haystack {
		if byte(char) == needle[0] && i+len(needle) <= len(haystack) {
			ok := true
			for j, need := range needle {
				if byte(need) != haystack[i+j] {
					ok = false
				}
			}

			if ok {
				return i
			}
		}
	}
	return res
}