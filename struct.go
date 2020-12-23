package main

import (
	"fmt"
	"reflect"
)

// 구조체 정의, 구조체는 struct 필드만을 가지며, 메서드는 분리하여 정의한다.
type person struct {
	name string
	age int
}

func newPerson() (pAddress *person) {
	pAddress = new(person)
	pAddress.name = "이름을 입력해주세요"
	pAddress.age = 100

	return
}

func main() {
	// 구조체 객체 생성
	p := person{"정우영", 18} // 값 초기화
	fmt.Println(reflect.TypeOf(p)) // main.person
	fmt.Println(p) // {정우영 18}

	p.name = "정우영"
	p.age = 18

	fmt.Println(p) // {정우영 18}

	// person type을 가진 p1 변수 선 생성 후 할당
	var p1 person
	p1 = person{name: "홍길동", age: 20}
	fmt.Println(p1) // {홍길동 20}

	// new() 함수로 구조체 생성
	p2 := new(person) // 모든 필드를 zero value로 초기화 후 person 객체의 포인터를 리턴한다.
	fmt.Println(p2) // &{"" 0}
	fmt.Println(*p2) // {"" 0}

	// 구조체 포인터에도 .(dot)으로 필드에 접근할 수 있다.
	p2.name = "필드 접근 할 수 있다고"
	fmt.Println(p2.name) // 필드 접근 할 수 있다고

	// 생성자 함수, 구조체를 초기화해주는 함수
	p3 := newPerson()
	fmt.Println(*p3) // {이름을 입력해주세요 100}
}
