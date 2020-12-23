package main

import "fmt"

type rectangle struct {
	width, height int
}

// receiver 변수로 rectangle 객체를 받는다. call by value 이므로 객체가 변경되지 않는다.
func (r rectangle) area() int {
	r.width++
	r.height++
	return r.width * r.height
}

// receiver 변수로 rectangle 객체의 주소를 받는다. call by reference 이므로 객체가 변경된다.
func (r *rectangle) referenceArea() int {
	r.width++
	r.height++
	return r.width * r.height
}

func main() {
	r := rectangle{width: 13, height: 15}
	fmt.Println(r.area()) // 224
	fmt.Println(r) // {13 15}
	fmt.Println(r.referenceArea()) // 224
	fmt.Println(r) // {14 16}
}
