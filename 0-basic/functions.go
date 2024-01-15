package main

import (
	"fmt"
	"os"
)

func sum(n ...int) (total int) {
	total = 0
	for _, value := range n {
		total += value
	}

	return
}

func calc() func(x int) int {
	a, b := 3, 5
	return func(x int) int {
		return a*x + b
	}
}

func ReadHello() {
	file, err := os.Open("test.txt")
	defer file.Close() // 지연 호출 (finally)

	if err != nil {
		return
	}

	buf := make([]byte, 100)
	if _, err = file.Read(buf); err != nil {
		return
	}

	fmt.Println(string(buf))
}

func PanicRecover() {
	defer func() {
		s := recover()
		fmt.Println(s) // panic의 메시지 받아옴
	}() // 바로 호출 (지연되어 실행)

	panic("Error!!!")
}