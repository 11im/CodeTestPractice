func solution(lines [][]int) int {

	res := 0
	common := make(map[float64]int)

	for _, val := range lines {
		for i := val[0]; i < val[1]; i++ {
			line := float64(i) + 0.5
			_, err := common[line]
			if err {
				common[line]++
			} else {
				common[line] = 1
			}
		}
	}

	for _, cnt := range common {
		if cnt > 1 {
			res++
		}
	}
	return res
}
