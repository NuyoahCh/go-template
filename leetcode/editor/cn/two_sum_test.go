/*
 * @lc app=leetcode.cn id=1 lang=golang
 * @lcpr version=30304
 *
 * [1] 两数之和
 */

package leetcode_solutions

import (
	"fmt"
	"testing"
)

// @lc code=start
func twoSum(nums []int, target int) []int {
	valToIndex := make(map[int]int)
	for i, num := range nums {
		need := target - num
		if j, ok := valToIndex[need]; ok {
			return []int{j, i}
		}
		valToIndex[num] = i
	}
	return nil
}

// @lc code=end

func TestTwoSum(t *testing.T) {
	// your test code here
	num := []int{2, 7, 11, 15}
	target := 9
	fmt.Printf("twoSum(num): %v\n", twoSum(num, target))
}

/*
// @lcpr case=start
// [2,7,11,15]\n9\n
// @lcpr case=end

// @lcpr case=start
// [3,2,4]\n6\n
// @lcpr case=end

// @lcpr case=start
// [3,3]\n6\n
// @lcpr case=end

*/
