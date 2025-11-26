/*
 * @lc app=leetcode.cn id=54 lang=golang
 * @lcpr version=30304
 *
 * [54] 螺旋矩阵
 */

package leetcode_solutions

import (
	"fmt"
	"testing"
)

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	var res []int
	m, n := len(matrix), len(matrix[0])
	upper_bound, lower_bound := 0, m-1
	left_bound, right_bound := 0, n-1
	for len(res) < m*n {
		if upper_bound <= lower_bound {
			for j := left_bound; j <= right_bound; j++ {
				res = append(res, matrix[upper_bound][j])
			}
			upper_bound++
		}
		if left_bound <= right_bound {
			for i := upper_bound; i <= lower_bound; i++ {
				res = append(res, matrix[i][right_bound])
			}
			right_bound--
		}
		if upper_bound <= lower_bound {
			for j := right_bound; j >= left_bound; j-- {
				res = append(res, matrix[lower_bound][j])
			}
			lower_bound--
		}
		if left_bound <= right_bound {
			for i := lower_bound; i >= upper_bound; i-- {
				res = append(res, matrix[i][left_bound])
			}
			left_bound++
		}
	}
	return res
}

// @lc code=end

func TestSpiralMatrix(t *testing.T) {
	// your test code here
	num := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Printf("spiralOrder(num): %v\n", spiralOrder(num))
}

/*
// @lcpr case=start
// [[1,2,3],[4,5,6],[7,8,9]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,3,4],[5,6,7,8],[9,10,11,12]]\n
// @lcpr case=end

*/
