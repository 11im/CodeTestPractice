import (
	"strings"
)

var word []string = []string{"aya", "ye", "woo", "ma"}

func solution(babbling []string) int {

	res := 0
	for _, babb := range babbling {
		if canSay(babb) {
			res++
		}
	}

	return res
}

func canSay(babb string) bool {
	csIdx := 0
	for _, cs := range word {
		if strings.Contains(babb, cs) {
			csIdx += len(cs)
		}
	}

	return csIdx == len(babb)
}