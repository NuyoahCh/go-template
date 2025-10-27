package leetcode_solutions

import (
	"fmt"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	valToIndex := make(map[int]int)
	for i, num := range nums {
		need := target - num
		if j, ok := valToIndex[need]; ok {
			return []int{j, i}
		}
		valToIndex[num] = i
	}
	return nil
}

//leetcode submit region end(Prohibit modification and deletion)

func TestTwoSum(t *testing.T) {
	// your test code here
	arr := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(arr, target))
}
