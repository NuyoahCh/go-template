/*
 * @lc app=leetcode.cn id=867 lang=golang
 * @lcpr version=30304
 *
 * [867] 转置矩阵
 */

package leetcode_solutions

import (
	"fmt"
	"testing"
)

// @lc code=start
// transpose 返回输入矩阵的转置矩阵
// 转置矩阵的定义：原矩阵的行变成列，列变成行
// 原矩阵 dimensions: m × n → 转置矩阵 dimensions: n × m
func transpose(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, m)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[j][i] = matrix[i][j]
		}
	}
	return res
}

// @lc code=end

func TestTransposeMatrix(t *testing.T) {
	// your test code here
	num := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Printf("transpose(num): %v\n", transpose(num))
}

/*
// @lcpr case=start
// [[1,2,3],[4,5,6],[7,8,9]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,3],[4,5,6]]\n
// @lcpr case=end

*/
