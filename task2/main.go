package main

import (
	"fmt"
	"strconv"
	"task2/test"
)

func main() {
	// test01()
	// test02()

	//test.SaveAndPrintNum()
	test.PrintNumRoutine()
}

func test01() {
	value := 25
	test.AddNum(&value)
	fmt.Println("value:" + strconv.Itoa(value))
}

func test02() {
	intArray := []int{21, 321, -15, 28}
	test.MultiplicationNum(&intArray)
	fmt.Println(intArray)
}
