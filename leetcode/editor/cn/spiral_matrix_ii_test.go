/*
 * @lc app=leetcode.cn id=59 lang=golang
 * @lcpr version=30304
 *
 * [59] 螺旋矩阵 II
 */

package leetcode_solutions

import (
	"fmt"
	"testing"
)

// @lc code=start
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	upper_bound, lower_bound := 0, n-1
	left_bound, right_bound := 0, n-1
	num := 1
	for num <= n*n {
		if upper_bound <= lower_bound {
			for j := left_bound; j <= right_bound; j++ {
				matrix[upper_bound][j] = num
				num++
			}
			upper_bound++
		}
		if left_bound <= right_bound {
			for i := upper_bound; i <= lower_bound; i++ {
				matrix[i][right_bound] = num
				num++
			}
			right_bound--
		}
		if upper_bound <= lower_bound {
			for j := right_bound; j >= left_bound; j-- {
				matrix[lower_bound][j] = num
				num++
			}
			lower_bound--
		}
		if left_bound <= right_bound {
			for i := lower_bound; i >= upper_bound; i-- {
				matrix[i][left_bound] = num
				num++
			}
			left_bound++
		}
	}
	return matrix
}

// @lc code=end

func TestSpiralMatrixIi(t *testing.T) {
	// your test code here
	n := 3
	fmt.Printf("generateMatrix(n): %v\n", generateMatrix(n))
}

/*
// @lcpr case=start
// 3\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/
