package main

import (
	"fmt"
)

func main() {
	// 배열 사용하기
	a := [...]int{32, 29, 87, 15, 23}
	for _, value := range a { // _: index
		fmt.Println(value)
	}

	// 슬라이스(가변 배열) 사용하기
	var b []int = make([]int, 5, 10) // 타입, 크기, (용량)
	b[3] = 3
	b = append(b, 3, 4, 5)
	fmt.Printf("len: %d, cap: %d\n\n", len(b), cap(b))

	var p *int = new(int)
	*p = 30
	fmt.Println(p)
}

func testfunc() {
	/*
	// 슬라이스를 여러 개의 값으로 분해 -> 가변인자로 넘김
	fmt.Println(sum(b...))

	f := calc()       // calc() 함수의 결괏값(클로져)를 f에 저장
	fmt.Println(f(1)) // 지역변수가 소멸되지 않음
	fmt.Println(f(2)) // 즉, 실행 흐름을 저장한 상태
	*/
}