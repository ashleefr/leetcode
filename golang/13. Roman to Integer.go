func romanToInt(s string) int {
	roman := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	res := 0
	lenght := len(s)
	for i := 0; i < lenght; i++ {
		if i < lenght-1 && roman[s[i]] < roman[s[i+1]] {
			res = res - roman[s[i]]
		} else {
			res = res + roman[s[i]]
		}
	}

	return res
}