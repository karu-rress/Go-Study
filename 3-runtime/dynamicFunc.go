package main

import (
	"fmt"
	"reflect"
)

func h(args []reflect.Value) []reflect.Value {
	fmt.Println("Hello, world!")
	return nil
}

func DyFunc() {
	var hello func()
	fn := reflect.ValueOf(&hello).Elem() // hello의 주소 전달, elem으로 값 정보 가져오기
	v := reflect.MakeFunc(fn.Type(), h)  // h의 함수 정보 생성

	fn.Set(v) // fn(hello의 값 정보)에 v를 설정, 함수 연결
	hello()
}

///////

func Sum(args []reflect.Value) []reflect.Value {
	a, b := args[0], args[1]
	if a.Kind() != b.Kind() {
		fmt.Println("타입이 다릅니다.")
		return nil
	}

	switch a.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return []reflect.Value{reflect.ValueOf(a.Int() + b.Int())}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return []reflect.Value{reflect.ValueOf(a.Uint() + b.Uint())}
	case reflect.Float32, reflect.Float64:
		return []reflect.Value{reflect.ValueOf(a.Float() + b.Float())}
	case reflect.String:
		return []reflect.Value{reflect.ValueOf(a.String() + b.String())}
	default:
		return []reflect.Value{}
	}
}

func makeSum(fptr interface{}) {
	fn := reflect.ValueOf(fptr).Elem()    // 인터페이스에서 값 정보 가져오기
	v := reflect.MakeFunc(fn.Type(), Sum) // Sum 함수 정보 생성
	fn.Set(v)
}

func Dynamic() {
	var intSum func(int, int) int64
	var uintSum func(uint, uint) uint64
	var floatSum func(float32, float32) float64
	var stringSum func(string, string) string

	makeSum(&intSum)
	makeSum(&uintSum)
	makeSum(&floatSum)
	makeSum(&stringSum)

	fmt.Println(intSum(-1, 6))
	fmt.Println(uintSum(2, 4))
	fmt.Println(floatSum(3.1, 2.3))
	fmt.Println(stringSum("Hi, ", "Go"))
}
