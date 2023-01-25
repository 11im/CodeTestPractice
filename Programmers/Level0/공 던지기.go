func solution(numbers []int, k int) int {
	res := 0
	i := 0
	leng := len(numbers)

	for k > 0 {
		res = numbers[i%leng]
		k--
		i += 2
	}
	return res
}