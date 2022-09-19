package main

import (
	"fmt"
	"math/rand"
	"time"

	//"math/rand"

	//"math/rand"
	"sync"
	//"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go fmt.Println(i)
	}
	time.Sleep(time.Second)

	ch := make(chan int)
	go func() {
		fmt.Println("hello from goroutine")
		ch <- 3 //数据写入Channel
	}()
	k := <-ch //从Channel中取数据并赋值
	fmt.Println("=====", k)

	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	loopFunc()
	time.Sleep(time.Second)

	ch1 := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			rand.Seed(time.Now().UnixNano())
			n := rand.Intn(10) // n will be between 0 and 10
			fmt.Println("putting: ", n)
			ch1 <- i
		}
		close(ch1)
	}()
	fmt.Println("hello from main")
	for v := range ch1 {
		fmt.Println("receiving: ", v)
	}

}

func loopFunc() {
	lock := sync.Mutex{}
	for i := 0; i < 3; i++ {
		go func(i int) {
			lock.Lock()
			defer lock.Unlock()
			fmt.Println("loopFunc:", i)
		}(i)
	}
}
