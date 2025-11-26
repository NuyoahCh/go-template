/*
 * @lc app=leetcode.cn id=234 lang=golang
 * @lcpr version=30304
 *
 * [234] 回文链表
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
func isPalindrome01(head *ListNode) bool {
	// 初始化快慢指针
	slow, fast := head, head
	// 查找链表中间节点
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 存在奇数的情况
	if fast != nil {
		slow = slow.Next
	}
	// 反转查找比对
	left, right := head, reverse(slow)
	// 遍历查找
	for right != nil {
		// 存在不是回文的情况
		if left.Val != right.Val {
			return false
		}
		// 继续前进
		left = left.Next
		right = right.Next
	}
	return true
}

// 反转链表
func reverse(head *ListNode) *ListNode {
	// base case
	if head == nil || head.Next == nil {
		return head
	}
	// 反转核心操作
	last := reverse(head.Next)
	head.Next.Next = head
	// 避免成环
	head.Next = nil
	return last
}

// @lc code=end

func TestPalindromeLinkedList(t *testing.T) {
	// your test code here
	last1 := CreateHead([]int{1, 2, 2, 1})
	res := isPalindrome01(last1)
	fmt.Println(res)
}

/*
// @lcpr case=start
// [1,2,2,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n
// @lcpr case=end

*/
