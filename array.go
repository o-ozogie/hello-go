package main

import "fmt"

func main() {
	// 정수형 3개의 요소를 갖는 배열 a
	var a [3]int
	fmt.Println(a) // [0 0 0]

	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Println(a) // [1 2 3]

	var a1 = [3]int{1, 2, 3} // 배열 선언과 함께 초기화
	fmt.Println(a1) // [1 2 3]

	var a2 = [...]int{1, 2, 3} // 초기화 요소의 갯수만큼 배열 크기를 자동으로 정한다.
	fmt.Println(a2) // [1 2 3]

	// 다차원 배열
	var ma [3][4]int
	fmt.Println(ma) // [[0 0 0 0] [0 0 0 0] [0 0 0 0]]

	// 다차원 배열 초기화
	var ma2 = [3][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(ma2) // [[1 2 3 4] [5 6 7 8] [9 10 11 12]]
}
