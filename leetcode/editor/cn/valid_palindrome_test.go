/*
 * @lc app=leetcode.cn id=125 lang=golang
 * @lcpr version=30304
 *
 * [125] 验证回文串
 */

package leetcode_solutions

import (
	"fmt"
	"strings"
	"testing"
	"unicode"
)

// @lc code=start
func isPalindrome(s string) bool {
	var sb strings.Builder
	for i := 0; i < len(s); i++ {
		c := s[i]
		if unicode.IsLetter(rune(c)) || unicode.IsDigit(rune(c)) {
			sb.WriteByte(byte(unicode.ToLower(rune(c))))
		}
	}
	filtered := sb.String()
	left, right := 0, len(filtered)-1
	for left < right {
		if filtered[left] != filtered[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// @lc code=end

func TestValidPalindrome(t *testing.T) {
	// your test code here
	s := "A man, a plan, a canal: Panama"
	res := isPalindrome(s)
	fmt.Print(res)
}

/*
// @lcpr case=start
// "A man, a plan, a canal: Panama"\n
// @lcpr case=end

// @lcpr case=start
// "race a car"\n
// @lcpr case=end

// @lcpr case=start
// " "\n
// @lcpr case=end

*/
