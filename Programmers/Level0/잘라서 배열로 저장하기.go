func solution(my_str string, n int) []string {

	var res []string
	idx := 0

	for i := 0; i < len(my_str); i += n {
		idx = i + n
		if idx > len(my_str) {
			idx = len(my_str)
		}

		res = append(res, my_str[i:idx])
	}

	return res
}