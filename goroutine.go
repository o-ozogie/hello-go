package main

import (
	"fmt"
	"runtime"
	"sync"
)

func say(s string) {
	fmt.Println(s)
}


func main() {
	// 함수를 동기적으로 호출
	say("sync")

	var wait sync.WaitGroup
	wait.Add(2)

	go func() { // goroutine, 논리적 스레딩을 통한 비동기적 함수 호출
		say("async1")
		defer wait.Done()
	}()

	go func() { // goroutine, 논리적 스레딩을 통한 비동기적 함수 호출
		say("async2")
		defer wait.Done()
	}()

	wait.Wait()

	// sync
	// sync2
	// sync1

	runtime.GOMAXPROCS(2) // 머신이 복수개의 CPU를 가진 경우, 다중 CPU에서 병렬처리하도록 한다.
}
