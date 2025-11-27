/*
 * @lc app=leetcode.cn id=32 lang=golang
 * @lcpr version=30304
 *
 * [32] 最长有效括号
 */

package leetcode_solutions

import (
	"fmt"
	"testing"
)

// @lc code=start
func longestValidParentheses(s string) int {
	stk := []int{}
	dp := make([]int, len(s)+1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stk = append(stk, i)
			dp[i+1] = 0
		} else {
			if len(stk) > 0 {
				leftIndex := stk[len(stk)-1]
				stk = stk[:len(stk)-1]
				len := i - leftIndex + 1 + dp[leftIndex]
				dp[i+1] = len
			} else {
				dp[i+1] = 0
			}
		}
	}
	res := 0
	for i := 0; i < len(dp); i++ {
		res = max(res, dp[i])
	}
	return res
}

// @lc code=end

func TestLongestValidParentheses(t *testing.T) {
	// your test code here
	s := "(()"
	fmt.Printf("longestValidParentheses(s): %v\n", longestValidParentheses(s))
}

/*
// @lcpr case=start
// "(()"\n
// @lcpr case=end

// @lcpr case=start
// ")()())"\n
// @lcpr case=end

// @lcpr case=start
// ""\n
// @lcpr case=end

*/
