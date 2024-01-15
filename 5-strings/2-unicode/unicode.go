package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func Unicode() {
	fmt.Println(unicode.Is(unicode.Hangul, '한'))
	fmt.Println(unicode.Is(unicode.Latin, '한'))

	fmt.Println(unicode.Is(unicode.Han, '漢'))
	fmt.Println(unicode.Is(unicode.Hangul, '漢'))

	fmt.Println(unicode.Is(unicode.Latin, 'A'))
	fmt.Println(unicode.Is(unicode.Hangul, 'A'))

	fmt.Println(unicode.In('한', unicode.Han, unicode.Hangul, unicode.Latin))

	///

	fmt.Println(unicode.IsGraphic('한')) // 표시되는 문자인가?
	fmt.Println(unicode.IsGraphic('\n'))

	fmt.Println(unicode.IsLetter('A')) // 문자인가?
	fmt.Println(unicode.IsLetter('1'))

	fmt.Println(unicode.IsDigit('1'))    // 숫자인가?
	fmt.Println(unicode.IsControl('\r')) // 제어문자인가?

	fmt.Println(unicode.IsMark('\u17c9')) // 마크인가?
	fmt.Println(unicode.IsPrint('1'))     // 표시할 수 있는가?
	fmt.Println(unicode.IsPunct('.'))     // 문장부호인가?

	fmt.Println(unicode.IsSpace(' '))  // 공백인가?
	fmt.Println(unicode.IsSymbol('★')) // 심볼인가?

	fmt.Println(unicode.IsUpper('A')) // 대문자인가?
	fmt.Println(unicode.IsLower('a')) // 소문자인가?
}

func Utf8() {
	s1 := "안녕하세요"
	// len(s1): 한글은 3바이트로 저장하므로 15

	fmt.Println(utf8.RuneCountInString(s1)) // 실제 길이

	b := []byte(s1)
	r, size := utf8.DecodeRune(b)
	fmt.Println(r, size)

	r, size = utf8.DecodeRune(b[3:])
	fmt.Println(r, size)

	r, size = utf8.DecodeLastRune(b)
	fmt.Println(r, size)

	r, size = utf8.DecodeLastRune(b[:len(b)-3])
	fmt.Println(r, size)

	// r = s1[0] 한글은 3바이트이므로 이런 식으로 처리하면 안 됨!
	r, _ = utf8.DecodeRuneInString(s1)
	fmt.Printf("%c\n", r)

	fmt.Println(utf8.Valid(b)) // 올바른 UTF-8 문자열인가? ValidRune, ValidString, ...
}

func main() {
	Unicode()
	Utf8()
}
