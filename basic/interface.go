package main

import (
	"fmt"
)

// 함수 두 개를 포함하는 인터페이스, 파라미터와 반환형이 모두 같아야 한다.
type shape interface {
	area() int
	perimeter() int
}

// 구조체에서 인터페이스를 구현하고자 한다면 인터페이스에 포함된 함수를 모두 구현하면 된다.
type circle struct {
	radius int
}

func (c circle) area() int {
	return c.radius * c.radius
}

func (c circle) perimeter() int {
	return c.radius * 2 * 3
}

type triangle struct {
	bottomSide int
	height int
}

func (t triangle) area() int {
	return t.height * t.bottomSide / 2
}

func (t triangle) perimeter() int {
	return t.bottomSide * 3
}

// shape 인터페이스를 구현한 모든 구조체가 s에 들어갈 수 있다. (다형성)
func printf(s shape) {
	fmt.Println(s)
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}

// empty interface, 모든 Type을 나타내기 위해 사용한다. (duck typing)
func println(content interface{}) {
	fmt.Println(content)
}

func main() {
	printf(circle{radius: 1}) // {1} 1 6
	printf(triangle{bottomSide: 1, height: 2}) // {1 2} 1 3

	println(circle{radius: 10}.perimeter()) // 60
	println(triangle{bottomSide: 10, height: 20}.perimeter()) // 30

	// interface type 변수 x와 타입 T에 대하여 x.(T)로 표현했을 때,
	// x는 nil이 아니며 x는 T 타입에 속한다는 점을 확인하는 것을 Type Assertion 이라 한다.
	// x가 nil이거나 x의 타입이 T가 아니라면 런타임 에러, 맞으면 T 타입의 x를 리턴한다.
	var x interface{} = 0
	x = "hello"
	fmt.Println(x) // hello

	y, existY := x.(string)
	fmt.Println(y, existY) // hello true

	z, existZ := x.(int)
	fmt.Println(z, existZ) // 0 false

}
