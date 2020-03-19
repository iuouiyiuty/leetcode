package main

/*
3. 无重复字符的最长子串
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:
输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例 2:
输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3:
输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/

//func main() {
//fmt.Println(lengthOfLongestSubstring("abcabcbb"))
//fmt.Println(lengthOfLongestSubstring("bbbbb"))
//fmt.Println(lengthOfLongestSubstring("pwwkew"))
//fmt.Println(lengthOfLongestSubstring("dvdf"))
//
//fmt.Println(lengthOfLongestSubstring2("abcabcbb"))
//fmt.Println(lengthOfLongestSubstring2("bbbbb"))
//fmt.Println(lengthOfLongestSubstring2("pwwkew"))
//fmt.Println(lengthOfLongestSubstring2("dvdf"))
//fmt.Println(lengthOfLongestSubstring2("abba"))
//}

// 暴力解决
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	max := 0     // 最长子串长度
	curStr := "" // 当前子串
	for _, v := range s {
		if curStr == "" {
			curStr += string(v)
			continue
		}
		flag := false // 是否出现重复字符
		// 可以用hashmap减小复杂度
		for i, cv := range curStr {
			if v == cv {
				// 比较当前子串长度
				if max < len(curStr) {
					max = len(curStr)
				}
				// 修改当前子串
				curStr = curStr[i+1:] + string(v)
				flag = true
				break
			}
		}
		if !flag {
			curStr += string(v)
		}
	}
	if max < len(curStr) {
		max = len(curStr)
	}
	return max
}

// 滑动窗口
func lengthOfLongestSubstring2(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	// 最大长度，左下标，右下标，[start,end)
	max, start, end := 0, 0, 0
	m := map[rune]int{}
	for i, v := range s {
		// 不重复
		idx, ok := m[v]
		if ok {
			// 判断大小
			if end-start > max {
				max = end - start
			}
			// 防止start左移
			if idx+1 > start {
				start = idx + 1
			}
		}
		// 赋值
		m[v] = i
		// 向右移动一位
		end++
	}
	if end-start > max {
		max = end - start
	}
	return max
}
