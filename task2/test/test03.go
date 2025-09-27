package test

import (
	"fmt"
	"sync"
	"time"
)

/*
编写一个程序，使用通道实现两个协程之间的通信。
一个协程生成从1到10的整数，并将这些整数发送到通道中，
另一个协程从通道中接收这些整数并打印出来
*/

var wg3 sync.WaitGroup

// 发送数字到管道
func saveNumRoutine(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Printf("[写入]到管道:%v\n", i)
		time.Sleep(time.Millisecond * 100)
	}

	close(ch)

	wg3.Done()
}

// 打印数字
func printNumRoutine(ch chan int) {
	for v := range ch {
		fmt.Printf("[读取]打印:%v\n", v)
		time.Sleep(time.Millisecond * 50)
	}

	wg3.Done()
}

func SaveAndPrintNum() {

	var ch = make(chan int, 10)

	wg3.Add(1)
	go saveNumRoutine(ch)

	wg3.Add(1)
	go printNumRoutine(ch)

	wg3.Wait()
}
