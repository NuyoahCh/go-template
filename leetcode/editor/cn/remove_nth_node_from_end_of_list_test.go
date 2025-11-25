/*
 * @lc app=leetcode.cn id=19 lang=golang
 * @lcpr version=30304
 *
 * [19] 删除链表的倒数第 N 个结点
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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{-1, head}
	x := findFromEnd(dummy, n+1)
	x.Next = x.Next.Next
	return dummy.Next
}

func findFromEnd(head *ListNode, k int) *ListNode {
	p1 := head
	for i := 0; i < k; i++ {
		p1 = p1.Next
	}
	p2 := head
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p2
}

// @lc code=end

func TestRemoveNthNodeFromEndOfList(t *testing.T) {
	// your test code here
	list1 := CreateHead([]int{1, 2, 3, 4, 5})
	n := 2
	res := removeNthFromEnd(list1, n)
	PrintList(res)

}

/*
// @lcpr case=start
// [1,2,3,4,5]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1]\n1\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n1\n
// @lcpr case=end

*/
