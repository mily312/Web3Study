package main

import (
	"fmt"
	"strconv"
	"task2/test"
)

func main() {
	// test01()
	// test02()

	// test.SaveAndPrintNum()
	// test.PrintNumRoutine()
	// test.Test06()
	// test.Test07()
	//test.Test08()

	test.Test09()

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
