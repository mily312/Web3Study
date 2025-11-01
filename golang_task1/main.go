package main

import (
	"fmt"
	"task1/test"
)

func main() {
	//test.FindOneElement()
	//fmt.Println(test.IsValid("()[]{}"))

	var strs = []string{"flower", "flow", "flight"}
	result := test.LongestCommonPrefix(strs)
	fmt.Println(result)
}
