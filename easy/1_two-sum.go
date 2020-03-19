package main

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/

//func main() {
//	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
//	fmt.Println(twoSum([]int{2, 7, 11, 15}, 18))
//	fmt.Println(twoSum([]int{2, 7, 11, 15}, 19))
//	fmt.Println(twoSum([]int{3, 3}, 6))
//	fmt.Println(twoSum([]int{0, 4, 3, 0}, 0))
//}

func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}

	m := make(map[int]int, 0)
	// 遍历数组
	for k, v := range nums {
		// 判断结果集是否存在另一个符合要求的数
		if i, ok := m[target-v]; ok {
			// 存在则返回
			return []int{i, k}
		}
		// 不村组则放入结果集，继续遍历
		m[v] = k
	}

	return nil
}
