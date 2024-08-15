func isPalindrome(x int) bool {
	if (x < 0) || (x != 0 && x%10 == 0) {
		return false
	}

	reversed := 0
	original := x

	for x > 0 {
		digit := x % 10
		reversed = reversed*10 + digit
		x = x / 10
	}

	return original == reversed
}