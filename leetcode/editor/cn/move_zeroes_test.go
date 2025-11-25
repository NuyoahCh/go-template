/*
 * @lc app=leetcode.cn id=283 lang=golang
 * @lcpr version=30304
 *
 * [283] 移动零
 */

package leetcode_solutions

import (
	"fmt"
	"testing"
)

// @lc code=start
func moveZeroes(nums []int) {
	for p := moveElement(nums, 0); p < len(nums); p++ {
		nums[p] = 0
	}
	return
}

func moveElement(nums []int, val int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

// @lc code=end

func TestMoveZeroes(t *testing.T) {
	// your test code here
	num := []int{0, 1, 0, 3, 12}
	moveZeroes(num)
	fmt.Print(num)
}

/*
// @lcpr case=start
// [0,1,0,3,12]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n
// @lcpr case=end

*/
