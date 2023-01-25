import "strconv"

func solution(my_string string) int {
	res := 0
	cal := true
	tmp := ""

	for i := 0; i < len(my_string); i++ {
		// + or -
		if tmp == "-" {
			cal = false
			tmp = ""
		} else if tmp == "+" {
			cal = true
			tmp = ""
		}

		if string(my_string[i]) == " " || i == len(my_string)-1 {
			if i == len(my_string)-1 {
				tmp += string(my_string[i])
			}
			num, _ := strconv.Atoi(tmp)
			if cal {
				res += num
				tmp = ""
			} else {
				res -= num
				tmp = ""
			}
		} else {
			tmp += string(my_string[i])
		}
	}

	return res
}