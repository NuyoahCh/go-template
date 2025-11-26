/*
 * @lc app=leetcode.cn id=15 lang=golang
 * @lcpr version=30304
 *
 * [15] 三数之和
 */

package leetcode_solutions

import (
	"fmt"
	"sort"
	"testing"
)

// @lc code=start
func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)
	for first := 0; first < n-2; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := n - 1
		target := -nums[first]
		for second := first + 1; second < n; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}

// @lc code=end

func TestThreeSum(t *testing.T) {
	// your test code here
	num := []int{-1, 0, 1, 2, -1, -4}
	fmt.Printf("threeSum(num): %v\n", threeSum(num))
}

/*
// @lcpr case=start
// [-1,0,1,2,-1,-4]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,1]\n
// @lcpr case=end

// @lcpr case=start
// [0,0,0]\n
// @lcpr case=end

*/
