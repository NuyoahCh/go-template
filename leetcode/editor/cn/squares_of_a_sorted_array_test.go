/*
 * @lc app=leetcode.cn id=977 lang=golang
 * @lcpr version=30304
 *
 * [977] 有序数组的平方
 */

package leetcode_solutions

import (
	"fmt"
	"testing"
)

// @lc code=start
func sortedSquares(nums []int) []int {
	// 初始化
	n := len(nums)
	i, j, p := 0, n-1, n-1
	// 这个地方一定要去新创建一个数组，不然都冗余到一起了
	res := make([]int, n)
	for i <= j {
		if abs(nums[i]) >= abs(nums[j]) {
			res[p] = nums[i] * nums[i]
			i++
		} else {
			res[p] = nums[j] * nums[j]
			j--
		}
		p--
	}
	return res
}

// helper func : abs 绝对值函数
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @lc code=end

func TestSquaresOfASortedArray(t *testing.T) {
	// your test code here
	num := []int{-4, -1, 0, 3, 10}
	fmt.Print(sortedSquares(num))
}

/*
// @lcpr case=start
// [-4,-1,0,3,10]\n
// @lcpr case=end

// @lcpr case=start
// [-7,-3,2,3,11]\n
// @lcpr case=end

*/
