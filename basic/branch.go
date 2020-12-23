package main

import "fmt"

func main() {
	// 분기문, 조건식이 참이면 분기 안의 내용을 실행한다. 조건식은 항상 boolean 식으로 표현되어야 한다. (0 또는 1 같은거 못 넣음)
	if true {
		fmt.Println("true") // true
	} else if true { // else if (if 두 개 안됨)
		fmt.Println("else if true") // true
	} else {
		fmt.Println("else")
	}

	// optional statement, 분기문 이전에 간단한 문장을 실행할 수 있다.
	// if문에 국한되지 않고 조건을 가지는 문법(switch, for 등)에 사용할 수 있다.
	five := 5
	ten := 10

	if five *= 2 ; five == ten {
		fmt.Println("optional statement")
	}

	// switch, 다수의 조건식을 확인하는 경우
	x := 0
	switch { // 조건이 없을 수 있다. 기본적으로 true이다. if문 단순화 할 때 자주 쓴다. 조건이 표현식일수도 있다. (optional statement)
	case x == 0:
		fmt.Println(x) // 0
		x++
		fallthrough // x == 1이지만 fallthrough가 없으면(기본적으로) 다음 case로 이동하지 않고 종료한다.
	case x == 1:
		fmt.Println(x) // 1
	default:
		fmt.Println("default") // fallthrough로 끝까지 내려오거나 만족하는 case가 없으면 실행된다.
	}
}
