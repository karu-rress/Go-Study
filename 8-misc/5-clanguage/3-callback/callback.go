package main

/*
// Go언어 함수는 extern으로
extern void goCallback(void *p, int n);

// Go에서 쓸 건 반드시 static inline
static inline void CExample(void *p) {
	goCallback(p, 100);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

//export goCallback
func goCallback(p unsafe.Pointer, n C.int) {
	f := *(*func(C.int))(p) // p를 함수 포인터 타입으로 변환, 역참조
	f(n)
}

func main() {
	f := func(n C.int) {
		fmt.Println(n)
	}

	C.CExample(unsafe.Pointer(&f))
}
