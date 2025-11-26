/*
 * @lc app=leetcode.cn id=5 lang=golang
 * @lcpr version=30304
 *
 * [5] 最长回文子串
 */

package leetcode_solutions

import (
	"fmt"
	"testing"
)

// @lc code=start
func longestPalindrome(s string) string {
	// 空字符串
	res := ""
	// 遍历查找
	for i := 0; i < len(s); i++ {
		// 考虑到奇偶性
		s1 := palindrome(s, i, i)
		s2 := palindrome(s, i, i+1)
		// 取最大值
		if len(res) < len(s1) {
			res = s1
		}
		if len(res) < len(s2) {
			res = s2
		}
	}
	// 返回最长的回文串
	return res
}

// 中心扩散法，从中间到两边进行扩散
func palindrome(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	// 去到回文的所有数值
	return s[l+1 : r]
}

// @lc code=end

func TestLongestPalindromicSubstring(t *testing.T) {
	// your test code here
	s := "babad"
	res := longestPalindrome(s)
	fmt.Print(res)
}

/*
// @lcpr case=start
// "babad"\n
// @lcpr case=end

// @lcpr case=start
// "cbbd"\n
// @lcpr case=end

*/
