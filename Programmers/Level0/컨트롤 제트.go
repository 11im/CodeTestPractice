import "strconv"

func solution(s string) int {
	res := 0
	last := 0
	tmp := ""

	for i := 0; i < len(s); i++ {
		if string(s[i]) == "Z" {
			res -= last
			continue
		}

		if s[i] != 32 {
			tmp += string(s[i])
		}

		if s[i] == 32 || i == len(s)-1 {
			num, _ := strconv.Atoi(tmp)
			res += num
			last = num
			tmp = ""
		}
	}
	return res
}