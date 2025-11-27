/*
 * @lc app=leetcode.cn id=33 lang=golang
 * @lcpr version=30304
 *
 * [33] 搜索旋转排序数组
 */

package leetcode_solutions

import (
	"fmt"
	"testing"
)

// @lc code=start
func search(nums []int, target int) int {
    // 左右都闭的搜索区间
    left, right := 0, len(nums) - 1
    // 因为是闭区间，所以结束条件为 left > right
    for left <= right {
        mid := left + (right - left) / 2
        // 首先检查 nums[mid]，是否找到 target
        if nums[mid] == target {
            return mid
        }
        if nums[mid] >= nums[left] {
            // mid 落在断崖左边，此时 nums[left..mid] 有序
            if target >= nums[left] && target < nums[mid] {
                // target 落在 [left..mid-1] 中
                right = mid - 1
            } else {
                // target 落在 [mid+1..right] 中
                left = mid + 1
            }
        } else {
            // mid 落在断崖右边，此时 nums[mid..right] 有序
            if target <= nums[right] && target > nums[mid] {
                // target 落在 [mid+1..right] 中
                left = mid + 1
            } else {
                // target 落在 [left..mid-1] 中
                right = mid - 1
            }
        }
    }
    // while 结束还没找到，说明 target 不存在
    return -1
}

// @lc code=end

func TestSearchInRotatedSortedArray(t *testing.T) {
	// your test code here
	num := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	fmt.Printf("search(num, target): %v\n", search(num, target))
}

/*
// @lcpr case=start
// [4,5,6,7,0,1,2]\n0\n
// @lcpr case=end

// @lcpr case=start
// [4,5,6,7,0,1,2]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1]\n0\n
// @lcpr case=end

*/
