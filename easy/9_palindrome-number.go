package main

import "fmt"

/*
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:

输入: 121
输出: true
示例 2:

输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3:

输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。
进阶:

你能不将整数转为字符串来解决这个问题吗？
*/

//func main() {
//	fmt.Println(isPalindrome(11))
//
//	fmt.Println(isPalindrome2(121))
//	fmt.Println(isPalindrome2(12345678987654321))
//	fmt.Println(isPalindrome2(11))
//	fmt.Println(isPalindrome2(-121))
//	fmt.Println(isPalindrome2(10))
//}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	str := fmt.Sprint(x)
	length := len(str)

	for i := range str {
		if str[i] != str[length-i-1] {
			return false
		}
	}
	return true
}

// 进阶，参考整数反转
func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	res := 0
	tmp := x
	for tmp > 0 {
		res = res*10 + tmp%10
		tmp = tmp / 10
	}

	return res == x
}
