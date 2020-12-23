package main

import (
	"fmt"
)

func main() {
	// 익명함수, func 뒤에 함수명을 생략하여 선언할 수 있고, 변수에 할당해서 사용할 수 있다.
	sum := func(elements ...int) (result int) {
			for _, element := range elements {
				result += element
			}
			return
		}

	result := sum(1, 2, 3, 4, 5)
	fmt.Println(result) // 15

	fmt.Println(calculate(sum, 1, 2)) // 3
	fmt.Println(calculate(func(elements ...int) int { return elements[0] - elements[1] }, 2, 1)) // 1
	fmt.Println(compactCalculate(sum, 1, 2)) // 3
}

// 함수는 일급 함수로서 기본 타입과 동일하게 취급되며,
// 다른 함수의 파라미터로 전달하거나 함수의 리턴값으로 사용 될 수도 있다.
func calculate(f func(elements ...int) int, a int, b int) int {
	result := f(a, b)
	return result
}

// type 문은 구조체, 인터페이스 등 custom type을 정의하기 위해 사용된다.
// type 문으로 함수의 원형 또한 정의 할 수 있다.
// 이처럼 함수의 원형을 정의하고 함수를 타 메서드에 전달하고 리턴받는 기능을 delegate라 한다.
type typeCalculate func(elements ...int) int

// 함수의 원형을 정의하여 가독성이 나아졌다.
func compactCalculate(f typeCalculate, a int, b int) int {
	result := f(a, b)
	return result
}
