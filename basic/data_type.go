package main

import (
	"fmt"
	"strconv"
)

func main() {
	// bool, 참, 거짓을 가지는 불린 타입
	// string, 밑에서 다룸
	// integer, int, int8, int16, int32, int64
	//			uint, uint8, uint16, uint32, uint64, uintptr (unsigned)
	// float, float32 float64 complex64 complex128
	// string, 한번 생성하면 수정될 수 없는 Immutable한 값
	// `` BackQuote, 문자열 그대로의 값을 가진다, escape 무시한다.
	backQuote := `hello, \n
world`
	fmt.Println(backQuote) // hello, \n
						   // world

	// "" 쌍따옴표, 복수 라인에 걸쳐 쓸 수 없으며, escape 문자를 특별하게 해석한다.
	doubleQuotes := "hello, \nworld"
	fmt.Println(doubleQuotes) // hello,
							  // world

	// 데이터 타입 변환, 숫자간 변환
	var i int = 100
	var u uint = uint(i) + 1
	var f float32 = float32(u) + .123
	fmt.Println(u, f)

	// 데이터 타입 변환, "100" -> 100 (strconv 패키지 Atoi 함수)
	var integer int
	var err error

	integer, err = strconv.Atoi("100")
	fmt.Println(integer, err) // 100 <nil>

	// 변환할 수 없는 문자열의 경우, 변수에 0과 err에 오류가 담겨서 반환된다.
	integer, err = strconv.Atoi("no integer")
	fmt.Println(integer, err) // 0 strconv.Atoi: parsing "no integer": invalid syntax
}
