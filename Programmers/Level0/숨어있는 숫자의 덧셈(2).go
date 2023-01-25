import "strconv"

func solution(my_string string) int {
	res := 0
	tmp := ""
	for i := 0; i < len(my_string); i++ {
		val := string(my_string[i])
		_, err := strconv.Atoi(val)
		if err != nil {
			num, _ := strconv.Atoi(tmp)
			res += num
			tmp = ""
		} else {
			tmp += val
		}
		if err == nil && i == len(my_string)-1 {
			num, _ := strconv.Atoi(tmp)
			res += num
		}
	}
	return res
}