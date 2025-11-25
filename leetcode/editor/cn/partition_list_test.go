/*
 * @lc app=leetcode.cn id=86 lang=golang
 * @lcpr version=30304
 *
 * [86] 分隔链表
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
func partition(head *ListNode, x int) *ListNode {
	// 初始化
	dummy1 := &ListNode{-1, nil}
	dummy2 := &ListNode{-1, nil}
	p := head
	p1, p2 := dummy1, dummy2
	for p != nil {
		if p.Val < x {
			p1.Next = p
			p1 = p1.Next
		} else {
			p2.Next = p
			p2 = p2.Next
		}
		// 防止成环
		temp := p.Next
		p.Next = nil
		p = temp
	}
	p1.Next = dummy2.Next
	return dummy1.Next
}

// @lc code=end

func TestPartitionList(t *testing.T) {
	// your test code here
	list1 := CreateHead([]int{1, 4, 3, 2, 5, 2})
	x := 3
	res := partition(list1, x)
	PrintList(res)
}

/*
// @lcpr case=start
// [1,4,3,2,5,2]\n3\n
// @lcpr case=end

// @lcpr case=start
// [2,1]\n2\n
// @lcpr case=end

*/
