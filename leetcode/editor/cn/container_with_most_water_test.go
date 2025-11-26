/*
 * @lc app=leetcode.cn id=11 lang=golang
 * @lcpr version=30304
 *
 * [11] 盛最多水的容器
 */

package leetcode_solutions

import (
	"fmt"
	"testing"
)

// @lc code=start
func maxArea(height []int) int {
	res := 0
	left, right := 0, len(height)-1
	for left <= right {
		area := min(height[left], height[right]) * (right - left)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
		res = max(res, area)
	}
	return res
}

// @lc code=end

func TestContainerWithMostWater(t *testing.T) {
	// your test code here
	num := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Printf("maxArea(num): %v\n", maxArea(num))
}

/*
// @lcpr case=start
// [1,8,6,2,5,4,8,3,7]\n
// @lcpr case=end

// @lcpr case=start
// [1,1]\n
// @lcpr case=end

*/
