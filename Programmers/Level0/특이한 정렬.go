import "sort"

func solution(numlist []int, n int) []int {

	abs := make(Arrs, len(numlist))
	res := make([]int, len(numlist))
	var tmp int

	for idx, val := range numlist {
		tmp = val - n
		if tmp < 0 {
			tmp = -tmp
		}
		abs[idx].abs = tmp
		abs[idx].origin = val
	}

	sort.Sort(abs)
	for sortedIdx, sortedArr := range abs {
		for _, val := range numlist {
			if val == sortedArr.origin {
				res[sortedIdx] = val
			}
		}
	}

	return res
}

type Arr struct {
	abs    int
	origin int
}

type Arrs []Arr

func (a Arrs) Len() int {
	return len(a)
}

func (a Arrs) Less(i, j int) bool {
	switch {
	case a[i].abs == a[j].abs:
		if a[i].origin < a[j].origin {
			return false
		} else {
			return true
		}
	case a[i].abs < a[j].abs:
		return true
	case a[i].abs > a[j].abs:
		return false
	}
	return false
}
func (a Arrs) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
