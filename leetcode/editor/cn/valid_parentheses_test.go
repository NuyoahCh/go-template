/*
 * @lc app=leetcode.cn id=20 lang=golang
 * @lcpr version=30304
 *
 * [20] 有效的括号
 */

package leetcode_solutions

import (
	"fmt"
	"testing"
)

// @lc code=start
func isValid(s string) bool {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if len(stack) != 0 && isPair(stack[len(stack)-1], s[i]) {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

func isPair(a, b byte) bool {
	if a == '(' && b == ')' {
		return true
	}
	if a == '[' && b == ']' {
		return true
	}
	if a == '{' && b == '}' {
		return true
	}
	return false
}

// @lc code=end

func TestValidParentheses(t *testing.T) {
	// your test code here
	s := "()[]{}"
	fmt.Printf("isValid(s): %v\n", isValid(s))
}

/*
// @lcpr case=start
// "()"\n
// @lcpr case=end

// @lcpr case=start
// "()[]{}"\n
// @lcpr case=end

// @lcpr case=start
// "(]"\n
// @lcpr case=end

// @lcpr case=start
// "([])"\n
// @lcpr case=end

// @lcpr case=start
// "([)]"\n
// @lcpr case=end

*/
