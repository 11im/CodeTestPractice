func solution(keyinput []string, board []int) []int {
	res := make([]int, 2)
	commands := []string{"up", "down", "right", "left"}
	change := [][]int{[]int{0, 1}, []int{0, -1}, []int{1, 0}, []int{-1, 0}}

	xLimit := []int{board[0] / 2, -board[0] / 2}
	yLimit := []int{board[1] / 2, -board[1] / 2}
	for _, command := range keyinput {
		for idx, val := range commands {
			if val == command {
				res[0] += change[idx][0]
				res[1] += change[idx][1]
			}
		}
		if res[0] > xLimit[0] {
			res[0] = xLimit[0]
		} else if res[0] < xLimit[1] {
			res[0] = xLimit[1]
		}

		if res[1] > yLimit[0] {
			res[1] = yLimit[0]
		} else if res[1] < yLimit[1] {
			res[1] = yLimit[1]
		}
	}

	return res
}
