package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func Mutex() {
	var data = []int{}
	var mutex = new(sync.Mutex)

	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock()
			data = append(data, i)
			mutex.Unlock()

			runtime.Gosched() // 다른 고루틴에게 CPU 양보
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock()
			data = append(data, i)
			mutex.Unlock()

			runtime.Gosched() // 다른 고루틴에게 CPU 양보
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println(len(data))
	fmt.Println()
}

func RwMutex() {
	var data int = 0
	var rwMutex = new(sync.RWMutex)

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.Lock() // 쓰기 잠금
			data += 1
			fmt.Println("write: ", data)
			time.Sleep(10 * time.Millisecond)
			rwMutex.Unlock()
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.RLock()
			fmt.Println("read 1: ", data)
			time.Sleep(1 * time.Second)
			rwMutex.RUnlock()
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.RLock()
			fmt.Println("read 2: ", data)
			time.Sleep(1 * time.Second)
			rwMutex.RUnlock()
		}
	}()

	time.Sleep(10 * time.Second)
}

func main() {
	Mutex()
	RwMutex()
}
