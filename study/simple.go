package main

import (
	"fmt"
	"time"
)

func main() {
	//带缓冲的channel
	ch := make(chan int, 3)
	numWorkers := 10
	//启动固定数量的worker
	for i := 0; i < numWorkers; i++ {
		go worker(ch)
	}

	//发送任务给worker
	hellaTasks := getTasks()
	for task := range hellaTasks {
		fmt.Println("发送")
		fmt.Println(task)
		ch <- task
	}

	time.Sleep(time.Second)
}
func worker(ch chan int) {
	for {
		//接受任务
		task := <-ch
		fmt.Println("执行")
		fmt.Println(task)
		time.Sleep(time.Second)
	}
}

func getTasks() []int {
	return []int{1, 2, 3}
}
