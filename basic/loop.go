package main

import "fmt"

func main() {
	// for문, golang의 유일한 반복문이다. 문법은 C와 동일하다.
	sum := 0
	for i := 1 ; i <= 100 ; i++ {
		sum += i
	}
	fmt.Println(sum) // 5050

	// for문에 초기값과 증감식을 생략하면 while문처럼 동작한다.
	n := 1
	for n < 200 {
		n *= 2
	}
	fmt.Println(n) // 256

	// for문에 아무 것도 안 쓰면 무한루프가 된다.
	//for {
	//	fmt.Println("can't stop")
	//}

	// for range 문, 루비의 each, 자바의 for each와 동일한 맥락의 문법이다.
	names := []string{"a", "b", "c", "d", "e"}

	for index, name := range names {
		fmt.Println(index, name)
	} // 0 a 1 b 2 c 3 d 4 e

	// break, continue, goto
	a := 1
	for a < 15 {
		if a == 5 {
			a *= 2
			continue // 밑 명령들을 무시하고 다음 loop를 실행한다.
		}
		a++
		if a > 10 {
			break // 반복문을 탈출한다.
		}
	}
	if a == 11 {
		goto END // END의 레이블로 이동한다.
	}
	fmt.Println(a) // goto문이 실행되었기 때문에 생략된다.
END:
	fmt.Println("End") // End

	// break, continue와 레이블을 함께 사용하여 지정된 레이블로 이동할 수 있다.
	I:
	for i := 1 ; i <= 100 ; i++{
		for j := 1 ; j <= 100 ; j++ {
			if i + j == 100 {
				fmt.Println(i, j) // 1 99
				break I
			}
		}
	}
}
