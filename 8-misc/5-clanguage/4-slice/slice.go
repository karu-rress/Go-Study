package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define ARRAY_LENGTH 5

// 슬라이스를 void *로 받아옴
void SliceToArray(void *p) {
	char *a = (char *)p;

	printf("%c\n", a[0]);
	printf("%s\n", a);
}

int *GetArray() {
	int *a = (int *)malloc(sizeof(int) * ARRAY_LENGTH);
	memset(a, 0, sizeof(int) * ARRAY_LENGTH);

	a[0] = 21; a[1] = -15; a[2] = 68; a[3] = 72; a[4] = -33;

	return a;
}

int GetLength() {
	return ARRAY_LENGTH;
}
*/
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	b := []byte("Hello, world!")

	C.SliceToArray(unsafe.Pointer(&b[0]))

	var a *C.int = C.GetArray()
	length := int(C.GetLength())

	hdr := reflect.SliceHeader{ // 배열을 감싸는 슬라이스 헤더
		Data: uintptr(unsafe.Pointer(a)),
		Len:  length,
		Cap:  length,
	}

	// 슬라이스 헤더 hdr은 []C.int, hdr의 레퍼런스는 그거의 포인터
	// 이해 안 되면 오른쪽에서 왼쪽으로 읽기 (형변환 후 역참조)
	s := *(*[]C.int)(unsafe.Pointer(&hdr))
	fmt.Println(s)
	fmt.Println(s[2:5])

	C.free(unsafe.Pointer(a))
}
