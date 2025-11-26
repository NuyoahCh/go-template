/*
 * @lc app=leetcode.cn id=76 lang=golang
 * @lcpr version=30304
 *
 * [76] 最小覆盖子串
 */

package leetcode_solutions

import (
	"fmt"
	"math"
	"testing"
)

// @lc code=start
func minWindow(s string, t string) string {
	// 需要窗口，实际窗口
	need, window := make(map[byte]int), make(map[byte]int)
	// 映射记录所需元素的个数
	for i := range t {
		need[t[i]]++
	}
	// 左右边界
	left, right := 0, 0
	// 有效值
	valid := 0
	// 记录起始值，和子串长度
	start, length := 0, math.MaxInt32
	for right < len(s) {
		c := s[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for valid == len(need) {
			if right-left < length {
				start = left
				length = right - left
			}
			d := s[left]
			left++

			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	if length == math.MaxInt32 {
		return ""
	}
	return s[start : start+length]
}

// @lc code=end

func TestMinimumWindowSubstring(t *testing.T) {
	// your test code here
	s1 := "ADOBECODEBANC"
	t1 := "ABC"
	fmt.Printf("minWindow(s1, t1): %v\n", minWindow(s1, t1))
}

/*
// @lcpr case=start
// "ADOBECODEBANC"\n"ABC"\n
// @lcpr case=end

// @lcpr case=start
// "a"\n"a"\n
// @lcpr case=end

// @lcpr case=start
// "a"\n"aa"\n
// @lcpr case=end

*/
