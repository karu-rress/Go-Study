package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
)

func File() {
	// Save
	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	s := "Hello, world!"

	if _, err := file.Write([]byte(s)); err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
	file.Close()
	fmt.Println("저장 완료")

	// Open
	file, err = os.Open("hello.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		return
	}

	var data = make([]byte, fi.Size())
	if _, err := file.Read(data); err != nil {
		return
	}
	fmt.Println(string(data))
}

func FileIO() {
	file, err := os.OpenFile(
		"test2.txt",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC, // 없으면 생성, 있으면 내용 지움
		os.FileMode(0644),
	)
	if err != nil {
		return
	}
	defer file.Close()

	if _, err := file.Write([]byte("안녕하세요")); err != nil {
		return
	}

	var fi fs.FileInfo
	if fi, err = file.Stat(); err != nil {
		return
	}
	data := make([]byte, fi.Size()) // io.SEEK_SET is deprecated
	file.Seek(0, io.SeekStart)

	if _, err = file.Read(data); err != nil {
		return
	}

	fmt.Println(string(data))
}

// ioutil is deprecated. Replaced with os
func Os() {
	s := "Hello, world!"

	err := os.WriteFile("test3.txt", []byte(s), os.FileMode(644))
	if err != nil {
		return
	}

	data, err := os.ReadFile("test3.txt")
	if err != nil {
		return
	}

	fmt.Println(string(data))
}

func main() {
	File()
	FileIO()
	Os()
}
