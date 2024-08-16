func maxDistance(arrays [][]int) int {
	min, max := arrays[0][0], arrays[0][len(arrays[0])-1]
	result := 0
	for _, arr := range arrays[1:] {
		localMin, localMax := arr[0], arr[len(arr)-1]
		result = getMax(result, getMax(localMax-min, max-localMin))
		min = getMin(min, localMin)
		max = getMax(max, localMax)
	}

	return result
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}