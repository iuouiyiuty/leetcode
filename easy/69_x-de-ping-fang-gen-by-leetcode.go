package main

import "fmt"

/*
69. x 的平方根
实现 int sqrt(int x) 函数。

计算并返回 x 的平方根，其中 x 是非负整数。

由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

示例 1:
输入: 4
输出: 2

示例 2:
输入: 8
输出: 2
说明: 8 的平方根是 2.82842...,
     由于返回类型是整数，小数部分将被舍去。
*/

func main() {
	fmt.Println(mySqrt(8))
}

// 2分法
func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	left, right := 2, x/2

	for left <= right {
		m := (right + left) / 2 // 中数
		n := m * m              // 中数平方
		if n > x {              // 大于x则右边界-1
			right = m - 1
		} else if n < x { // 小于x则左边界+1
			left = m + 1
		} else { // 整除则正好
			return m
		}
	}
	return right
}
