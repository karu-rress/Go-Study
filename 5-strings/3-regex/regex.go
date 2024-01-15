package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 한 번만 사용할 때
	matched, _ := regexp.MatchString("[0-9]+\\+[0-9]*", "1+")
	fmt.Println(matched) // true

	// 여러 번 사용할 때라면 컴파일
	var validEmail, _ = regexp.Compile(
		"^[_a-z0-9+-.]+@[[a-z0-9-]+(?P<domain1>\\.[a-z0-9-]+)*(?P<domain2>\\.[a-z]{2,4})$",
	) // ${domain1}.${domain2}

	fmt.Println(validEmail.MatchString("myname@gmail.com"))

	s2 := validEmail.FindStringSubmatch("karu@rolling.ress")
	fmt.Println(s2)

	/*
		기타 함수
		FindString: 찾은 문자열
		FindStringSubmatch: 찾은 문자열, 그룹
		FindStringSubmatchIndex: 위치
		FindAllString: 모든 문자열을 슬라이스로
		FindAllStringIndex: 모든 문자열의 위치를 슬라이스로

		ReplaceAllString: 바꾸기
		ReplaceAllLiteralString: 정규표현식 기호 무시
		Split: 나누기
	*/
}