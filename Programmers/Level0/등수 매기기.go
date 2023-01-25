import "sort"

func solution(score [][]int) []int {
	res := make([]int, len(score))
	tmp := make([]int, len(score))
	tmpB := make([]int, len(score))

	for idx, arr := range score {
		tmp[idx] = arr[0] + arr[1]
	}

	copy(tmpB, tmp)

	sort.Sort(sort.Reverse(sort.IntSlice(tmp)))
	for sortedIdx, sortedVal := range tmp {
		for idx, val := range tmpB {
			if sortedVal == val && res[idx] == 0 {
				res[idx] = sortedIdx + 1
			}
		}
	}

	return res
}
