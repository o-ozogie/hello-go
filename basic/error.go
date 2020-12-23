package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
)

func main() {
	f, err := os.Open("~/Desktop/invalid-file")
	fmt.Println(reflect.TypeOf(f), reflect.TypeOf(err))
	if err != nil {
		log.Fatal(err.Error()) // 에러 메시지 출력 후 os.Exit(1)
	}
	fmt.Println(f.Name()) // ~/Desktop/invalid-file
}
