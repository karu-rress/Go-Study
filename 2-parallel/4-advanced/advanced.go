package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
)

func Cond() {
	var mutex = new(sync.Mutex)
	var cond = sync.NewCond(mutex)

	c := make(chan bool, 3)
	wg := new(sync.WaitGroup) // 대기 그룹 생성

	for i := 0; i < 3; i++ {
		wg.Add(1) // 반복할 때마다 1씩 추가
		go func(n int) {
			defer wg.Done() // 지연호출: 고루틴이 끝나기 직전에 실행
			mutex.Lock()
			c <- true
			fmt.Println("wait begin : ", n)
			cond.Wait() // 조건 변수 대기
			fmt.Println("wait end : ", n)
			mutex.Unlock()
		}(i) // 고루틴이니까 i 직접 사용 안 하고 매개변수로!
	}

	for i := 0; i < 3; i++ {
		<-c // 고루틴 3개가 모두 실행될 때까지 기다림
	}

	for i := 0; i < 3; i++ {
		mutex.Lock()
		fmt.Println("signal : ", i)
		cond.Signal() // 대기하고 있는 고루틴을 하나씩 깨움
		// 반복문 없이 한 번에 깨우려면 Broadcast 사용
		// cond.Broadcast()
		mutex.Unlock()
	}

	wg.Wait() // 모든 고루틴이 끝날 때까지 대기
}

func Once() {
	once := new(sync.Once)
	wg := new(sync.WaitGroup)
	hello := func() {
		fmt.Println("Hello!")
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println("goroutine: ", n)
			once.Do(hello)
		}(i)
	}

	wg.Wait()
}

type Data struct {
	tag    string
	buffer []int
}

func Pool() {
	wg := new(sync.WaitGroup)
	pool := sync.Pool{
		New: func() interface{} { // Get 사용시 호출될 함수 정의
			data := new(Data)
			data.tag = "new"
			data.buffer = make([]int, 10)
			return data
		},
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			data := pool.Get().(*Data) // 풀에서 *Data 타입으로 데이터를 가져옴
			for index := range data.buffer {
				data.buffer[index] = rand.Intn(100) // 랜덤 값 저장
			}
			fmt.Println(data)
			data.tag = "used" // 객체가 사용되었다는 태그 설정
			pool.Put(data)    // 풀에 객체 보관
		}()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			data := pool.Get().(*Data) // 풀에서 *Data 타입으로 데이터를 가져옴

			n := 0
			for index := range data.buffer {
				data.buffer[index] = n

				n += 2
			}
			fmt.Println(data)
			data.tag = "used" // 객체가 사용되었다는 태그 설정
			pool.Put(data)    // 풀에 객체 보관
		}()
	}

	wg.Wait()
}

// 사용할 수 있다면 Mutex에 비해 Atomic이 성능상 훨씬 좋음
func Atomic() {
	var data int32 = 0
	wg := new(sync.WaitGroup)

	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&data, 1)
			wg.Done()
		}()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&data, -1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(data)
}

func main() {
	Cond()
	Once()
	Pool()
	Atomic()
}
