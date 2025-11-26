/*
 * @lc app=leetcode.cn id=14 lang=golang
 * @lcpr version=30304
 *
 * [14] 最长公共前缀
 */

package leetcode_solutions

import (
	"fmt"
	"testing"
)

// @lc code=start
func longestCommonPrefix(strs []string) string {
	m := len(strs)
	n := len(strs[0])
	for col := 0; col < n; col++ {
		for row := 1; row < m; row++ {
			thisStr, prevStr := strs[row], strs[row-1]
			if col >= len(thisStr) || col >= len(prevStr) || thisStr[col] != prevStr[col] {
				return strs[row][:col]
			}
		}
	}
	return strs[0]
}

// @lc code=end

func TestLongestCommonPrefix(t *testing.T) {
	// your test code here
	s := []string{"flower", "flow", "flight"}
	fmt.Printf("longestCommonPrefix(s): %v\n", longestCommonPrefix(s))
}

/*
// @lcpr case=start
// ["flower","flow","flight"]\n
// @lcpr case=end

// @lcpr case=start
// ["dog","racecar","car"]\n
// @lcpr case=end

*/
