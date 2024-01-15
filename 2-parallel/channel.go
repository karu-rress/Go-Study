package main

import (
	"fmt"
	"time"
)

// 보내기 전용 채널
func producer(c chan<- int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
	c <- 100
}

// 받기 전용 채널
func consumer(c <-chan int) {
	for i := range c {
		fmt.Println(i)
	}

	fmt.Println(<-c)
}

func Chan() {
	c := make(chan int)

	go producer(c)
	go consumer(c)

	fmt.Scanln()
}

func Num(a, b int) <-chan int {
	out := make(chan int)
	go func() {
		out <- a
		out <- b
		close(out)
	}()
	return out
}

func Sum(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		r := 0
		for i := range c {
			r += i
		}
		out <- r
	}()
	return out
}

func ChanFunc() {
	c := Num(1, 2) // 1과 2가 들어있는 채널이 리턴됨
	out := Sum(c)  // 채널 자체를 매개변수로 넘겨서 모두 더함
	fmt.Println(<-out)
}

func Select() {
	c1 := make(chan int)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- 10
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for {
			c2 <- "Hello, world!"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			select {
			// case c1 <- 10: // 매번 채널에 값을 보냄
			case i := <-c1: // 채널 c1에 값이 들어왔다면?
				fmt.Println("c1: ", i)
			case s := <-c2:
				fmt.Println("c2: ", s)
			}
		}
	}()

	time.Sleep(5 * time.Second)
}
