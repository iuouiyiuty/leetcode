package main

import (
	"math"
)

/*
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21
注意:

假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31,  2^31 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
*/

//func main() {
//	fmt.Println(reverse(123))
//	fmt.Println(reverse(-123))
//	fmt.Println(reverse(1 << 29))
//}

func reverse(x int) int {
	res := int64(0) // 64位防止溢出

	for x != 0 {
		res = res*int64(10) + int64(x%10) // 每次乘以10，加余数，这样就实现了倒序
		x = x / 10
	}
	// 判断是否溢出
	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}

	return int(res)
}
