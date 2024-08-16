func findTheDifference(s string, t string) byte {
	sumS, sumT := 0, 0
	for _, i := range s {
		sumS += int(i)
	}

	for _, i := range t {
		sumT += int(i)
	}

	diff := sumT - sumS

	return byte(diff)
}