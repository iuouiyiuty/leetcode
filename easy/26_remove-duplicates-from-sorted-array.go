package main

import "fmt"

/*
26. 删除排序数组中的重复项
给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

示例 1:
给定数组 nums = [1,1,2],
函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。
你不需要考虑数组中超出新长度后面的元素。

示例 2:
给定 nums = [0,0,1,1,1,2,2,3,3,4],
函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。
你不需要考虑数组中超出新长度后面的元素。
*/

func main() {
	arr := []int{1, 1, 2}
	fmt.Println(removeDuplicates(arr))
	fmt.Println(arr)

	arr2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(arr2))
	fmt.Println(arr2)

}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}

	validLen := len(nums)
	i := 1
	for {
		if i >= validLen {
			break
		}
		if nums[i] == nums[i-1] {
			// 全部向前移1位
			for j := i; j < validLen-1; j++ {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
			validLen -= 1
			i--
		}
		i++
	}

	return validLen
}

//func removeDuplicates2(nums []int) int {
//	index := 0
//	for i:=1; i < len(nums); i++ {
//		// 覆盖重复位，会修改原数组元素
//		if nums[i] != nums[index] {
//			index++
//			nums[index] = nums[i]
//		}
//	}
//	return index+1
//}
