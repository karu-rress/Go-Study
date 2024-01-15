package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)

	title := flag.String("title", "", "영화 이름")
	runtime := flag.Int("runtime", 0, "상영 시간")
	rating := flag.Float64("rating", 0.0, "평점")
	released := flag.Bool("released", false, "개봉 여부")

	// 명령줄 옵션 분석
	flag.Parse()

	// 명령줄 옵션이 없다면, 사용법 출력
	if flag.NFlag() == 0 {
		flag.Usage()
		return
	}

	fmt.Printf(
		"영화 이름: %s\n상영 시간: %d분\n평점 %f\n",
		*title,
		*runtime,
		*rating,
	) // 포인터임.

	if *released {
		fmt.Println("개봉 여부: 개봉")
	} else {
		fmt.Println("개봉 여부: 미개봉")
	}
}
