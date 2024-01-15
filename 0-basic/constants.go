package main

const (
	x, y      int = 30, 50
	age, name     = 20, "karu"
)

// enum
const (
	mon = iota
	tue
	wed
	_ // iota 건너뛰기
	thu
	fri = 1 << iota // 연산을 할 수도 있음
)
