package main

import "fmt"

func main() {
	helloWorld := "hello, world"
	function(helloWorld) // hello, world

	callByValue(helloWorld)
	fmt.Println(helloWorld) // hello, world

	callByReference(&helloWorld)
	fmt.Println(helloWorld) // hello, world!!!

	variableParameterFunction(1) // 1
	variableParameterFunction(1, 2, 3, 4) // 1 2 3 4

	s := sum(2, 3)
	fmt.Println(s) // 5

	max, min := maxMin(1, 2, 3, 4, 5)
	fmt.Println(max, min) // 5 1

	max, min = betterMaxMin(6, 7, 8, 9, 10)
	fmt.Println(max, min) // 10 6
}

// string 타입의 msg라는 변수를 파라미터로 받는 함수, 리턴값은 없다.
func function(msg string) {
	fmt.Println(msg)
}

// 변수의 값을 복사하여 연산했으므로 값을 호출한 함수에서는 값의 변화가 없다.
func callByValue(str string) {
	str += "!!!"
}

// 변수의 주소를 파라미터에 받아 주소의 값을 변경했으므로 호출한 함수에서 값의 변화가 있다.
func callByReference(strAddress *string) {
	*strAddress += "!!!"
}

// 함수 인자를 가변적으로 받기 위해서  ...를 사용한다.
func variableParameterFunction(args ...int) {
	for _, e := range args {
		fmt.Println(e)
	}
}

// 리턴값이 int형인 함수
func sum(a int, b int) int {
	return a + b
}

// 리턴값이 복수 개인 함수
func maxMin(elements ...int) (int, int) {
	max := elements[0]
	min := elements[0]

	for _, element := range elements {
		if element > max {
			max = element
		} else if element < min {
			min = element
		}
	}

	return max, min
}

// named return parameter, 함수 선언에서 리턴값에 해당하는 변수를 할당하여 가독성을 높일 수 있다.
func betterMaxMin(elements ...int) (max int, min int) {
	max = elements[0]
	min = elements[0]

	for _, element := range elements {
		if element > max {
			max = element
		} else if element < min {
			min = element
		}
	}

	return
}
