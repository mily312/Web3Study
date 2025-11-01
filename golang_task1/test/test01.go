package test

import "fmt"

/*
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
找出那个只出现了一次的元素。
可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
*/
func FindOneElement() {
	var arrays [9]int = [9]int{13, 5, 21, 24, 13, 8, 5, 24, 21}

	//map记录出现的次数
	var mapCount = make(map[int]int)

	for i := 0; i < len(arrays); i++ {
		value := arrays[i]
		mapCount[value] += 1
	}

	for k, v := range mapCount {
		if v == 1 {
			fmt.Println(k)
		}
	}
}
