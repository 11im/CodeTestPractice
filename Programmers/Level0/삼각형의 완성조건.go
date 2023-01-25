func solution(sides []int) int {
	res := 0
	if sides[0] < sides[1] {
		sides[1], sides[0] = sides[0], sides[1]
	}

	min := sides[0] - sides[1] + 1

	for i := min; i <= sides[0]; i++ {
		res++
	}

	max := sides[0] + sides[1]
	for i := sides[0] + 1; i < max; i++ {
		res++
	}
	return res
}
