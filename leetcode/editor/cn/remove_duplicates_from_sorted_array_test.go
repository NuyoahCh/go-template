/*
 * @lc app=leetcode.cn id=26 lang=golang
 * @lcpr version=30304
 *
 * [26] 删除有序数组中的重复项
 */

package leetcode_solutions

import (
	"fmt"
	"testing"
)

// @lc code=start
func removeDuplicates(nums []int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[slow] != nums[fast] {
			// 保证第一个元素不会丢失
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	// 返回长度
	return slow + 1
}

// @lc code=end

func TestRemoveDuplicatesFromSortedArray(t *testing.T) {
	// your test code here
	num1 := []int{1, 1, 2}
	fmt.Print(removeDuplicates(num1))
}

/*
// @lcpr case=start
// [1,1,2]\n
// @lcpr case=end

// @lcpr case=start
// [0,0,1,1,1,2,2,3,3,4]\n
// @lcpr case=end

*/
