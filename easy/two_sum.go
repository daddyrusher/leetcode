package easy

func twoSumSlow(nums []int, target int) []int {
	result := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = []int{i, j}
			}
		}
	}

	return result
}

func twoSumFast(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, found := numMap[complement]; found {
			return []int{j, i}
		}
		numMap[num] = i
	}
	return nil
}
