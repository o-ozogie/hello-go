package main

import (
	"fmt"
	"reflect"
)

func main() {
	// int형 변수 a 선언, 초깃값 없으면 Zero Value
	var a int
	fmt.Println(a) // 0
	// float32형 변수 b 선언, 초깃값 11.23
	var b float32 = 11.23
	fmt.Println(b) // 11.23
	// 동일한 타입 변수들 초깃값 설정
	var i, j, k = 1, 2, 3
	fmt.Println(i, j, k) // 1 2 3

	// 상수 PI 선언, 무조건 초기화 해줘야 하고 타입 안쓰면 타입 추론해서 넣음 변수도 마찬가지
	const PI = 3.14
	fmt.Println(reflect.TypeOf(PI)) // float64
	// 상수 여러 개 묶어서 지정
	const (
		NAME = "정우영"
		AGE = 18
	)
	fmt.Println(NAME, AGE) // 정우영 18
	// iota, 상수에서 0부터 값을 순차적으로 부여하기 위한 식별자
	const (
		A = iota
		B
		C
	)
	fmt.Println(A, B, C) // 0 1 2

	// short assignment statement, :=
	nationalRule := "hello, world"
	fmt.Println(nationalRule) // hello, world
}



