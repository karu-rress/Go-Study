package main

import (
	"fmt"
	"os"
)

func Sstream() {
	s1 := fmt.Sprint(1, 1.1, "Hi")
	fmt.Println(s1)

	i1 := "1\n1.1\nHello"
	var num int
	var f float32
	var s string
	n, _ := fmt.Sscan(i1, &num, &f, &s)
	fmt.Println("입력 개수: ", n)
	fmt.Println(num, f, s)
}

func Fstream() {
	file1, _ := os.Create("test.txt")
	
	fmt.Fprintln(file1, 1, 1.1, "Hi")
	file1.Close()

	file1, _ = os.Open("test.txt")
	defer file1.Close()

	var num int
	var f float32
	var s string
	n, _ := fmt.Fscan(file1, &num, &f, &s)

	fmt.Println("입력 개수: ", n)
}

func main() {
	Sstream()
	Fstream()
}