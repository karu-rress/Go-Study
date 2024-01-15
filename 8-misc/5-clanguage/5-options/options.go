package main

/*
// LDFLAGS: 링크 옵션, CFLAGS: 컴파일 옵션
#cgo LDFLAGS: -lz

#include <stdio.h>
#include <string.h>
#include <zlib.h>

void CExample() {
    char *text = "Hello, world!";
    unsigned long textSize = strlen(text);
    char buffer[1024] = "";
    unsigned long bufferSize = 1024;

    // zlib 압축
    int err = compress(buffer, &bufferSize, text, textSize);
    printf("Compressed size: %lu\n", bufferSize);
    printf("zlib Version: %s\n", zlibVersion());
}
*/
import "C"

func main() {
	C.CExample()
}