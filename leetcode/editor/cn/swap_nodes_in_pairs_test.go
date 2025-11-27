/*
 * @lc app=leetcode.cn id=24 lang=golang
 * @lcpr version=30304
 *
 * [24] 两两交换链表中的节点
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
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	first := head
	second := head.Next
	other := head.Next.Next
	second.Next = first
	first.Next = swapPairs(other)
	return second
}

// @lc code=end

func TestSwapNodesInPairs(t *testing.T) {
	// your test code here
	list1 := CreateHead([]int{1, 2, 3, 4})
	res := swapPairs(list1)
	PrintList(res)
}

/*
// @lcpr case=start
// [1,2,3,4]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

*/
