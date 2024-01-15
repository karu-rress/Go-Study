package main

import (
	"fmt"
	"log"
	"time"
)

type HelloOneError struct {
	time  time.Time
	value int
}

func (e HelloOneError) Error() string {
	return fmt.Sprintf("%v: %d는 1이 아닙니다.", e.time, e.value)
}

func helloOne(n int) (string, error) {
	if n == 1 {
		s := fmt.Sprintln("Hello", n)
		return s, nil
	}

	return "", HelloOneError{time.Now(), n}
}

func main() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()

	s, err := helloOne(1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)

	s, err = helloOne(2)
	if err != nil {
		// Print: 에러 문자열만 출력
		// log.Print(err)

		// Fatal: 에러 문자열 출력 후 프로그램 정상 종료 (code 1)
		// log.Fatal(err)

		// Panic: 런타임 에러 발생시킴
		// panic(err)
		log.Panic(err) // 그렇지만 recover로 살렸죠?

		// err.(HelloOneError).value 이렇게 값 꺼내오기 가능
	}

	fmt.Println(s)
}
