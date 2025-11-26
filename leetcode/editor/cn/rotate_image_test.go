/*
 * @lc app=leetcode.cn id=48 lang=golang
 * @lcpr version=30304
 *
 * [48] 旋转图像
 */

package leetcode_solutions

import (
	"testing"
)

// @lc code=start
func rotate(matrix [][]int) {
	n := len(matrix)
	if n == 0 {
		return
	}
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for _, row := range matrix {
		reverseNum(row)
	}
}

func reverseNum(arr []int) {
	i, j := 0, len(arr)-1
	for j > i {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}

}

// @lc code=end

func TestRotateImage(t *testing.T) {
	// your test code here
	num := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(num)
}

/*
// @lcpr case=start
// [[1,2,3],[4,5,6],[7,8,9]]\n
// @lcpr case=end

// @lcpr case=start
// [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]\n
// @lcpr case=end

*/
