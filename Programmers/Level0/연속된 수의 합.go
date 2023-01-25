func solution(num int, total int) []int {
	var res []int
	var left, right, mid int

	if num%2 == 1 {
		mid = total / num
		left = num / 2
		right = num / 2
	} else {
		mid = total / num
		left = num/2 - 1
		right = num / 2
	}

	for i := mid - left; i <= mid+right; i++ {
		res = append(res, i)
	}

	return res
}