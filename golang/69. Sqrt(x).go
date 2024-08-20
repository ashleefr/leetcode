func mySqrt(x int) int {
	res := 1
	for res*res < x {
		res++
	}

	for res*res > x {
		res--
	}
	return res
}