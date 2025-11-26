/*
 * @lc app=leetcode.cn id=1329 lang=golang
 * @lcpr version=30304
 *
 * [1329] 将矩阵按对角线排序
 */

package leetcode_solutions

import (
	"fmt"
	"sort"
	"testing"
)

// @lc code=start
func diagonalSort(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	diagonals := make(map[int][]int)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			diagonalID := i - j
			diagonals[diagonalID] = append(diagonals[diagonalID], mat[i][j])
		}
	}

	for _, didiagonal := range diagonals {
		sort.Slice(didiagonal, func(i, j int) bool {
			return didiagonal[i] > didiagonal[j]
		})
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			diagonalID := i - j
			mat[i][j] = diagonals[diagonalID][len(diagonals[diagonalID])-1]
			diagonals[diagonalID] = diagonals[diagonalID][:len(diagonals[diagonalID])-1]
		}
	}
	return mat
}

// @lc code=end

func TestSortTheMatrixDiagonally(t *testing.T) {
	// your test code here
	num := [][]int{{3, 3, 1, 1}, {2, 2, 1, 2}, {1, 1, 1, 2}}
	fmt.Printf("diagonalSort(num): %v\n", diagonalSort(num))
}

/*
// @lcpr case=start
// [[3,3,1,1],[2,2,1,2],[1,1,1,2]]\n
// @lcpr case=end

// @lcpr case=start
// [[11,25,66,1,69,7],[23,55,17,45,15,52],[75,31,36,44,58,8],[22,27,33,25,68,4],[84,28,14,11,5,50]]\n
// @lcpr case=end

*/
