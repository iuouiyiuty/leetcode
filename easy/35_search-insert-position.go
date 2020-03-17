package main

import "fmt"

/*
35. 搜索插入位置
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

你可以假设数组中无重复元素。

示例 1:
输入: [1,3,5,6], 5
输出: 2

示例 2:
输入: [1,3,5,6], 2
输出: 1

示例 3:
输入: [1,3,5,6], 7
输出: 4

示例 4:
输入: [1,3,5,6], 0
输出: 0
*/

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0))
	fmt.Println(searchInsert([]int{1, 3, 5, 6, 8}, 7))
	fmt.Println(searchInsert([]int{1, 3, 5, 6, 8}, 9))

	fmt.Println(searchInsert2([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert2([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert2([]int{1, 3, 5, 6}, 7))
	fmt.Println(searchInsert2([]int{1, 3, 5, 6}, 0))
	fmt.Println(searchInsert2([]int{1, 3, 5, 6, 8}, 7))
	fmt.Println(searchInsert2([]int{1, 3, 5, 6, 8}, 9))
}

// 直接遍历
func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}
	return len(nums)
}

// 2分法查找，适合数据量大的情况
func searchInsert2(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		if nums[0] >= target {
			return 0
		}
		return 1
	}

	front := 0
	end := len(nums) - 1

	for front <= end {
		mid := (front + end) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			end = mid - 1
		}
		if nums[mid] < target {
			front = mid + 1
		}
	}

	return front
}
