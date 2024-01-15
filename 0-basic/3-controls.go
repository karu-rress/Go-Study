package main

import (
	"fmt"
	"math/rand"
)

func ForSwitch() {
	// Break, continue with label
Loop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 2 {
				continue Loop
				// break Loop
			}
			fmt.Println(i, j)
		}
	}
	fmt.Print("Ended nested for loop\n\n")

	// 변수만 선언하고 조건식은 없음! (세미콜론)
	switch i := rand.Intn(10); {
	case i <= 5:
		fmt.Println("i is less than or equal to 5")
	case i >= 6:
		fmt.Println("i is greater than or equal to 6")
	}
}
