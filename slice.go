package main

import "fmt"

func main() {
	// 고정된 크기를 지정하지 않고, 지정된 크기를 동적으로 변경할 수 있는, 부분 배열을 발췌할 수 있는 슬라이스
	var a []int // nil slice, 별도의 길이와 용량을 지정하지 않은 슬라이스 (== nil)
	a = []int{1, 2, 3}
	a[1] = 10
	fmt.Println(a, len(a), cap(a)) // [1 10 3] 3 3

	// make 함수를 사용하여 개발자가 슬라이스의 길이와 용량을 임의로 지정할 수 있다.
	// 길이 : 슬라이스의 최초 길이, 용량 : 슬라이스의 최대 길이
	ma := make([]int, 5, 6)
	fmt.Println(ma, len(ma), cap(ma)) // [0 0 0 0 0] 5 6
	ma = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(ma, len(ma), cap(ma)) // [1 2 3 4 5 6 7] 7 7

	// 부분 슬라이스 [시작 인덱스:끝 인덱스 + 1], 처음 인덱스는 포함, 끝 인덱스는 배제한다.
	fmt.Println(ma[2:5]) // [3 4 5]

	// a에 값을 추가한다. 원래 배열에 값이 추가되는 것이 아닌, 값이 추가된 배열이 반환된다.
	fmt.Println(append(a, 4)) // [1 10 3 4]
	fmt.Println(a) // [1 10 3]

	// 슬라이스의 길이가 슬라이스의 용량을 초과하는 경우, 슬라이스의 용량 크기를 두 배로 증가한다.
	var sliceA []int
	for i := 1 ; i <= 15 ; i++ {
		sliceA = append(sliceA, i)
		fmt.Println(len(sliceA), cap(sliceA))
	} // 1 1
	  // 2 2
	  // 3 4
	  // 4 4
	  // 5 8
	  // 6 8
	  // 7 8
	  // 8 8
	  // 9 16
	  // 10 16
	  // 11 16
	  // 12 16
	  // 13 16
	  // 14 16
	  // 15 16
	fmt.Println(sliceA) // [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15]

	// 슬라이스의 통합
	newSlice := append(a, sliceA...) // sliceA의 요소를 unpacking하여 slice a에 추가한다.
	fmt.Println(newSlice) // [1 10 3 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15]

	newSlice = make([]int, len(a), cap(a) * 2)
	copy(newSlice, a) // 슬라이스는 실제 배열의 포인터 정보만을 가지므로, 실제로 복사는 복사되는 배열 데이터를 복사받는 배열로 복제 하는 것
	fmt.Println(newSlice) // [1 10 3]
}
