/*
 * @lc app=leetcode.cn id=88 lang=golang
 * @lcpr version=30304
 *
 * [88] 合并两个有序数组
 */

package leetcode_solutions

import "testing"

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, p := m-1, n-1, len(nums1)-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[p] = nums1[i]
			i--
		} else {
			nums1[p] = nums2[j]
			j--
		}
		p--
	}
	for j >= 0 {
		nums1[p] = nums2[j]
		p--
		j--
	}
}

// @lc code=end

func TestMergeSortedArray(t *testing.T) {
	// your test code here
	num1 := []int{1, 2, 3, 0, 0, 0}
	num2 := []int{2, 5, 6}
	m := 3
	n := 2
	merge(num1, m, num2, n)
}

/*
// @lcpr case=start
// [1,2,3,0,0,0]\n3\n[2,5,6]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1]\n1\n[]\n0\n
// @lcpr case=end

// @lcpr case=start
// [0]\n0\n[1]\n1\n
// @lcpr case=end

*/
