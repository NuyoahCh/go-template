/*
 * @lc app=leetcode.cn id=445 lang=golang
 * @lcpr version=30304
 *
 * [445] 两数相加 II
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
func addTwoNumbersII(l1 *ListNode, l2 *ListNode) *ListNode {
	// 把链表元素转入栈中
	var stk1 []int
	for l1 != nil {
		stk1 = append(stk1, l1.Val)
		l1 = l1.Next
	}
	var stk2 []int
	for l2 != nil {
		stk2 = append(stk2, l2.Val)
		l2 = l2.Next
	}

	// 接下来基本上是复用我在第 2 题的代码逻辑
	// 注意新节点要直接插入到 dummy 后面

	// 虚拟头结点（构建新链表时的常用技巧）
	dummy := &ListNode{-1, nil}

	// 记录进位
	carry := 0
	// 开始执行加法，两条链表走完且没有进位时才能结束循环
	for len(stk1) > 0 || len(stk2) > 0 || carry > 0 {
		// 先加上上次的进位
		val := carry
		if len(stk1) > 0 {
			val += stk1[len(stk1)-1]
			stk1 = stk1[:len(stk1)-1]
		}
		if len(stk2) > 0 {
			val += stk2[len(stk2)-1]
			stk2 = stk2[:len(stk2)-1]
		}
		// 处理进位情况
		carry = val / 10
		val = val % 10
		// 构建新节点，直接接在 dummy 后面
		newNode := &ListNode{Val: val, Next: dummy.Next}
		dummy.Next = newNode
	}
	// 返回结果链表的头结点（去除虚拟头结点）
	return dummy.Next
}

// @lc code=end

func TestAddTwoNumbersIi(t *testing.T) {
	// your test code here
	l1 := CreateHead([]int{7, 2, 4, 3})
	l2 := CreateHead([]int{5, 6, 4})
	res := addTwoNumbersII(l1, l2)
	PrintList(res)
}

/*
// @lcpr case=start
// [7,2,4,3]\n[5,6,4]\n
// @lcpr case=end

// @lcpr case=start
// [2,4,3]\n[5,6,4]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n[0]\n
// @lcpr case=end

*/
