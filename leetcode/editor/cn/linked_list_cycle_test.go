/*
 * @lc app=leetcode.cn id=141 lang=golang
 * @lcpr version=30304
 *
 * [141] 环形链表
 */

package leetcode_solutions

import (
	"fmt"
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
func hasCycle(head *ListNode) bool {
	// 快慢指针
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		// 存在环
		if slow == fast {
			return true
		}
	}
	return false
}

// @lc code=end

func TestLinkedListCycle(t *testing.T) {
	// your test code here
	list1 := CreateHead([]int{3, 2, 0, -4})
	res := hasCycle(list1)
	fmt.Print(res)
}

/*
// @lcpr case=start
// [3,2,0,-4]\n1\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n0\n
// @lcpr case=end

// @lcpr case=start
// [1]\n-1\n
// @lcpr case=end

*/
