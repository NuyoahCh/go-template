/*
 * @lc app=leetcode.cn id=2 lang=golang
 * @lcpr version=30304
 *
 * [2] 两数相加
 */

package leetcode_solutions

import (
	"testing"
)

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 初始化虚拟头节点
	dummy := &ListNode{-1, nil}
	p := dummy
	p1, p2 := l1, l2
	// 进位值
	carry := 0
	// 核心执行逻辑
	for p1 != nil || p2 != nil || carry > 0 {
		// 存储进位值
		val := carry
		if p1 != nil {
			val += p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			val += p2.Val
			p2 = p2.Next
		}
		carry = val / 10
		val = val % 10
		p.Next = &ListNode{val, nil}
		p = p.Next
	}
	return dummy.Next
}

// @lc code=end

func TestAddTwoNumbers(t *testing.T) {
	// your test code here
	l1 := CreateHead([]int{2, 4, 3})
	l2 := CreateHead([]int{5, 6, 4})
	PrintList(addTwoNumbers(l1, l2))
}

/*
// @lcpr case=start
// [2,4,3]\n[5,6,4]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n[0]\n
// @lcpr case=end

// @lcpr case=start
// [9,9,9,9,9,9,9]\n[9,9,9,9]\n
// @lcpr case=end

*/
