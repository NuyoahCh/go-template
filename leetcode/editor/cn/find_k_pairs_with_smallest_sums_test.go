/*
 * @lc app=leetcode.cn id=373 lang=golang
 * @lcpr version=30304
 *
 * [373] 查找和最小的 K 对数字
 */

package leetcode_solutions

import (
	"container/heap"
	"fmt"
	"testing"
)

// @lc code=start
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	// 存储三元组 (num1[i], nums2[i], i)
	// i 记录 nums2 元素的索引位置，用于生成下一个节点
	pq := &PriorityQueue{}
	heap.Init(pq)

	// 按照 23 题的逻辑初始化优先级队列
	for i := 0; i < len(nums1); i++ {
		heap.Push(pq, []int{nums1[i], nums2[0], 0})
	}

	res := [][]int{}
	// 执行合并多个有序链表的逻辑
	for pq.Len() > 0 && k > 0 {
		cur := heap.Pop(pq).([]int)
		k--
		// 链表中的下一个节点加入优先级队列
		nextIndex := cur[2] + 1
		if nextIndex < len(nums2) {
			heap.Push(pq, []int{cur[0], nums2[nextIndex], nextIndex})
		}

		res = append(res, []int{cur[0], cur[1]})
	}
	return res
}

// PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue [][]int

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// 按照数对的元素和升序排序
	return (pq[i][0] + pq[i][1]) < (pq[j][0] + pq[j][1])
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.([]int))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// @lc code=end

func TestFindKPairsWithSmallestSums(t *testing.T) {
	// your test code here
	nums1 := []int{1, 7, 11}
	nums2 := []int{2, 4, 6}
	k := 3
	fmt.Println(kSmallestPairs(nums1, nums2, k))
}

/*
// @lcpr case=start
// [1,7,11]\n[2,4,6]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1,1,2]\n[1,2,3]\n2\n
// @lcpr case=end

*/
