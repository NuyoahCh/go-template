/*
 * @lc app=leetcode.cn id=1260 lang=golang
 * @lcpr version=30304
 *
 * [1260] 二维网格迁移
 */

package leetcode_solutions

import (
	"fmt"
	"testing"
)

// @lc code=start

func shiftGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	mn := m * n
	k = k % mn
	// 把二维 grid 抽象成一维数组
	reverse02(grid, mn-k, mn-1)
	// 先把最后 k 个数翻转
	reverse02(grid, 0, mn-k-1)
	// 然后把前 mn - k 个数翻转
	reverse02(grid, 0, mn-1)
	// 最后把整体翻转

	// 转化成 Java List
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
		for j := 0; j < n; j++ {
			res[i][j] = grid[i][j]
		}
	}
	return res
}

// 通过一维数组的索引访问二维数组的元素
func get(grid [][]int, index int) int {
	n := len(grid[0])
	i, j := index/n, index%n
	return grid[i][j]
}

// 通过一维数组的索引修改二维数组的元素
func set(grid [][]int, index int, val int) {
	n := len(grid[0])
	i, j := index/n, index%n
	grid[i][j] = val
}

// 翻转一维数组的索引区间元素
func reverse02(grid [][]int, i, j int) {
	for i < j {
		temp := get(grid, i)
		set(grid, i, get(grid, j))
		set(grid, j, temp)
		i++
		j--
	}
}

// @lc code=end

func TestShift2dGrid(t *testing.T) {
	// your test code here
	num := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	k := 9
	fmt.Printf("shiftGrid(num, k): %v\n", shiftGrid(num, k))
}

/*
// @lcpr case=start
// [[1,2,3],[4,5,6],[7,8,9]]\n1\n
// @lcpr case=end

// @lcpr case=start
// [[3,8,1,9],[19,7,2,5],[4,6,11,10],[12,0,21,13]]\n4\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,3],[4,5,6],[7,8,9]]\n9\n
// @lcpr case=end

*/
