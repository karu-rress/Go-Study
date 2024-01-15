package main

import "fmt"

func Printer() {
	var num int = 10
	var f float32 = 3.2
	var c complex64 = 1.2 + 4.3i
	var b bool = true
	var a []int = []int{1, 2, 3}
	var m map[string]int = map[string]int{"Hello": 1}
	var p *int = new(int)

	fmt.Printf("int: %d\n", num)
	fmt.Printf("float: %f\n", f)
	fmt.Printf("complex: %f\n", c)
	fmt.Printf("bool: %t\n", b)
	fmt.Printf("slice: %v\n", a)
	fmt.Printf("map: %v\n", m)
	fmt.Printf("ptr: %p\n", p)
	// 구조체, 인터페이스도 v로 출력 (default 지정자)

	/*
		%v: 모든 값. %+v: 필드도 함께 표시
		%#v: 타입과 값을 함께 표시
		%T: 타입을 표시
		%%: % 문자
		%c: rune 문자 하나
		%d, %o, %x, %X: 정수
		%q: 문자열을 escape하여 출력 (\ 붙여서)
		%U: 유니코드 형식
		%b: 이진수 정수 또는 2의 제곱형 실수
		%e, %E: 부동소수점
		%f, %F: 고정소수점
		%g, %G: 짧은 것
		%p: 포인터, 채널
	*/
}

func Scanner() {
	var s1, s2 string
	// Scanln - 띄어쓰기로 구분
	// Scan - 띄어쓰기 및 개행으로 구분
	n, _ := fmt.Scanln(&s1, &s2)

	fmt.Println("입력 개수: ", n)
	fmt.Println(s1, s2)
}

func main() {
	Printer()
	Scanner()
}
