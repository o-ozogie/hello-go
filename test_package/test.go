package test_package

import "fmt"

// public 식별자 (식별자 이름이 대문자로 시작)
func Test() {
	fmt.Println("test_package package Test function public")
}

// non-public 식별자 (식별자 이름이 소문자로 시작)
func test() {
	fmt.Println("test_package package test function private")
}

// 패키지가 import될 때 호출되는 초기화 변수
func init() {
	fmt.Println("test_package initialize")
}
