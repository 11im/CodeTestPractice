func solution(id_pw []string, db [][]string) string {
	res := "fail"

	for _, pair := range db {
		if pair[0] == id_pw[0] {
			if pair[1] == id_pw[1] {
				res = "login"
			} else {
				res = "wrong pw"
			}
			break
		}
	}

	return res
}