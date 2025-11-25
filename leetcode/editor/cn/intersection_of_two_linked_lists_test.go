/*
 * @lc app=leetcode.cn id=160 lang=golang
 * @lcpr version=30304
 *
 * [160] 相交链表
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
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB
	for p1 != p2 {
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}
		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p2
}

// @lc code=end

func TestIntersectionOfTwoLinkedLists(t *testing.T) {
	// your test code here
	list1 := CreateHead([]int{4, 1, 8, 4, 5})
	list2 := CreateHead([]int{5, 6, 1, 8, 4, 5})
	res := getIntersectionNode(list1, list2)
	PrintList(res)
}

/*
// @lcpr case=start
// 8\n[4,1,8,4,5]\n[5,6,1,8,4,5]\n2\n3\n
// @lcpr case=end

// @lcpr case=start
// 2\n[1,9,1,2,4]\n[3,2,4]\n3\n1\n
// @lcpr case=end

// @lcpr case=start
// 0\n[2,6,4]\n[1,5]\n3\n2\n
// @lcpr case=end

*/
