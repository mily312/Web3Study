package test

/*
示例 1：

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
示例 2：

输入：nums = [3,2,4], target = 6
输出：[1,2]
示例 3：

输入：nums = [3,3], target = 6
输出：[0,1]
*/

func TwoSum(nums []int, target int) []int {

	//map[数组中元素值]数组下标
	mapIndex := map[int]int{}

	for k, v := range nums {
		if index, ok := mapIndex[target-v]; ok {

			return []int{k, index}
		}

		mapIndex[v] = k

	}
	return nil

}
