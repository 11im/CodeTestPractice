func solution(common []int) int {
	res := 0
	last := len(common) - 1

	flag, num := isCD(common[0], common[1], common[2])

	if flag {
		res = common[last] + num
	} else {
		res = common[last] * num
	}

	return res
}

func isCD(a, b, c int) (bool, int) {
	if b-a == c-b {
		return true, b - a
	} else {
		return false, b / a
	}
}