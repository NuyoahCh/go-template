/*
 * @lc app=leetcode.cn id=25 lang=golang
 * @lcpr version=30304
 *
 * [25] K 个一组翻转链表
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
func reverseKGroup(head *ListNode, k int) *ListNode {
	// base case
	if head == nil || head.Next == nil {
		return head
	}
	// base case
	if !hasKNode(head, k) {
		return head
	}

	newHead := reverseKN(head, k)
	head.Next = reverseKGroup(successorK, k)

	return newHead

}

// 反转后继节点
var successorK *ListNode

// 满足一组节点的要求
func hasKNode(head *ListNode, k int) bool {
	for i := 0; i < k; i++ {
		if head == nil {
			return false
		}
		head = head.Next
	}
	return true
}

// 反转部分链表
func reverseKN(head *ListNode, n int) *ListNode {
	// base case
	if n == 1 {
		// 记录下一个节点
		successorK = head.Next
		return head
	}
	// 核心反转逻辑
	last := reverseKN(head.Next, n-1)
	head.Next.Next = head
	head.Next = successorK
	return last
}

// @lc code=end

func TestReverseNodesInKGroup(t *testing.T) {
	// your test code here
	list1 := CreateHead([]int{1, 2, 3, 4, 5})
	k := 2
	res := reverseKGroup(list1, k)
	PrintList(res)
}

/*
// @lcpr case=start
// [1,2,3,4,5]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5]\n3\n
// @lcpr case=end

*/
