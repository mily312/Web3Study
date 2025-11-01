package test

/*
题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
*/
func MultiplicationNum(intArray *[]int) {
	for i := 0; i < len(*intArray); i++ {
		arr := *intArray
		arr[i] = arr[i] * 2
	}
}
