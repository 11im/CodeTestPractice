// 자릿수 더하기
func solution(n int) int {
    res := 0    
    for n > 0{
        res += n%10
        n = n/10
    }
    return res
}

// 문자열안에 문자열
import "strings"
func solution(str1 string, str2 string) int {
    if strings.Contains(str1, str2){
        return 1
    } else {
        return 2
    }
}

// 제곱수 판별하기
func solution(n int) int {
    res := 2
    i := 1
    for {
        if i*i == n {
            res = 1
            break
        } else if i*i > n {
            break
        } else{
            i++
        }
    }
    return res
}

// 숨어있는 숫자의 덧셈(1)
func solution(my_string string) int {
    nums := []string{"1","2","3","4","5","6","7","8","9"}
    res := 0
    for i := 0 ; i < len(my_string) ; i++ {
        for j := 0; j < len(nums); j++{
            if string(my_string[i]) == nums[j] {
                res += j+1
            }
        }
    }   
    return res
}

// 개미 군단
func solution(hp int) int {
    att := []int{5,3,1}   
    res := 0
    for _, ant := range att{
        res += hp / ant
        hp = hp % ant
    }
    return res
}

// 모음 제거
func solution(my_string string) string {
    vowels := []string{"a","e","i","o","u"}
    res := ""
    var flag bool
    for i := 0; i < len(my_string); i++{
        flag = false
        val := string(my_string[i])
        for _, vw := range(vowels){
            if val == vw{
                flag = true
                break
            }
        }
        if flag == false{
            res += val
        }
    }   
    return res
}