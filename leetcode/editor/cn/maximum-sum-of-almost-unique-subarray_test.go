package leetcode_solutions

import "testing"

// leetcode submit region begin(Prohibit modification and deletion)
func maxSum(nums []int, m int, k int) (ans int64) {
	s := int64(0)
	cnt := map[int]int{}
	for i, x := range nums {
		s += int64(x)
		cnt[x]++

		left := i - k + 1
		if left < 0 {
			continue
		}

		if len(cnt) >= m {
			ans = max(ans, s)
		}

		out := nums[left]
		s -= int64(out)
		cnt[out]--
		if cnt[out] == 0 {
			delete(cnt, out)
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMaximumSumOfAlmostUniqueSubarray(t *testing.T) {
	// your test code here

}
