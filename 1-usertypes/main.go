package main

import (
	"fmt"
)

type Rectangle struct {
	width  int
	height int
}

func NewRectangle(width, height int) *Rectangle {
	return &Rectangle{width, height}
}

// 리시버 변수는 *로 받거나 일반 구조체로 받음
// 포인터로 받으면 this, 그냥 받으면 const this 같은 느낌
// 안 쓰면 _로 받기
func (rect *Rectangle) area() int {
	return rect.width * rect.height
}

func testRectangle() {
	var rect1 *Rectangle = new(Rectangle)
	rect2 := Rectangle{45, 62}
	rect3 := NewRectangle(1, 2)
	_ = rect1

	fmt.Println(rect3)
	fmt.Println(rect2.area())
}

type Person struct {
	name string
	age  int
}

func (p Person) greeting() {
	fmt.Println("HI!")
}

type Student struct {
	Person // is ~ a: student is person!
	school string
	grade int	
} // 오버라이딩도 가능하긴 함

func main() {
	testRectangle()

	var s Student
	s.greeting()
}

