package test

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
使用原子操作（ sync/atomic 包）实现一个无锁的计数器。
启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值
*/

var wg09 sync.WaitGroup

var count09 int32

func AtomicPrintNum() {

	for i := 0; i < 1000; i++ {
		atomic.AddInt32(&count09, 1)
	}

	wg09.Done()
}

func Test09() {
	for i := 0; i < 10; i++ {
		wg09.Add(1)
		go AtomicPrintNum()
	}

	wg09.Wait()
	fmt.Println(count09)
}
