/*
 * @lc app=leetcode.cn id=80 lang=golang
 * @lcpr version=30304
 *
 * [80] 删除有序数组中的重复项 II
 */

package leetcode_solutions

import (
	"fmt"
	"testing"
)

// @lc code=start
func removeDuplicates(nums []int) int {
	// base case
	if len(nums) == 0 {
		return 0
	}
	// 快慢指针
	slow, fast := 0, 0
	// 计数
	count := 0
	for fast < len(nums) {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		} else if slow < fast && count < 2 {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
		count++
		if fast < len(nums) && nums[fast] != nums[fast-1] {
			count = 0
		}
	}
	return slow + 1
}

// @lc code=end

func TestRemoveDuplicatesFromSortedArrayIi(t *testing.T) {
	// your test code here
	num := []int{1, 1, 1, 2, 2, 3}
	fmt.Print(removeDuplicates(num))
}

/*
// @lcpr case=start
// [1,1,1,2,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [0,0,1,1,1,1,2,3,3]\n
// @lcpr case=end

*/
