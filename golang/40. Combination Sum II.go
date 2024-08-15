func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	sort.Ints(candidates)
	backtrack(candidates, target, []int{}, 0, &res)
	return res
}

func backtrack(candidates []int, target int, current []int, start int, res *[][]int) {
	if target == 0 {
		// Make a copy of the current combination since we're going to modify it
		temp := make([]int, len(current))
		copy(temp, current)
		*res = append(*res, temp)
		return
	}

	for i := start; i < len(candidates); i++ {
		// Skip duplicate candidates
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}

		if candidates[i] > target {
			break
		}

		current = append(current, candidates[i])
		backtrack(candidates, target-candidates[i], current, i+1, res)
		// Backtrack, remove the last element
		current = current[:len(current)-1]
	}
}