package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const Message = "Hello, world!"

func BufferedFile() {
	file, err := os.OpenFile(
		"test.txt",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		os.FileMode(0644),
	)
	if err != nil {
		return
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	w.WriteString(Message)
	w.Flush()

	r := bufio.NewReader(file)
	fi, _ := file.Stat()
	b := make([]byte, fi.Size())

	file.Seek(0, io.SeekStart)
	r.Read(b)

	fmt.Println(string(b))
}

func String() {
	file, err := os.OpenFile(
		"test2.txt",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		os.FileMode(0644),
	)
	if err != nil {
		return
	}
	defer file.Close()

	s := Message
	r := strings.NewReader(s) // 문자열로 io.Reader 인터페이스를 따름

	w := bufio.NewWriter(file)

	w.ReadFrom(r)            // r(s)에서 데이터를 읽어 w 버퍼에 저장
	fmt.Fprintf(w, "%d", 10) // Fprintf 쓸 수도 있음
	w.Flush()                // 파일에 저장
}

func Stdio() {
	s := Message
	r := strings.NewReader(s)

	io.Copy(os.Stdout, r)
}

func ReadWriter() {
	file, err := os.OpenFile(
		"test3.txt",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		os.FileMode(0644),
	)
	if err != nil {
		return
	}
	defer file.Close()

	r := bufio.NewReader(file)
	w := bufio.NewWriter(file)
	rw := bufio.NewReadWriter(r, w)

	rw.WriteString(Message)
	rw.Flush()

	file.Seek(0, io.SeekStart)
	data, _, _ := rw.ReadLine() // 한 줄을 읽어서 저장
	fmt.Println(string(data))
}

func main() {
	BufferedFile()
	String()
	Stdio()
	ReadWriter()
}
