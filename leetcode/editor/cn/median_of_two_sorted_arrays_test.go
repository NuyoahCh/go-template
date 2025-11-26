/*
 * @lc app=leetcode.cn id=4 lang=golang
 * @lcpr version=30304
 *
 * [4] 寻找两个正序数组的中位数
 */

package leetcode_solutions

import (
	"fmt"
	"testing"
)

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	total := m + n
	res := make([]int, total)
	i, j, k := 0, 0, 0
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			res[k] = nums1[i]
			k++
			i++
		} else {
			res[k] = nums2[j]
			k++
			j++
		}
	}

	for i < m {
		res[k] = nums1[i]
		k++
		i++
	}

	for j < n {
		res[k] = nums2[j]
		k++
		j++
	}

	if total%2 == 1 {
		return float64(res[total/2])
	} else {
		mid1 := res[total/2-1]
		mid2 := res[total/2]
		return float64(mid1+mid2) / 2
	}
}

// @lc code=end

func TestMedianOfTwoSortedArrays(t *testing.T) {
	// your test code here
	num1 := []int{1, 3}
	num2 := []int{2}
	fmt.Printf("findMedianSortedArrays(num1, num2): %v\n", findMedianSortedArrays(num1, num2))
}

/*
// @lcpr case=start
// [1,3]\n[2]\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n[3,4]\n
// @lcpr case=end

*/
