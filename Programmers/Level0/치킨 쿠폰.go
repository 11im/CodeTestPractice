func solution(chicken int) int {
	res := 0
	for i := 1; i <= chicken; i++ {
		if i%10 == 0 {
			chicken++
			res++
		}
	}
	return res
}