func solution(array []int) int {
	res := 0
	for _, val := range array {
		for val > 0 {
			if val%10 == 7 {
				res++
			}
			val = val / 10
		}
	}
	return res
}