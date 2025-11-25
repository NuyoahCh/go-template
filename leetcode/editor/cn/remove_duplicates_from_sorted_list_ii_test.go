/*
 * @lc app=leetcode.cn id=82 lang=golang
 * @lcpr version=30304
 *
 * [82] 删除排序链表中的重复元素 II
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
func deleteDuplicates01(head *ListNode) *ListNode {
	dummyDup, dummyUniq := &ListNode{Val: 101}, &ListNode{Val: 101}
	pDup, pUniq := dummyDup, dummyUniq
	p := head
	for p != nil {
		if (p.Next != nil && p.Val == p.Next.Val) || (p.Val == pDup.Val) {
			pDup.Next = p
			pDup = pDup.Next
		} else {
			pUniq.Next = p
			pUniq = pUniq.Next
		}
		p = p.Next
	}
	pDup.Next = nil
	pUniq.Next = nil
	return dummyUniq.Next
}

// @lc code=end

func TestRemoveDuplicatesFromSortedListIi(t *testing.T) {
	// your test code here
	list1 := CreateHead([]int{1, 2, 3, 3, 4, 4, 5})
	res := deleteDuplicates(list1)
	PrintList(res)
}

/*
// @lcpr case=start
// [1,2,3,3,4,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,1,2,3]\n
// @lcpr case=end

*/
