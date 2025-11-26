/*
 * @lc app=leetcode.cn id=344 lang=golang
 * @lcpr version=30304
 *
 * [344] 反转字符串
 */

package leetcode_solutions

import "testing"

// @lc code=start
func reverseString(s []byte) {
	left, right := 0, len(s)-1
	// 还是采用闭区间的方式
	for left <= right {
		temp := s[left]
		s[left] = s[right]
		s[right] = temp
		left++
		right--
	}
}

// @lc code=end

func TestReverseString(t *testing.T) {
	// your test code here
	s1 := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s1)
}

/*
// @lcpr case=start
// ["h","e","l","l","o"]\n
// @lcpr case=end

// @lcpr case=start
// ["H","a","n","n","a","h"]\n
// @lcpr case=end

*/
