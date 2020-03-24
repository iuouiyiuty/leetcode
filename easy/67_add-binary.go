package main

import (
	"fmt"
	"strconv"
)

/*
67. 二进制求和
给定两个二进制字符串，返回他们的和（用二进制表示）。

输入为非空字符串且只包含数字 1 和 0。

示例 1:
输入: a = "11", b = "1"
输出: "100"

示例 2:
输入: a = "1010", b = "1011"
输出: "10101"
*/

func main() {
	fmt.Println(addBinary("1010", "1011"))
}

func addBinary(a string, b string) string {
	aa, _ := strconv.ParseInt(a, 2, 64)
	bb, _ := strconv.ParseInt(b, 2, 64)
	return strconv.FormatInt(aa+bb, 2)
}
