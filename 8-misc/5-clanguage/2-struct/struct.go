package main

/*
#include <stdlib.h>

typedef struct {
	char *name;
	int age;
} PERSON;

PERSON* create(char *name, int age) {
	PERSON *p = (PERSON *)malloc(sizeof(PERSON));

	p->name = name;
	p->age = age;

	return p;
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	var p *C.PERSON

	name := C.CString("Maria")
	age := C.int(20)

	p = C.create(name, age)

	fmt.Println(C.GoString(p.name))
	fmt.Println(p.age)

	C.free(unsafe.Pointer(name)) // CString은 꼭 해제
	C.free(unsafe.Pointer(p))    // malloc도 당연히 해제
}
