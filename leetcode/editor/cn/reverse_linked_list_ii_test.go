/*
 * @lc app=leetcode.cn id=92 lang=golang
 * @lcpr version=30304
 *
 * [92] 反转链表 II
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
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// base case
	if head == nil || head.Next == nil {
		return head
	}
	// base case
	if left == 1 {
		return reverseN(head, right)
	}
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

// 后驱节点
var successor *ListNode

// 反转部分链表
func reverseN(head *ListNode, n int) *ListNode {
	// base case
	if n == 1 {
		// 记录下一个节点
		successor = head.Next
		// 返回头节点
		return head
	}
	// 反转核心执行
	last := reverseN(head.Next, n-1)
	head.Next.Next = head
	// 反转后继节点
	head.Next = successor
	return last
}

// @lc code=end

func TestReverseLinkedListIi(t *testing.T) {
	// your test code here
	list1 := CreateHead([]int{1, 2, 3, 4, 5})
	left := 2
	right := 4
	res := reverseBetween(list1, left, right)
	PrintList(res)
}

/*
// @lcpr case=start
// [1,2,3,4,5]\n2\n4\n
// @lcpr case=end

// @lcpr case=start
// [5]\n1\n1\n
// @lcpr case=end

*/
