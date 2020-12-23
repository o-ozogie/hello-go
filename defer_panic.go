package main

import "fmt"

func panic_recover() {
	// recover, panic() 함수에 의한 패닉상태를 정상상태로 되돌리는 함수이다.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("error,", r)
		}
	}()

	// panic, 함수를 즉시 멈추고 함수의 defer 구문을 모두 실행한 뒤 즉시 리턴한다.
	// 상위 함수에도 똑같이 적용되며 계속 콜스택을 타고 올라가며, 마지막에는 프로그램이 에러와 함께 종료된다.
	panic("panic!")
}

func d() {
	// defer, defer을 호출하는 함수가 리턴하기 직전에 특정 구문을 실행한다.
	fmt.Println("before defer")
	defer fmt.Println("defer")
	fmt.Println("after defer")
}

func main() {
	d()
	// before defer
	// after defer
	// defer

	panic_recover()

	//recover에 의한 회복으로 다음 구문이 정상적으로 실행된다.
	fmt.Println("status normal") // status normal
}
