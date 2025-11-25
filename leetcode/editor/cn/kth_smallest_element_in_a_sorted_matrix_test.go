/*
 * @lc app=leetcode.cn id=378 lang=golang
 * @lcpr version=30304
 *
 * [378] 有序矩阵中第 K 小的元素
 */

package leetcode_solutions

import (
	"fmt"
	"testing"
)

// @lc code=start
func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	// 初始化二分查找的范围：最小值为矩阵左上角元素，最大值为矩阵右下角元素
	left, right := matrix[0][0], matrix[n-1][n-1]
	// 进行二分查找，直到左右边界重合
	for left < right {
		mid := left + (right-left)/2 // 防止溢出，计算中间值
		// 检查矩阵中小于等于mid的元素个数是否大于等于k
		if check(matrix, mid, k, n) {
			right = mid // 如果满足，说明第k小元素在mid左侧（含mid）
		} else {
			left = mid + 1 // 否则说明第k小元素在mid右侧
		}
	}
	// 最终left即为第k小元素
	return left
}

func check(matrix [][]int, mid, k, n int) bool {
	// 从左下角开始扫描：初始行为最后一行(i=n-1)，列为第一列(j=0)
	i, j := n-1, 0
	num := 0 // 统计小于等于mid的元素个数

	for i >= 0 && j < n {
		if matrix[i][j] <= mid {
			// 当前元素小于等于mid，则该列当前行及其上方的所有元素均小于等于mid
			num += i + 1 // 累加该列中满足条件的元素个数（共i+1个）
			j++          // 向右移动一列（检查下一列）
		} else {
			// 当前元素大于mid，则当前行该列及右侧列都不满足条件
			i-- // 向上移动一行（检查上一行）
		}
	}
	// 判断小于等于mid的元素总数是否大于等于k
	return num >= k
}

// @lc code=end

func TestKthSmallestElementInASortedMatrix(t *testing.T) {
	// your test code here
	matrix := [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}
	k := 8
	fmt.Println(kthSmallest(matrix, k))

}

/*
// @lcpr case=start
// [[1,5,9],[10,11,13],[12,13,15]]\n8\n
// @lcpr case=end

// @lcpr case=start
// [[-5]]\n1\n
// @lcpr case=end

*/
