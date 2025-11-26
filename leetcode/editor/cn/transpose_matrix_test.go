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
	// 获取原矩阵的维度
	m := len(matrix)    // 原矩阵的行数
	n := len(matrix[0]) // 原矩阵的列数（假设矩阵非空且每行长度相同）

	// 创建结果矩阵（转置后的矩阵）
	// 转置后维度变为 n × m，所以先创建 n 行
	res := make([][]int, n)

	// 初始化结果矩阵的每一行
	// 每一行需要有 m 个元素（对应原矩阵的 m 行）
	for i := 0; i < n; i++ {
		res[i] = make([]int, m)
	}

	// 执行转置操作
	// 遍历原矩阵的每一个元素
	for i := 0; i < m; i++ { // i: 原矩阵的行索引
		for j := 0; j < n; j++ { // j: 原矩阵的列索引
			// 关键转换：原矩阵 [i][j] → 转置矩阵 [j][i]
			// 原来的第 i 行第 j 列元素，变成转置矩阵的第 j 行第 i 列
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
