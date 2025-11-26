/*
 * @lc app=leetcode.cn id=167 lang=golang
 * @lcpr version=30304
 *
 * [167] 两数之和 II - 输入有序数组
 */

package leetcode_solutions

import (
	"fmt"
	"testing"
)

// @lc code=start
func twoSum01(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left <= right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return []int{-1, -1}
}

// @lc code=end

func TestTwoSumIiInputArrayIsSorted(t *testing.T) {
	// your test code here
	num := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum01(num, target))
}

/*
// @lcpr case=start
// [2,7,11,15]\n9\n
// @lcpr case=end

// @lcpr case=start
// [2,3,4]\n6\n
// @lcpr case=end

// @lcpr case=start
// [-1,0]\n-1\n
// @lcpr case=end

*/
