import "strconv"

func solution(quiz []string) []string {
	var res []string
	tmp := make([]string, 5)
	var chr string
	var x, y, z, cnt int64
	var flag bool

	for _, q := range quiz {

		cnt = 0
		q += " "

		for _, val := range q {
			if string(val) == " " {
				tmp[cnt] = string(chr)
				cnt++
				chr = ""
			} else {
				chr += string(val)
			}
		}

		x, _ = strconv.ParseInt(tmp[0], 10, 64)
		y, _ = strconv.ParseInt(tmp[2], 10, 64)
		z, _ = strconv.ParseInt(tmp[4], 10, 64)

		if tmp[1] == "-" {
			flag = x-y == z
		} else {
			flag = x+y == z
		}

		if flag {
			res = append(res, "O")
		} else {
			res = append(res, "X")
		}

	}

	return res
}