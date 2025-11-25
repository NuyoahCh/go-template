/*
 * @lc app=leetcode.cn id=23 lang=golang
 * @lcpr version=30304
 *
 * [23] 合并 K 个升序链表
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
func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}
	return mergeKListsHelper(lists, 0, n-1)
}

func mergeKListsHelper(lists []*ListNode, start, end int) *ListNode {
	if start == end {
		return lists[start]
	}
	mid := start + (end - start) >> 1
	left := mergeKListsHelper(lists, start, mid)
	right := mergeKListsHelper(lists, mid+1, end)
	return mergeTwoList(left, right)
}

func mergeTwoList(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{-1, nil}
	p := dummy
	p1, p2 := list1, list2

	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		p = p.Next
	}
	if p1 == nil {
		p.Next = p2
	}
	if p2 == nil {
		p.Next = p1
	}
	return dummy.Next
	
}

// @lc code=end

func TestMergeKSortedLists(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [[1,4,5],[1,3,4],[2,6]]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [[]]\n
// @lcpr case=end

*/
