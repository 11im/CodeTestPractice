func solution(spell []string, dic []string) int {
	res := 2
	tmp := make([]int, len(spell))
	flag := true

	for _, myString := range dic {
		flag = true
		for _, myChar := range myString {
			for idx, alpha := range spell {
				if string(myChar) == alpha {
					tmp[idx]++
				}
			}
		}
		for i, val := range tmp {
			if val != 1 {
				flag = false
			}
			tmp[i] = 0
		}
		if flag {
			res = 1
		}
	}
	return res
}