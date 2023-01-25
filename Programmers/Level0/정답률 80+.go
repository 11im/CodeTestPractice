// 369게임
func solution(order int) int {
    res := 0
    for order > 0{
        if order%10 == 3 || order%10 == 6 || order % 10 == 9{
            res++
        }
        order = order/10
    }
    return res
}

// 문자열 정렬하기(2)
import 
("strings"
 "sort")
func solution(my_string string) string {
    var tmp []string
    res := ""
    my_string = strings.ToLower(my_string)
    for i:=0; i < len(my_string); i++{
        tmp = append(tmp,string(my_string[i]))
    }
    sort.Sort(sort.StringSlice(tmp))
    for _, val := range(tmp){
        res += val
    }
    return res
}

// 합성수 찾기
func solution(n int) int {
    res := 0
    for i := 1 ; i <= n ; i++{
        cnt := 0
        for j := 1 ; j <= i ; j++{
            if i % j == 0 {
                cnt++
            }
        }
        if cnt > 2 {
            res++
        }
    } 
    return res
}

// 모스 부호(1)
func solution(letter string) string {
    res :=""
    tmp := ""
    morse := map[string]string{ 
		".-":"a","-...":"b","-.-.":"c","-..":"d",".":"e","..-.":"f",
		"--.":"g","....":"h","..":"i",".---":"j","-.-":"k",".-..":"l",
		"--":"m","-.":"n","---":"o",".--.":"p","--.-":"q",".-.":"r",
		"...":"s","-":"t","..-":"u","...-":"v",".--":"w","-..-":"x",
		"-.--":"y","--..":"z"}
    for i := 0 ; i < len(letter) ; i++ {
        chr := string(letter[i])
        if chr != " "{
            tmp += chr
        }
        if chr == " "  || i == len(letter)-1 {
            res += morse[tmp]
            tmp = ""
        }
    }
    return res
}

// 중복된 문자 제거
func solution(my_string string) string {
    res := ""
    tmp := make(map[string]int)
    for i := 0 ; i < len(my_string) ; i++ {
        val := string(my_string[i])
        _, err := tmp[val]
        if err {
            continue   
        } else{
            tmp[val] = 1
            res += val
        }   
    }
    return res
}

// 팩토리얼
func solution(n int) int {
    res := 1
    var tmp float64 = 1
    for {
        if tmp * float64(res) < float64(n) {
            res++
            tmp = tmp * float64(res)
        } else {
            break
        }
    }
    return res
}

// A로 B 만들기
func solution(before string, after string) int {
    res := 0
    befMap := make(map[string]int)
    for i := 0 ; i < len(before) ; i++ {
        val := string(before[i])
        if _, ok := befMap[val]; ok{
            befMap[val]++
        } else {
            befMap[val] = 1
        }
    }
    for i := 0 ; i < len(after) ; i++ {   
        val := string(after[i])
        _, ok := befMap[val]
        if ok == false || befMap[val] <= 0 {
            res = 0
            break
        }
        if ok && befMap[val] > 0 {
            res = 1
            befMap[val]--
        }
    }
    return res
}

// 2차원으로 만들기
func solution(num_list []int, n int) [][]int {
	res := make([][]int, len(num_list)/n)
	j := 0
	for idx, val := range num_list {
		if idx%n == 0 && idx != 0 {
			j++
		}
		res[j] = append(res[j], val)
	}
	return res
}

// 가까운 수
func solution(array []int, n int) int {
    val := 100
    res := 0
    for _, num := range array {
        tmp := num - n
        if tmp < 0 {
            tmp = -tmp
        }
        if tmp < val {
            val = tmp
            res = num
        }
        if tmp == val && res > num{
                res = num
        }
    }
    return res
}

// K의 개수
func solution(i int, j int, k int) int {
    res := 0
    for ; i <= j; i++ {
        tmp := i
        for tmp > 0 {
            if tmp % 10 == k {
                res++
            }
            tmp = tmp/10
        }
    }
    return res
}

// 진료순서 정하기
import "sort"
func solution(emergency []int) []int {
	res := make([]int, len(emergency))
	tmp := make([]int, len(emergency))
	copy(tmp, emergency)
	sort.Sort(sort.Reverse(sort.IntSlice(emergency)))
	for idx1, val1 := range tmp {
		for idx2, val2 := range emergency {
			if val1 == val2 {
				res[idx1] = idx2 + 1
			}
		}
	}
	return res
}

// 한 번 만 등장한 문자
func solution(s string) string {
    res := ""
    tmp := make([]int, 26)
    for _, val := range(s) {
        tmp[val - 'a']++
    }
    for idx, cnt := range(tmp) {
        if cnt == 1{
            res += string(idx+'a')
        }
    }   
    return res
}