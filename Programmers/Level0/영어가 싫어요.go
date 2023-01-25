
func solution(numbers string) int64 {
	numstrs := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	tmp := ""
	var res int64 = 0

	for i := 0; i < len(numbers); i++ {
		tmp += string(numbers[i])
		for idx, val := range numstrs {
			if val == tmp {
				res = res*10 + int64(idx)
				tmp = ""
			}
		}
	}

	return res
}