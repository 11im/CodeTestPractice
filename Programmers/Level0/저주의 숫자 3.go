func solution(n int) int {

	res := 0
	i := 1
	var tmp []int

	for {
		if i%3 != 0 && have3(i) {
			tmp = append(tmp, i)
		}
		if len(tmp) == n {
			break
		}
		i++
	}

	res = tmp[n-1]
	return res
}

func have3(n int) bool {
	flag := true

	for n > 0 {
		if n%10 == 3 {
			flag = false
		}
		n = n / 10
	}
	return flag
}