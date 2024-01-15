package main

/*
	compress/bzip2
	compress/zlib
	compress/flate
	compress/lzw
*/
import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func Compress() {
	file, err := os.OpenFile(
		"hello.txt.gz",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		os.FileMode(0644),
	)
	if err != nil {
		return
	}
	defer file.Close()

	w := gzip.NewWriter(file)
	defer w.Close()

	s := "Hello, world!"
	w.Write([]byte(s))
	w.Flush()
}

func Decompress() {
	file, err := os.Open("hello.txt.gz")
	if err != nil {
		return
	}
	defer file.Close()

	r, err := gzip.NewReader(file)
	if err != nil {
		return
	}
	defer r.Close()

	if b, err := io.ReadAll(r); err == nil {
		fmt.Println(string(b))
	}
}

func main() {
	Compress()
	Decompress()
}
