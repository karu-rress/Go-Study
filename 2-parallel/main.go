package main

import (
	"fmt"
	"runtime"
	"time"
)

func hello() {
	fmt.Println("Hello")
}

func main() {
	go hello() // 고루틴으로 실행
	fmt.Println(runtime.GOMAXPROCS(0))

	fmt.Scanln()

	for i := 0; i < 5; i++ {
		// go로 하면 반복문이 끝난 다음에 실행됨.
		// 클로저 내부에서 i에 접근하면 5만 뜸. (반복문 끝)
		go func(n int) {
			fmt.Println("HI", n)
		}(i)
	}

	syncChan()
	asyncChan()

}

func sum(a int, b int, c chan int) {
	c <- a + b
}

func syncChan() {
	done := make(chan bool)
	count := 3

	go func() {
		for i := 0; i < count; i++ {
			// 이미 채널에 값이 있으면 대기함
			done <- true
			fmt.Println("고루틴: ", i)
			time.Sleep(1 * time.Second)
		}
	}()

	for i := 0; i < count; i++ {
		// 고루틴이 채널에 값을 보낼 때까지 대기 후 가져옴
		<-done
		fmt.Println("호출자: ", i)
	}
}

func asyncChan() {
	done := make(chan bool, 2)
	count := 4

	go func() {
		for i := 0; i < count; i++ {
			// 채널이 가득 차면 대기함
			done <- true
			fmt.Println("고루틴: ", i)
		}
		close(done) // 채널 닫기
	}()

	// 채널을 닫을 때까지 반복. 닫지 않았으면 대기
	for i := range done {
		fmt.Println("호출자: ", i)
	}
}
