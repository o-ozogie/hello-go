package main

import (
	"fmt"
	"reflect"
)

// 수신 전용 채널을 리턴하는 함수, 채널에 값을 보내놓고 호출하는 함수에서 리턴된 채널을 수신
func receiveOnly() <- chan int {
	c := make(chan int)

	go func() {
		defer close(c) // 채널 닫기, 해당 채널로의 송신을 막는다. 수신은 가능하다.
		c <- 456
	}()

	return c
}

// 송신 전용 채널을 리턴하는 함수, 채널에서 값을 송신하고 호출하는 함수에서 리턴된 채널에 송신
func sendOnly() chan <- int {
	c := make(chan int)

	go func() {
		defer close(c)
		fmt.Println(<- c)
	}()

	return c
}

func unbufferedSend() {
	unbufferedChannel := make(chan int)
	unbufferedChannel <- 123 // deadlock, 송신자가 없음
}

func main() {
	// 정수형 채널을 만든다.
	c := make(chan int)

	fmt.Println(reflect.TypeOf(c)) // chan int

	go func() {
		c <- 123 // c 채널에 123을 보낸다.
	}()

	i := <-c // 채널로부터 123을 받는다. 채널이 데이터를 받을 때 까지 대기하므로 wait 코드가 필요없다.
	fmt.Println(i)


	bufferedChannel := make(chan int, 3)
	bufferedChannel <- 123 // 정상 실행
	fmt.Println(<- bufferedChannel)

	// 단방향 채널, 기본값은 양방향 채널이며 단방향 채널 또한 생성할 수 있다.
	fmt.Println(<- receiveOnly()) // 456
	sendOnly() <- 123 // 123

	// range 구문 다음의 채널에서 값을 수신한다. 채널이 닫히면 수신한 값을 반환한다.
	for channel := range receiveOnly() {
		fmt.Println(channel)
	} // 456

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		c1 <- "one"
	}()

	go func() {
		c2 <- "two"
	}()

	var recv1, recv2 string

	for recv1 == "" || recv2 == "" { // 모든 채널에서 데이터를 받기 위한 반복문
		select { // 복수의 채널을 기다리다가 데이터를 받은 채널을 실행하는 기능을 제공한다. switch문처럼 실행 후 select문을 종료한다.
		case recv1 = <- c1:
			fmt.Println(recv1)
		case recv2 = <- c2:
			fmt.Println(recv2)
		default: // hit된 채널이 없을 때 실행한다.
			fmt.Println("waiting")
		}
		fmt.Println("end select")
	}
}
