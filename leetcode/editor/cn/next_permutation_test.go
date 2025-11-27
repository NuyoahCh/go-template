/*
 * @lc app=leetcode.cn id=31 lang=golang
 * @lcpr version=30304
 *
 * [31] 下一个排列
 */

package leetcode_solutions

import (
	"slices"
	"testing"
)

// @lc code=start
func nextPermutation(nums []int) {
	n := len(nums)
	// i+1<n-1
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := n - 1
		for nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	slices.Reverse(nums[i+1:])
}

// @lc code=end

func TestNextPermutation(t *testing.T) {
	// your test code here
	num := []int{1, 2, 3}
	nextPermutation(num)
}

/*
// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [3,2,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,5]\n
// @lcpr case=end

*/
