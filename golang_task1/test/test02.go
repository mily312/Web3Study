package test

import "fmt"

/*有效的括号

考察：字符串处理、栈的使用

题目：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效

([)]

整体思路：1.遇到左边元素往切片里面加
	2.遇到右边元素：1）先判断切片里有没有元素，如果是符合要求的，肯定会有对应的左面元素
					2）再判断切片里的最后一个元素，和这个右面元素是不是一对
*/
func IsValid(str string) bool {

	if len(str)%2 != 0 {
		fmt.Println("无效的括号！！")
		return false
	}

	//构建符号对
	mapParis := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	stSlice := []rune{}

	//遍历str
	for _, c := range str {
		if mapParis[c] == 0 { //说明是左边的  {[(
			stSlice = append(stSlice, c)
		} else { // c是右边的
			// ({}) ([})
			// 没有左边元素，或者左边元素不对
			if len(stSlice) == 0 || stSlice[len(stSlice)-1] != mapParis[c] {
				return false
			}

			//出栈，去掉最后一个元素
			stSlice = stSlice[:len(stSlice)-1]
		}
	}

	return len(stSlice) == 0
}
