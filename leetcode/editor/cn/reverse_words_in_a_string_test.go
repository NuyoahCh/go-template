/*
 * @lc app=leetcode.cn id=151 lang=golang
 * @lcpr version=30304
 *
 * [151] 反转字符串中的单词
 */

package leetcode_solutions

import (
	"fmt"
	"strings"
	"testing"
)

// @lc code=start
func reverseWords(s string) string {
	var sb strings.Builder
	// 先清洗一下数据，把多于的空格都删掉
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c != ' ' {
			// 单词中的字母/数字
			sb.WriteByte(c)
		} else if sb.Len() > 0 && sb.String()[sb.Len()-1] != ' ' {
			// 单词之间保留一个空格
			sb.WriteByte(' ')
		}
	}
	if sb.Len() == 0 {
		return ""
	}
	// 末尾如果有空格，清除之
	cleaned := sb.String()
	if cleaned[len(cleaned)-1] == ' ' {
		cleaned = cleaned[:len(cleaned)-1]
	}

	// 清洗之后的字符串
	chars := []byte(cleaned)
	n := len(chars)
	// 进行单词的翻转，先整体翻转
	reverse01(chars, 0, n-1)
	// 再把每个单词翻转
	for i := 0; i < n; {
		for j := i; j < n; j++ {
			if j+1 == n || chars[j+1] == ' ' {
				// chars[i..j] 是一个单词，翻转之
				reverse01(chars, i, j)
				// 把 i 置为下一个单词的首字母
				i = j + 2
				break
			}
		}
	}
	// 最后得到题目想要的结果
	return string(chars)
}

// 翻转 arr[i..j]
func reverse01(arr []byte, i, j int) {
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

// @lc code=end

func TestReverseWordsInAString(t *testing.T) {
	// your test code here
	s := "the sky is blue"
	fmt.Println(reverseWords(s))
}

/*
// @lcpr case=start
// "the sky is blue"\n
// @lcpr case=end

// @lcpr case=start
// "  hello world  "\n
// @lcpr case=end

// @lcpr case=start
// "a good   example"\n
// @lcpr case=end

*/
