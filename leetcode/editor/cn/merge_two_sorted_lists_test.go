/*
 * @lc app=leetcode.cn id=21 lang=golang
 * @lcpr version=30008
 *
 * [21] 合并两个有序链表
 */

package leetcode_solutions

import (
	"testing"
)

// @lc code=start

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 存在合并产生新的链表，就使用虚拟头节点
	dummy := &ListNode{-1, nil}
	// 初始化
	p := dummy
	p1, p2 := l1, l2
	// 遍历两个链表节点
	for p1 != nil && p2 != nil {
		// 升序执行
		if p1.Val < p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		p = p.Next
	}
	// 存在长短不一致的情况
	if p1 == nil {
		p.Next = p2
	}
	if p2 == nil {
		p.Next = p1
	}
	return dummy.Next
}

// @lc code=end

func TestMergeTwoSortedLists(t *testing.T) {
	// Add test cases here
	l1 := CreateHead([]int{1, 2, 4})
	l2 := CreateHead([]int{1, 3, 4})

	result := mergeTwoLists(l1, l2)
	PrintList(result)
}

/*
// @lcpr case=start
// [1,2,4]\n[1,3,4]\n
// @lcpr case=end

// @lcpr case=start
// []\n[]\n
// @lcpr case=end

// @lcpr case=start
// []\n[0]\n
// @lcpr case=end

*/
