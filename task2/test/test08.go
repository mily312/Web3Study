package test

import (
	"fmt"
	"sync"
)

/*
编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。
启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值
*/
var mutex sync.Mutex
var wg sync.WaitGroup

var count = 0

func MutexPrintNum() {
	mutex.Lock()
	for i := 0; i < 1000; i++ {
		count++
	}
	mutex.Unlock()

	wg.Done()
}

func Test08() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go MutexPrintNum()
	}

	wg.Wait()
	fmt.Println(count)
}


