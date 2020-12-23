package main

import (
	"fmt"
	"reflect"
)

type myError struct {
	code string
	message string
}

func (e *myError) Error() string {
	// 표준 error 인터페이스 구현
	return e.code + ", " + e.message
}

func NotFound() error {
	// 에러 구조체의 주소를 반환
	return &myError{code: "404", message: "Not Found"}
}

func IntervalServerError() error {
	return &myError{code: "500", message: "Interval Server Error"}
}

func main() {
	notFound := NotFound()
	fmt.Println(notFound, reflect.TypeOf(notFound))

	intervalServerError := IntervalServerError()
	fmt.Println(intervalServerError, reflect.TypeOf(intervalServerError))
}
