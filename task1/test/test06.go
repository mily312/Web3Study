package test

import "slices"

/*
示例 1：

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2：

输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
示例 3：

输入：intervals = [[4,7],[1,4]]
输出：[[1,7]]
解释：区间 [1,4] 和 [4,7] 可被视为重叠区间。
*/

func merge(intervals [][]int) [][]int {

	var result [][]int

	// 按照左端点从小到大排序
	slices.SortFunc(intervals, func(p, q []int) int {
		return p[0] - q[0]
	})

	for _, intArray := range intervals {
		m := len(result)

		if m != 0 && result[m-1][1] >= intArray[0] { //区间合并
			result[m-1][1] = max(result[m-1][1], intArray[1]) //右端点取最大值
		} else {
			result = append(result, intArray)
		}

	}

	return result
}
