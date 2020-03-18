package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
38. 外观数列
「外观数列」是一个整数序列，从数字 1 开始，序列中的每一项都是对前一项的描述。前五项如下：

1.     1
2.     11
3.     21
4.     1211
5.     111221
1 被读作  "one 1"  ("一个一") , 即 11。
11 被读作 "two 1s" ("两个一"）, 即 21。
21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。

给定一个正整数 n（1 ≤ n ≤ 30），输出外观数列的第 n 项。

注意：整数序列中的每一项将表示为一个字符串。

示例 1:
输入: 1
输出: "1"
解释：这是一个基本样例。

示例 2:
输入: 4
输出: "1211"
解释：当 n = 3 时，序列是 "21"，其中我们有 "2" 和 "1" 两组，"2" 可以读作 "12"，
也就是出现频次 = 1 而 值 = 2；类似 "1" 可以读作 "11"。所以答案是 "12" 和 "11" 组合在一起，也就是 "1211"。
*/

//func main() {
//	for i := 1; i <= 30; i++ {
//		fmt.Printf("n: %d,\tres: %s\n", i, CountAndSay2(i))
//	}
//}

func CountAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	return generate(2, n, "1")
}

// 递归方式
func generate(start, n int, preStr string) string {
	// 终止条件
	if start > n {
		return preStr
	}

	curStr := "" // 当前元素
	times := 0   // 当前元素个数
	newStr := "" // start的输出
	for _, v := range preStr {
		// 元素相等，则次数++
		if string(v) == curStr {
			times++
			continue
		}
		// 当前元素为空，直接赋值，次数为1，进入下次循环
		if curStr == "" {
			curStr = string(v)
			times++
			continue
		}
		// 不相等，写入newStr
		// 次数不多没必要用strings.Builder
		newStr += fmt.Sprintf("%d%s", times, curStr)

		curStr = string(v)
		times = 1
	}
	//  把最后一次循环的结果写入newStr
	newStr += fmt.Sprintf("%d%s", times, curStr)

	return generate(start+1, n, newStr)
}

var arr = [30]string{"1"}

// 使用多层循环实现
// 这里使用strings.Builder依旧性能很差，比递归差20多倍，大头在底层slice扩容和其内存迁移
func CountAndSay2(n int) string {
	if n == 1 {
		return arr[0]
	}

	str := strings.Builder{}
	for i := 1; i < n; i++ {
		str.Reset()
		// 遍历arr[i-1]的元素，得出arr[i]的值
		for j := 0; j < len(arr[i-1]); j++ {
			// 当前元素
			cur := arr[i-1][0]
			// 次数
			times := 1
			// 循环匹配判断
			for k := j + 1; k < len(arr[i-1]) && cur == arr[i-1][k]; k++ {
				times++
			}
			str.WriteString(strconv.Itoa(times))
			str.WriteByte(cur)
		}
		// 把值写入
		arr[i] = str.String()
	}

	return generate(2, n, "1")
}
