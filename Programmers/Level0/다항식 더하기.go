import "strconv"

func solution(polynomial string) string {
	res := ""
	var tmp []string
	var tmpChar = ""
	coef := []int{0, 0}

	for i := 0; i < len(polynomial); i++ {
		if string(polynomial[i]) == " " && tmpChar != "" {
			tmp = append(tmp, tmpChar)
			tmpChar = ""
		} else if string(polynomial[i]) != " " && string(polynomial[i]) == "+" {
			continue
		} else if i == len(polynomial)-1 {
			tmpChar += string(polynomial[i])
			tmp = append(tmp, tmpChar)
		} else if string(polynomial[i]) != " " {
			tmpChar += string(polynomial[i])
		}
	}

	for _, val := range tmp {
		num, err := strconv.Atoi(val)
		if err != nil {
			if val == "x" {
				coef[0] += 1
			} else {
				num, _ = strconv.Atoi(val[0 : len(val)-1])
				coef[0] += num
			}
		} else {
			coef[1] += num
		}
	}

	if coef[0] > 1 {
		res += strconv.Itoa(coef[0]) + "x"
	} else if coef[0] == 1 {
		res += "x"
	}

	if coef[0] > 0 && coef[1] > 0 {
		res += " " + "+" + " "
	} else if coef[0] == 0 && coef[1] == 0 {
		res += "0"
	}

	if coef[1] > 0 {
		res += strconv.Itoa(coef[1])
	}

	return res
}