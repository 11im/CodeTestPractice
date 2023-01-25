import "sort"

func solution(array []int) int {
	resMap := make(map[int]int)
	var resArr []int
	for _, num := range array {
		if cnt, ok := resMap[num]; ok {
			resMap[num] = cnt + 1
		} else {
			resMap[num] = 1
		}
	}

	for _, cnt := range resMap {
		resArr = append(resArr, cnt)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(resArr)))

	if len(resArr) > 1 && resArr[0] == resArr[1] {
		return -1
	}

	res := 0

	for idx, cnt := range resMap {
		if cnt == resArr[0] {
			res = idx
			break
		}
	}

	return res

}