import "strconv"

func solution(bin1 string, bin2 string) string {
	res := ""
	if bin1 == "0" && bin2 == "0" {
		res += "0"
		return res
	}
	var intBin1, intBin2 int64
	intBin1, _ = strconv.ParseInt(bin1, 2, 64)
	intBin2, _ = strconv.ParseInt(bin2, 2, 64)
	intBin1 += intBin2
	for intBin1 > 0 {
		res = strconv.Itoa(int(intBin1%2)) + res
		intBin1 = intBin1 / 2
	}
	return res
}
