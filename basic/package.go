package main // main 패키지는 패키지 안의 main() 함수가 entry point가 된다.

// 패키지는 모듈화 단위이며, 패키지를 통해 코드를 세분화한다.
// 커스텀 패키지 "test_package" import
import (
	// alias를 사용하여 패키지 이름을 변경할 수 있다.
	testPackage "hello-go/basic/test_package" // test_package initialize
)

func main() {
	testPackage.Test() // // test_package package Test function public
	testPackage.Test() // test_package package Test function public
}
