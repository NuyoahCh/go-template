/*
 * @lc app=leetcode.cn id=206 lang=golang
 * @lcpr version=30304
 *
 * [206] 反转链表
 */

package leetcode_solutions

import "testing"

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	// base case
	if head == nil || head.Next == nil {
		return head
	}
	// 反转的核心逻辑
	last := reverseList(head.Next)
	head.Next.Next = head
	// 避免成环
	head.Next = nil
	return last
}

// @lc code=end

func TestReverseLinkedList(t *testing.T) {
	// your test code here
	list1 := CreateHead([]int{1, 2, 3, 4, 5})
	res := reverseList(list1)
	PrintList(res)
}

/*
// @lcpr case=start
// [1,2,3,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/
