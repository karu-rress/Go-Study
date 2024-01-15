package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string `tag1:"이름" tag2:"Name"`
	age  int    `tag1:"나이" tag2:"Age"`
}

func main() {
	var num int = 1
	fmt.Println(reflect.TypeOf(num)) // 자료형 이름 출력

	var f float64 = 3.14
	t := reflect.TypeOf(f)
	v := reflect.ValueOf(f)

	fmt.Println(t.Name(), t.Size())          // 자료형 이름과 크기
	fmt.Println(t.Kind() == reflect.Float64) // Float64인가?
	fmt.Println(v.Type())                    // 값이 담긴 변수의 자료형 이름
	fmt.Println(v.Float())                   // 값을 실수형으로 출력

	p := Person{}
	name, ok := reflect.TypeOf(p).FieldByName("name")
	fmt.Println(ok, name.Tag.Get("tag1"))

	age, ok := reflect.TypeOf(p).FieldByName("age")
	fmt.Println(ok, age.Tag.Get("tag2"))

	var a *int = new(int)
	*a = 1
	fmt.Println(reflect.ValueOf(a).Elem().Int()) // 포인터, 인터페이스는 Elem()을 써줘야 함
}
