func solution(numer1 int, denom1 int, numer2 int, denom2 int) []int {
	res := make([]int, 2)

	numer := numer1*denom2 + numer2*denom1
	denom := denom1 * denom2

	gcd := getGCD(numer, denom)

	res[0] = numer / gcd
	res[1] = denom / gcd

	return res
}

func getGCD(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for b > 0 {
		a, b = b, a%b
	}
	return a
}