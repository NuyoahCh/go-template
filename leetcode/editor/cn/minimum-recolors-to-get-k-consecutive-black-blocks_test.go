package leetcode_solutions

import (
	"fmt"
	"math"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func minimumRecolors(blocks string, k int) int {
	ans := math.MaxInt
	sum := 0
	for i := range blocks {
		if blocks[i] == 'W' {
			sum++
		}
		left := i - k + 1
		if left < 0 {
			continue
		}
		ans = min(ans, sum)
		if blocks[left] == 'W' {
			sum--
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMinimumRecolorsToGetKConsecutiveBlackBlocks(t *testing.T) {
	// your test code here
	s := "WBBWWBBWBW"
	k := 7
	fmt.Println(minimumRecolors(s, k))
}
