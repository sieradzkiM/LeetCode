package main

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Output: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]

*/

func two_sum(nums []int, target int) []int {
	// Create the hashtable
	m := make(map[int]int)
	// Loop throug the entire slice
	for index1, numL1 := range nums {
		// Check if target - numL1 exist in the hashtable
		if index2, found := m[target-numL1]; found {
			//  Return new slice []int with index1 and index2
			return []int{index2, index1}
		}
		// If not add numL1 as key and index 1 as value
		m[numL1] = index1
	}

	return nil
}
func two_sum_brutforce(nums []int, target int) []int {
	// Loop through the entire slice
	for i, numL1 := range nums {
		// Loop throught i+1 till the end - can't use the same element
		for j, numL2 := range nums[i+1:] {
			// match j index to nums[i+1} - otherwise j start from 0
			j += i + 1
			// Check if the sum of numL1 (return from first loop) and numL2 (return from second loop) equals target
			if numL1+numL2 == target {
				// If yes return the new slice of int with indexes from the passed slice to the function
				return []int{i, j}
			}
		}
	}

	return nil
}

func main() {
	twoSum([]int{2, 7, 11, 15}, 9)
	twoSum([]int{3, 2, 4}, 6)
	twoSum([]int{3, 3}, 6)
}
