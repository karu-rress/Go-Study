package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const Hello = "Hello, world!"

func String() {
	fmt.Println(strings.Contains(Hello, "wo"))    // 포함하는가?
	fmt.Println(strings.ContainsAny(Hello, "wo")) // 하나라도?
	fmt.Println(strings.Count(Hello, "o"))        // 몇개나?

	var r rune = '하'
	fmt.Println(strings.ContainsRune("안녕하세요", r)) // Rune 문자
	fmt.Println(strings.HasPrefix(Hello, "He"))   // 접두사가 있는가?
	fmt.Println(strings.HasSuffix(Hello, "rld!")) // 접미사가 있는가?

	fmt.Println(strings.Index(Hello, "ll"))     // 인덱스는?
	fmt.Println(strings.IndexAny(Hello, "elo")) // 하나라도?
	// IndexRune, IndexByte도 있음

	f := func(r rune) bool {
		return unicode.Is(unicode.Hangul, r) // 한글이면 true
	}
	fmt.Println(strings.IndexFunc("Go 언어", f)) // 함수값이 참인 게 있는가?

	/*
		문자열 조작 함수
		Join: 모두 연결
		Split: 구분자로 쪼개서 슬라이스로 리턴
		Fields: 공백을 기준으로 쪼개서 슬라이스로
		FieldsFunc: 함수값을 기준으로 쪼개서 슬라이스로
		Repeat: 문자열 반복
		Replace: 바꾸기
		Trim: 특정 문자열 제거
		TrimLeft, TrimRight: 왼쪽, 오른쪽 문자열 제거
		TrimFunc: 함수값을 기준으로 제거
		TrimLeftFunc, TrimRightFunc: 상동
		TrimSpace: 공백제거
		TrimSuffix: 접미사 제거
		NewReplacer: 교체 인스턴스 생성
		Replace: NewReplacer로 만든 걸로 문자열 교체
	*/
}

func Converts() {
	var n int
	n, err := strconv.Atoi("100")
	fmt.Println(n, err)
	/*
		FormatBool
		FormatFloat
		FormatInt
		FormatUint
		AppendBool: 불값을 문자열로 변환, 슬라이스 뒤에 추가
		AppendFloat
		AppendInt
		AppendUint
		ParseBool: 불 형태 문자열을 불로 반환
		ParseFloat
		ParseInt
		ParseUint
	*/
}

func main() {
	String()
	Converts()
}
