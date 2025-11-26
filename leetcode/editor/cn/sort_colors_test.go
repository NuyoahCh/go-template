/*
 * @lc app=leetcode.cn id=75 lang=golang
 * @lcpr version=30304
 *
 * [75] 颜色分类
 */

package leetcode_solutions

import (
	"testing"
)

// @lc code=start
func sortColors(nums []int) {
	// sort.Ints(nums)
	p0, p2 := 0, len(nums)-1
	p := 0
	for p <= p2 {
		if nums[p] == 0 {
			swap(nums, p0, p)
			p0++
		} else if nums[p] == 2 {
			swap(nums, p2, p)
			p2--
		} else if nums[p] == 1 {
			p++
		}

		// 因为小于 p0 都是 0.所以 p 不要小于 p0
		if p < p0 {
			p = p0
		}
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

// @lc code=end

func TestSortColors(t *testing.T) {
	// your test code here
	num := []int{2, 0, 2, 1, 1, 0}
	sortColors(num)
}

/*
// @lcpr case=start
// [2,0,2,1,1,0]\n
// @lcpr case=end

// @lcpr case=start
// [2,0,1]\n
// @lcpr case=end

*/
