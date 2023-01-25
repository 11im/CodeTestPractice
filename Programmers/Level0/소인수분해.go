func solution(n int) []int {
	var res []int
	tmp := 0
	for n >= 2 {
		tmp = getFirst(n)
		flag := true
		for _, val := range res {
			if val == tmp {
				flag = false
			}
		}
		if flag {
			res = append(res, tmp)
		}
		n = n / tmp
	}
	return res
}

func getFirst(n int) int {
	res := 0
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			res = i
			break
		}
	}
	return res
}