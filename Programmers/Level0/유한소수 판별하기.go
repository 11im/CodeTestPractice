func solution(a int, b int) int {
	res := 1
	tmpB := b
	if a < b {
		a, tmpB = b, a
	}

	for tmpB > 0 {
		a, tmpB = tmpB, a%tmpB
	}

	b = b / a
	for {
		if b <= 1 {
			break
		}
		if b%2 != 0 && b%5 != 0 {
			res = 2
			break
		}
		if b%2 == 0 {
			b = b / 2
		}
		if b%5 == 0 {
			b = b / 5
		}
	}

	return res
}
