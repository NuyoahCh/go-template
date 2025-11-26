/*
 * @lc app=leetcode.cn id=3 lang=golang
 * @lcpr version=30304
 *
 * [3] 无重复字符的最长子串
 */

package leetcode_solutions

import (
	"fmt"
	"testing"
)

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	// 存储字符
	cnt := [128]int{}
	left := 0
	ans := 0
	for right, c := range s {
		cnt[c]++
		for cnt[c] > 1 {
			cnt[s[left]]--
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

// @lc code=end

func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	// your test code here
	s := "abcabcbb"
	fmt.Printf("lengthOfLongestSubstring(s): %v\n", lengthOfLongestSubstring(s))
}

/*
// @lcpr case=start
// "abcabcbb"\n
// @lcpr case=end

// @lcpr case=start
// "bbbbb"\n
// @lcpr case=end

// @lcpr case=start
// "pwwkew"\n
// @lcpr case=end

*/
