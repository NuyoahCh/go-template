/*
 * @lc app=leetcode.cn id=23 lang=golang
 * @lcpr version=30304
 *
 * [23] 合并 K 个升序链表
 */

package leetcode_solutions

import (
	"reflect"
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
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return mergeKListsHelper(lists, 0, len(lists)-1)
}

func mergeKListsHelper(lists []*ListNode, start, end int) *ListNode {
	if start == end {
		return lists[start]
	}
	mid := start + (end-start)>>1
	left := mergeKListsHelper(lists, start, mid)
	right := mergeKListsHelper(lists, mid+1, end)
	return mergeTwoList(left, right)
}

func mergeTwoList(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{-1, nil}
	p := dummy
	p1, p2 := list1, list2
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		p = p.Next
	}
	if p1 == nil {
		p.Next = p2
	}
	if p2 == nil {
		p.Next = p1
	}
	return dummy.Next
}

// @lc code=end

// 辅助函数：将链表转为切片（用于断言测试结果）
func listToArray(head *ListNode) []int {
	res := make([]int, 0)
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

// TestMergeKSortedLists 测试合并K个升序链表
func TestMergeKSortedLists(t *testing.T) {
	// 定义测试用例：输入（链表数组的切片表示）、预期输出
	testCases := []struct {
		name     string  // 测试用例名称
		input    [][]int // 输入的K个升序链表（用切片模拟）
		expected []int   // 预期合并后的结果
	}{
		{
			name:     "示例1：3个非空链表",
			input:    [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
			expected: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			name:     "边界场景1：空链表数组",
			input:    [][]int{},
			expected: []int{},
		},
		{
			name:     "边界场景2：只有1个空链表",
			input:    [][]int{{}},
			expected: []int{},
		},
		{
			name:     "边界场景3：只有1个非空链表",
			input:    [][]int{{3, 6, 8}},
			expected: []int{3, 6, 8},
		},
		{
			name:     "特殊场景1：部分链表为空",
			input:    [][]int{{1, 3}, {}, {2, 5, 7}},
			expected: []int{1, 2, 3, 5, 7},
		},
		{
			name:     "特殊场景2：所有链表都为空",
			input:    [][]int{{}, {}, {}},
			expected: []int{},
		},
		{
			name:     "特殊场景3：链表长度差异大",
			input:    [][]int{{2}, {1, 3, 7}, {4, 5, 6, 9, 10}},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 9, 10},
		},
		{
			name:     "特殊场景4：包含负数",
			input:    [][]int{{-5, -3, -1}, {}, {-4, -2, 0}},
			expected: []int{-5, -4, -3, -2, -1, 0},
		},
		{
			name:     "特殊场景5：重复元素多",
			input:    [][]int{{2, 2, 2}, {2, 2}, {2}},
			expected: []int{2, 2, 2, 2, 2, 2},
		},
	}

	// 遍历所有测试用例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. 构造输入的链表数组
			lists := make([]*ListNode, len(tc.input))
			for i, nums := range tc.input {
				lists[i] = CreateHead(nums)
			}

			// 2. 调用待测试函数
			result := mergeKLists(lists)

			// 3. 转换结果为切片，与预期对比
			resultArr := listToArray(result)
			if !reflect.DeepEqual(resultArr, tc.expected) {
				t.Errorf("测试失败：\n输入：%v\n预期：%v\n实际：%v", tc.input, tc.expected, resultArr)
			}
		})
	}
}

/*
// @lcpr case=start
// [[1,4,5],[1,3,4],[2,6]]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [[]]\n
// @lcpr case=end

*/
