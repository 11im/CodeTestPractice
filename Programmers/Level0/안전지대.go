func solution(board [][]int) int {
	res := 0

	max := len(board)
	dx := []int{-1, 1, 0, 0, -1, -1, 1, 1}
	dy := []int{0, 0, 1, -1, 1, -1, 1, -1}

	var bomb [][]int

	for y, arr := range board {
		for x, here := range arr {
			if here == 1 {
				bomb = append(bomb, []int{x, y})
			}
		}
	}
	fmt.Println(bomb)
	bombx := 0
	bomby := 0
	for _, here := range bomb {
		for i := 0; i < len(dx); i++ {
			bombx = dx[i] + here[0]
			bomby = dy[i] + here[1]
			if bombx >= 0 && bombx < max && bomby >= 0 && bomby < max {
				board[bomby][bombx] = 1
			}
		}
	}

	for _, arr := range board {
		for _, here := range arr {
			if here == 0 {
				res++
			}
		}
	}

	return res
}
