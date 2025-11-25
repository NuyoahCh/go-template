/*
 * @lc app=leetcode.cn id=876 lang=golang
 * @lcpr version=30304
 *
 * [876] 链表的中间结点
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
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// @lc code=end

func TestMiddleOfTheLinkedList(t *testing.T) {
	// your test code here
	list1 := CreateHead([]int{1, 2, 3, 4, 5})
	PrintList(middleNode(list1))
}

/*
// @lcpr case=start
// [1,2,3,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5,6]\n
// @lcpr case=end

*/
