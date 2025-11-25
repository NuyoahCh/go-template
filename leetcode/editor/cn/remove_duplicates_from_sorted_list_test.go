/*
 * @lc app=leetcode.cn id=83 lang=golang
 * @lcpr version=30304
 *
 * [83] 删除排序链表中的重复元素
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
func deleteDuplicates(head *ListNode) *ListNode {
	// base case
	if head == nil {
		return nil
	}
	// 初始化双指针
	slow, fast := head, head
	// 遍历到链表的结尾
	for fast != nil {
		// 不存在重复节点
		if fast.Val != slow.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}
	// 断开和后面重复元素的连接
	slow.Next = nil
	return head
}

// @lc code=end

func TestRemoveDuplicatesFromSortedList(t *testing.T) {
	// your test code here
	list1 := CreateHead([]int{1, 1, 2})
	res := deleteDuplicates(list1)
	PrintList(res)
}

/*
// @lcpr case=start
// [1,1,2]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,2,3,3]\n
// @lcpr case=end

*/
