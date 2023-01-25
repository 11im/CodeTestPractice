func solution(A string, B string) int {
	res := -1
	idx := 0
	// i는 민 숫자, j는 검사용 인덱스
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A); j++ {
			idx = i + j
			if idx >= len(A) {
				idx -= len(A)
			}
			if A[j] != B[idx] {
				break
			}
			if j == len(A)-1 {
				res = i
				return res
			}
		}
	}

	return res
}
