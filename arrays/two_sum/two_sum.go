package two_sum

// TwoSum finds two numbers in the array that add up to the target
// Returns the indices of the two numbers
// Time Complexity: O(n)
// Space Complexity: O(n)
func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if j, found := seen[complement]; found {
			return []int{j, i}
		}
		seen[num] = i
	}

	return []int{}
}
