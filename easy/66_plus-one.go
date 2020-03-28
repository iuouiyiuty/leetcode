package main

/*
66. 加一
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。

示例 1:
输入: [1,2,3]
输出: [1,2,4]
解释: 输入数组表示数字 123。

示例 2:
输入: [4,3,2,1]
输出: [4,3,2,2]
解释: 输入数组表示数字 4321。
*/

//func main() {
//	fmt.Println(plusOne([]int{1, 2, 3}))
//	fmt.Println(plusOne([]int{1, 2, 3, 4}))
//	fmt.Println(plusOne([]int{1, 9, 9}))
//	fmt.Println(plusOne([]int{1, 9}))
//	fmt.Println(plusOne([]int{9, 9}))
//}

func plusOne(digits []int) []int {
	// 一位，且小于9，直接加1返回
	if len(digits) == 1 && digits[0] < 9 {
		digits[0]++
		return digits
	}

	for i := len(digits) - 1; i >= 0; i-- {
		// 小于9直接加1
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		// 进位
		digits[i] = 0
		// 最后一位，则首位添加一个元素，为1
		if i-1 < 0 {
			nd := make([]int, 0, len(digits)+1)
			nd = append(nd, 1)
			nd = append(nd, digits...)
			return nd
		}
	}
	return digits
}
