package main

import (
	"fmt"
	"time"
)

func main() {

	/**
	给定一个字符串数组
	[“I”,“am”,“stupid”,“and”,“weak”]
	用 for 循环遍历该数组并修改为
	[“I”,“am”,“smart”,“and”,“strong”]
	*/
	myarr := [5]string{"i", "am", "stuid", "and", "week"}
	newarr := [5]string{"i", "am", "smart", "and", "strong"}
	for i := 0; i < len(myarr); i++ {
		myarr[i] = newarr[i]
	}

	for _, value := range myarr {
		fmt.Printf("%s,", value)
	}
	fmt.Println()

	/**	队列：
	队列长度 10，队列元素类型为 int
	生产者：
	每 1 秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞
	消费者：
	每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞
	*/
	message := make(chan int, 10)
	done := make(chan int)
	defer close(message)

	// consumer
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-done:
				fmt.Println("consumer done")
				return
			default:
				fmt.Println("revice message: %d", <-message)
			}
		}
	}()

	// provider
	for i := 0; i < 10; i++ {
		message <- i * 10
		fmt.Println("send message: %d", i*10)
		time.Sleep(1 * time.Second)
	}
	time.Sleep(5 * time.Second)
	close(done)
	time.Sleep(1 * time.Second)
	fmt.Println("main process exit!")

}
