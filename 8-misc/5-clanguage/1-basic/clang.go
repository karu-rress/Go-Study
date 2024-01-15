package main

/*
#include <stdlib.h>
#include <time.h>

int sum(int a, int b) {
	return a + b;
}

void hello() {
	printf("Hello, world!\n");
}
*/
import "C"
import "fmt"

func main() {
	C.srand(C.uint(C.time(nil)))
	r := C.rand()
	fmt.Println(r)

	a, b := 1, 2
	r = C.sum(C.int(a), C.int(b))
	fmt.Println(r)

	C.hello()
}
