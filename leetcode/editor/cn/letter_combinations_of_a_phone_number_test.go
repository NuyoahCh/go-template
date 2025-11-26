/*
 * @lc app=leetcode.cn id=17 lang=golang
 * @lcpr version=30304
 *
 * [17] 电话号码的字母组合
 */

package leetcode_solutions

import (
	"fmt"
	"testing"
)

// @lc code=start
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var mapping = []string{
		"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz",
	}

	var res []string

	// backtrack
	backtrack(digits, 0, "", &res, mapping)
	return res

}

func backtrack(digits string, start int, combination string, res *[]string, mapping []string) {
	if len(combination) == len(digits) {
		*res = append(*res, combination)
		return
	}

	digit := digits[start] - '0'
	for _, c := range mapping[digit] {
		backtrack(digits, start+1, combination+string(c), res, mapping)
	}
}

// @lc code=end

func TestLetterCombinationsOfAPhoneNumber(t *testing.T) {
	// your test code here
	s := "23"
	fmt.Printf("letterCombinations(s): %v\n", letterCombinations(s))
}

/*
// @lcpr case=start
// "23"\n
// @lcpr case=end

// @lcpr case=start
// "2"\n
// @lcpr case=end

*/
