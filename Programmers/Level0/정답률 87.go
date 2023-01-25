// 세균 증식
func solution(n int, t int) int {
    res := n
    for i := 0; i < t ; i++ {
        res = res * 2
    }
    return res
}

// 가위 바위 보
func solution(rsp string) string {
    rock := "0"
    scissor := "2"
    paper := "5"
    res := ""
    for i := 0 ; i < len(rsp) ; i++{
        val := string(rsp[i])
        if val == rock{
            res += paper
        } else if val == scissor {
            res += rock
        } else {
            res += scissor
        }
    }
    return res
}

// 암호 해독
func solution(cipher string, code int) string {
    res := ""
    for i := code-1 ; i < len(cipher) ; i += code {
        res += string(cipher[i])
    }
    return res
}

// 대문자와 소문자
func solution(my_string string) string {
    res := ""
    low := []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
    upper := []string{"A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z"}
    for i := 0 ; i < len(my_string) ; i++ {
        for idx,lowchar := range(low){
            if string(my_string[i]) == lowchar{
                res += upper[idx]
            }
        }
        for idx,upperchar := range(upper){
            if string(my_string[i]) == upperchar{
                res += low[idx]
            }
        }
    }
    return res
}

// n의 배수 고르기
func solution(n int, numlist []int) []int {
    res := make([]int, 0)
    for _,num := range(numlist){
        if num%n == 0{
            res = append(res,num)
        }
    }
    return res
}

// 직각 삼각형 출력하기
import "fmt"
func solution() {
    var n int
    fmt.Scan(&n)
    
    for i := 1 ; i <= n ; i++{
        line := ""
        for j := 0 ; j < i ; j++{
            line +="*"
        }
        fmt.Println(line)
    }
}

// 주사위의 개수
func solution(box []int, n int) int {
    res := 1
    for _,line := range(box){
        res = res * (line/n)
    }
    return res
}

// 문자열 정렬하기(1)
import (
    "sort"
    "strconv"
)
func solution(my_string string) []int {
	res := []int{}
	for i := 0; i < len(my_string); i++ {
		num, err := strconv.Atoi(string(my_string[i]))
		if err != nil {
			continue
		} else{
			res = append(res, num)
		}
	}
	sort.Sort(sort.IntSlice(res))
	return res
}
