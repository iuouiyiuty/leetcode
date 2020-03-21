package main

/*
53. 最大子序和
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
*/

//func main() {
//	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
//	fmt.Println(maxSubArray([]int{2, 3, -6, 2, 4}))
//	fmt.Println(maxSubArray([]int{4, -1, 2, 1}))
//}

func maxSubArray(nums []int) int {
	curSum := nums[0] // 当前和
	maxSum := curSum  // 最大和

	for i := 1; i < len(nums); i++ {
		// 当前和+当前元素比当前元素小，则当前和=当前元素
		if curSum+nums[i] < nums[i] {
			curSum = nums[i]
		} else {
			curSum += nums[i]
		}

		if maxSum < curSum {
			maxSum = curSum
		}
	}
	return maxSum
}
