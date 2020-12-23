package main

import "fmt"

func main() {
	// 산술연산자 [+ - / * % ++ --]
	i := ((5 + 10) / 3) % 2
	fmt.Println(i) // 1
	i++
	fmt.Println(i) // 2
	i--
	fmt.Println(i) // 1

	// 관계연산자 [== != < >]
	// 논리연산자 [&& || !]

	// Bitwise, binary OR (b011 | b110 == b111)
	fmt.Println(3 | 6) // 7
	// Bitwise, binary AND (b011 & b110 == b010)
	fmt.Println(3 & 6) // 2
	// Bitwise, binary XOR (b011 ^ b110 == 101)
	fmt.Println(3 ^ 6) // 5
	// Bitwise, binary shift (b011 << 1 == b110)
	fmt.Println(3 << 1) // 6
	fmt.Println(3 >> 1) // 1

	// 할당연산자 [= += -= &= <<=]

	// 포인터 연산자 [& : 변수의 주소 값, * : 주소의 실제 값]
	helloWorld := "hello, world"
	var address *string = &helloWorld
	fmt.Println(address) // 0xc000108040
	fmt.Println(*address) // hello, world
}
