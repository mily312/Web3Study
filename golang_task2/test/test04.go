package test

import (
	"fmt"
	"sync"
)

/*
编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
*/
var wg4 sync.WaitGroup

func printNum1() {
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Printf("协程1打印奇数：%v\n", i)
		}
	}

	wg4.Done()
}

func printNum2() {
	for i := 2; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("协程2打印偶数：%v\n", i)
		}
	}

	wg4.Done()
}

func PrintNumRoutine() {

	wg4.Add(1)
	go printNum1()

	wg4.Add(1)
	go printNum2()

	wg4.Wait()
}
