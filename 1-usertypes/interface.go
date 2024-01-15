package main

import (
	"fmt"
	"strconv"
)

type MyInt int

func (i MyInt) Print() {
	fmt.Println(i)
}

type Rect struct {
	width, height int
}

func (r Rect) Print() {
	fmt.Println(r.width, r.height)
}

type Printer interface {
	Print()
}

func testInterface() {
	var i MyInt = 5
	r := Rect{10, 20}
	var p Printer

	p = i
	p.Print()

	p = r
	p.Print()
}

func implementsPrint() bool {
	var rect Rect
	_, ok := interface{}(rect).(Printer)

	return ok
}

type Any interface{}

// 빈 인터페이스는 모든 타입을 받을 수 있음
func getAny(a Any) {
	fmt.Println(a)
}

func formatString(arg interface{}) string {
	switch arg.(type) {
	case int:
		i := arg.(int) // int로 값 가져오기
		return strconv.Itoa(i)
	case string:
		s := arg.(string)
		return s
	case *Rect:
		r := arg.(*Rect)
		return strconv.Itoa(r.height * r.width)
	default:
		return "ERROR"
	}
}