func solution(dots [][]int) int {
	res := 0

	//case1 1,2 / 3,4
	line1 := getDiff(dots[0], dots[1])
	line2 := getDiff(dots[2], dots[3])

	if isPar(line1, line2) {
		return 1
	}

	//case2 1,3 / 2,4

	line1 = getDiff(dots[0], dots[2])
	line2 = getDiff(dots[1], dots[3])

	if isPar(line1, line2) {
		return 1
	}

	//case3 1,4 / 2,3
	line1 = getDiff(dots[0], dots[3])
	line2 = getDiff(dots[1], dots[2])

	if isPar(line1, line2) {
		return 1
	}

	return 0
}

// x,y 증가량
func getDiff(dot1, dot2 []int) []int {
	line := []int{dot2[0] - dot1[0], dot2[1] - dot1[1]}
	return line
}

func isPar(line1, line2 []int) bool {
	if line2[1]/line2[0] == line1[1]/line1[0] {
		return true
	}
	return false

}
