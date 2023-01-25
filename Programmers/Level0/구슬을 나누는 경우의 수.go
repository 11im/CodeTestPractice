func solution(balls int, share int) int {

	var res float64 = 1
	var devide float64 = 1

	for i := 0; i < share; i++ {
		res = res * float64((balls - i))
	}

	for i := share; i > 0; i-- {
		devide = devide * float64(i)
	}

	res = res / devide
	return int(res)
}