package main

/*
88. 合并两个有序数组
给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 num1 成为一个有序数组。

说明:
初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。

示例:
输入:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

输出: [1,2,2,3,5,6]
*/

//func main() {
//	nums1 := []int{1, 2, 3, 0, 0, 0}
//	nums2 := []int{2, 5, 6}
//	merge2([]int{2,0}, 1, []int{1}, 1)
//}

// 双指针，从前往后遍历，需要额外空间
func merge(nums1 []int, m int, nums2 []int, n int) {
	res := make([]int, 0, m+n)
	i, j := 0, 0
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}
	}
	if i < m {
		res = append(res, nums1[i:]...)
	}
	if j < n {
		res = append(res, nums2[j:]...)
	}
	copy(nums1, res)
}

// 双指针，从后往前遍历
func merge2(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1, nums2)
	}
	if n == 0 {
		return
	}
	for m > 0 || n > 0 {
		// nums1没了，把num2一个一个插进去
		if m <= 0 {
			nums1[m+n-1] = nums2[n-1]
			n--
			continue
		}
		// nums2没了，完成
		if n <= 0 {
			return
		}
		// 比较，然后插入
		if nums1[m-1] < nums2[n-1] {
			nums1[m+n-1] = nums2[n-1]
			n--
		} else {
			nums1[m+n-1] = nums1[m-1]
			m--
		}
	}
}
