package test

/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1：

输入：strs = ["flower","flow","flight"]
输出："fl"
示例 2：

输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。

整体思路：
	1.先取出切片中第一个元素，进行遍历
	2.第一个元素中第一个字符和第二个元素的第一个字符对比，如果相同，则和第三个元素第一个字符对比。。。
	3.如果都匹配上了，则进行第二个字符依次比较。步骤同第一步一样。。。。
	4.假如比对到第j个字符时，有匹配不上的，则说明j列前面的是匹配上了的
*/
func LongestCommonPrefix(strs []string) string {

	firstStr := strs[0] //取出第一个元素

	for j, v := range firstStr {
		for _, str := range strs {
			if j == len(str) || str[j] != byte(v) { //比对第j个字符是否相同，或者str是否字母缺失
				return firstStr[:j]
			}
		}
	}

	return firstStr
}
