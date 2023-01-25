// 가장 큰 수 찾기
func solution(array []int) []int {
    res := []int{}
    maxIdx := 0
    maxVal := 0
    for idx, val := range(array){
        if val > maxVal {
            maxVal = val
            maxIdx = idx
        }
    }
    res = append(res, maxVal, maxIdx)
    return res
}

// 배열 회전시키기
func solution(numbers []int, direction string) []int {
    var res []int
    if direction == "left"{
        for i := 1; i < len(numbers); i++{
            res = append(res, numbers[i])
        }
        res = append(res, numbers[0])
    } else {
        res = append(res, numbers[len(numbers)-1])
        for i := 0; i < len(numbers) -1; i++{
            res = append(res, numbers[i])
        }
    }       
    return res
}

// 외계행성의 나이
func solution(age int) string {
    res := ""
    vals:="abcdefghijklmnopqrstuvwxyz"
    for i := age ; i > 0; i=i/10{
        res = string(vals[i%10]) + res
    }
    return res
}

// 피자 나눠 먹기(2)
func solution(n int) int {
    gcd := getGCD(6,n)
    res := 6 * n / gcd
    res = res / 6
    return res
}
func getGCD(num1, num2 int) int {   
    if num1 < num2 {
        num1, num2 = num2, num1
    }
    for num2 > 0{
        num1, num2 = num2, num1%num2
    }
    return num1
}

// 최댓값 만들기(2)
import "sort"
func solution(numbers []int) int {
    sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
    res1 := numbers[0] * numbers[1]
    res2 := numbers[len(numbers)-1] * numbers[len(numbers)-2]
    if res1 > res2{
        return res1
    } else{
        return res2
    }
}

// 인덱스 바꾸기
func solution(my_string string, num1 int, num2 int) string {
    res := ""
    if num1 > num2 {
        num1,num2 = num2,num1
    }
    for i := 0; i < len(my_string); i++{
        switch i {
            case num1 :
                res += string(my_string[num2])
            case num2 :
                res += string(my_string[num1])
            default:
                res += string(my_string[i])
        }
    }
    return res
}

// 약수 구하기
func solution(n int) []int {
    var res []int
    for i := 1 ; i <= n/2 ; i++{
        if n % i == 0 {
            res = append(res, i)
        }
    }
    res = append(res, n)
    return res
}

// 숫자 찾기
func solution(num int, k int) int {
    res := 0
    idx := 0
    flag := false
    for num > 0 {
        if k == num % 10{
            res = idx
            flag = true
        }
        idx++
        num = num/10
    }
    if flag != true{
        res = -1
    } else{
        res = idx - res
    }   
    return res
}