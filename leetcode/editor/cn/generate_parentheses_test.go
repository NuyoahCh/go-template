/*
 * @lc app=leetcode.cn id=22 lang=golang
 * @lcpr version=30304
 *
 * [22] 括号生成
 */

package leetcode_solutions

import (
	"fmt"
	"testing"
)

// @lc code=start
func generateParenthesis(n int) []string {
	// 回溯路径
	var track string
	// 存储结果
	var res []string

	// 回溯函数
	var backtrack func(left, right int)

	// 执行回溯操作
	backtrack = func(left, right int) {
		// 不符合要求的情况
		if right < left {
			return
		}
		if left < 0 || right < 0 {
			return
		}
		// 符合就进行存储
		if left == 0 && right == 0 {
			res = append(res, track)
			return
		}

		track += "("
		backtrack(left-1, right)
		track = track[:len(track)-1]

		track += ")"
		backtrack(left, right-1)
		track = track[:len(track)-1]
	}
	// 初始化回溯
	backtrack(n, n)
	return res
}

// @lc code=end

func TestGenerateParentheses(t *testing.T) {
	// your test code here
	n := 3
	fmt.Printf("generateParenthesis(n): %v\n", generateParenthesis(n))
}

/*
// @lcpr case=start
// 3\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/
