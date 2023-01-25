func solution(dots [][]int) int {
	var res int
	xDots := []int{dots[0][0], dots[0][0]}
	yDots := []int{dots[0][1], dots[0][1]}

	for _, arr := range dots {
		if arr[0] < xDots[0] {
			xDots[0] = arr[0]
		} else if arr[0] > xDots[1] {
			xDots[1] = arr[0]
		}

		if arr[1] < yDots[0] {
			yDots[0] = arr[1]
		} else if arr[1] > yDots[1] {
			yDots[1] = arr[1]
		}
	}
	res = (xDots[1] - xDots[0]) * (yDots[1] - yDots[0])

	return res
}
