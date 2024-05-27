package main

func twoSum(nums []int, target int) []int {
	for idx, num := range nums {
		for idx2 := idx + 1; idx2 < len(nums); idx2++ {
			if num+nums[idx2] == target {
				return []int{idx, idx2}
			}
		}
	}

	return nil
}
