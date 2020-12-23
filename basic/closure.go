package main

import "fmt"

// 클로저는 함수 바깥의 변수를 참조하는 함수 값을 말한다.
// 바깥 함수는 내부에서 참조하는 변수값을 계속 유지한다.
func nextValue(startValue int) func() int {
	return func() int {
		startValue++ // 바깥 함수의 변수를 참조
		return startValue
	}
}

func main() {
	next := nextValue(3)

	fmt.Println(next()) // 4
	fmt.Println(next()) // 5
	fmt.Println(next()) // 6

	anotherNext := nextValue(10)

	fmt.Println(anotherNext()) // 11
	fmt.Println(anotherNext()) // 12
	fmt.Println(anotherNext()) // 13
}
