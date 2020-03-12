package main

import "fmt"

/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:
所有输入只包含小写字母 a-z 。
*/

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "floght"}))
	fmt.Println(longestCommonPrefix([]string{"flow", "flow", "flow"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 使用第一位做前缀
	prefix := strs[0]
	// 过滤掉prefix
	strs = strs[1:]

	for _, str := range strs {
		// 最大匹配的下标
		sameIdx := 0

		// 遍历取短字符长度
		min := len(prefix)
		if len(prefix) > len(str) {
			min = len(str)
		}
		for sameIdx = 0; sameIdx < min; sameIdx++ {
			// 不相同跳出匹配
			if prefix[sameIdx] != str[sameIdx] {
				break
			}
		}
		if sameIdx == 0 {
			return ""
		}
		// 更新prefix
		prefix = prefix[:sameIdx]
	}

	return prefix
}
