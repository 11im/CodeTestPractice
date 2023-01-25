// 짝수의 합
func solution(n int) int {
    res := 0
    i:= 0
    for i<= n {
        res+=i
        i+=2
    }
    return res
}

// 배열의 평균값
func solution(numbers []int) float64 {
    var res float64 = 0
    for _,num := range numbers {
        res += float64(num)
    }
    res = res/float64(len(numbers))
    return res
}

// 양꼬치
func solution(n int, k int) int {
    drink := k - (n/10)
    food := n * 12000
    drink = drink * 2000
    res := food + drink
    return res
}

// 배열 원소의 길이
func solution(strlist []string) []int {
    res := make([]int,len(strlist))
    for i,val := range(strlist){
        res[i] = len(val)
    }
    return res
}

// 피자 나눠먹기(1)
func solution(n int) int {   
    res := n / 7
    if n % 7 != 0 {
        res++
    }
    return res
}

// 점의 위치 구하기
func solution(dot []int) int {
    x := dot[0]
    y := dot[1]    
    switch {
        case x > 0 && y > 0 :
            return 1
        case x < 0 && y > 0 :
            return 2
        case x < 0 && y < 0 :
            return 3
    default :
            return 4
    }
}

// 아이스 아메리카노
func solution(money int) []int {
    COST := 5500
    res := []int{money/COST, money%COST}
    return res
}

// 짝수 홀수 개수
func solution(num_list []int) []int {
    res := make([]int,2)
    even := 0
    odd := 0
    for _,num := range(num_list){
        if num %2 == 0{
            even++
        } else{
            odd++
        }
    }
    res[0] = even
    res[1] = odd
    return res
}

// 문자열 뒤집기
func solution(my_string string) string {
    res := ""
    for i := len(my_string) -1; i >=0 ; i--{
        res += string(my_string[i])
    }
    return res
}

// 배열 뒤집기
func solution(num_list []int) []int {
    res := make([]int, len(num_list))
    for i,num := range num_list {
        res[len(num_list)-i-1] = num
    }
    return res
}

// 최댓값 만들기 (1)
import "sort"
func solution(numbers []int) int {
    sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
    res := numbers[0] * numbers[1]
    return res
}

// 삼각형의 완성 조건 (1)
func solution(sides []int) int {
    if sides[0]-sides[1]-sides[2] >= 0 || sides[1]-sides[0]-sides[2] >= 0 || sides[2]-sides[0]-sides[1] >= 0 {
        return 2
    }else{
        return 1
    }
}

// 머쓱이보다 큰 사람
func solution(array []int, height int) int {
    res := 0
    for _, num := range array{
        if num > height{
            res++
        }
    }
    return res
}

// 짝수는 싫어요
func solution(n int) []int {
    var res []int
    for i := 1; i <= n; i+=2{
        res = append(res,int(i))
    }
    return res
}

// 특정 문자 제거하기
func solution(my_string string, letter string) string {
    res := ""
    for i := 0 ; i < len(my_string); i++{
        if string(my_string[i]) == letter{
            continue
        }
        res += string(my_string[i])
    }   
    return res
}

// 피자 나눠먹기(3)
func solution(slice int, n int) int {
    res := n / slice
    if n % slice != 0{
        res++
    }
    return res
}

// 문자 반복 출력하기
func solution(my_string string, n int) string {
    res := ""
    for i := 0; i < len(my_string); i++{
        for j :=0; j < n; j++{
            res += string(my_string[i])
        }
    }      
    return res
}

// 배열 자르기
func solution(numbers []int, num1 int, num2 int) []int {
    leng := num2-num1+1
    res := make([]int, leng)
    for i := 0; i < leng; i++{
        res[i] = numbers[num1+i]
    }
    return res
}

// 중복된 숫자 개수
func solution(array []int, n int) int {    
    res := 0
    for _, num := range array {
        if num == n {
            res++
        }
    }       
    return res
}

// 배열 두 배 만들기
func solution(numbers []int) []int {
    for i,num := range(numbers){
        numbers[i] = num*2
    }
    return numbers
}

// 편지
func solution(message string) int {
    res := len(message) * 2    
    return res
}

// 중앙값 구하기
import "sort"
func solution(array []int) int {
    sort.Sort(sort.IntSlice(array))
    res := array[len(array)/2]
    return res
}

// 순서쌍의 개수
func solution(n int) int {
    max := n/2
    res := 1
    for i := 1 ; i <= max ; i++{
        if n % i == 0{
            res++
        }
    }   
    return res
}

// 옷가게 할인 받기
func solution(price int) int {
    var dc float64
    switch {
        case price >= 500000 :
            dc = 0.8
        case price >= 300000 :
            dc = 0.9
        case price >= 100000 :
            dc = 0.95
        default:
            dc = 1
    }
    res := int(float64(price) * dc)   
    return res
}

// 배열의 유사도
func solution(s1 []string, s2 []string) int {
    res := 0
    for _,var1 := range(s1){
        for _,var2 := range(s2){
            if var1 == var2{
                res++
            }
        }
    }   
    return res
}