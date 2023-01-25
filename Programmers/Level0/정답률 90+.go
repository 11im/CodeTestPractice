// 두 수의 곱
func solution(num1 int, num2 int) int {
	return num1 * num2
}

// 몫
func solution(num1 int, num2 int) int {
	return num1 / num2
}

// 나머지
func solution(num1 int, num2 int) int {
	return num1 % num2
}

// 두 수의 함
func solution(num1 int, num2 int) int {
	return num1 + num2
}

// 두 수의 차
func solution(num1 int, num2 int) int {
	return num1 - num2
}

// 숫자 비교하기
func solution(num1 int, num2 int) int {
	if num1 == num2 {
		return 1
	} else {
		return -1
	}
}

// 나이 출력
func solution(age int) int {
	return 2023 - age
}

// 두 수의 나눗셈
func solution(num1 int, num2 int) int {
	res := float64(num1) / float64(num2) * 1000
	return int(res)
}

// 각도기
func solution(angle int) int {
	switch {
	case angle > 0 && angle < 90:
		return 1
	case angle == 90:
		return 2
	case angle < 180 && angle > 90:
		return 3
	default:
		return 4
	}
}
