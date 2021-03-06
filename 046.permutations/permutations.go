package LeetCode046

func permute(nums []int) [][]int {
	n := len(nums)
	// vector 是一组可能的解答
	vector := make([]int, n)
	// taken[i] == true 表示 nums[i] 已经出现在了 vector 中
	taken := make([]bool, n)

	var ans [][]int

	makePermutation(0, n, nums, vector, taken, &ans)

	return ans
}

func makePermutation(cur, n int, nums, vector []int, taken []bool, ans *[][]int) {
	if cur == n { // 当前索引==数组长度，说明这一组比较到了最后
		tmp := make([]int, n)
		copy(tmp, vector)
		*ans = append(*ans, tmp)
		return
	}

	for i := 0; i < n; i++ {
		if !taken[i] {
			// 准备使用 nums[i]，所以，taken[i] == true
			taken[i] = true
			// NOTICE: 是 vector[cur]
			vector[cur] = nums[i] // cur的理解很重要

			makePermutation(cur+1, n, nums, vector, taken, ans)

			// 下一个循环中
			// vector[cur] = nums[i+1]
			// 所以，在这个循环中，恢复 nums[i] 自由
			taken[i] = false
		}
	}
}

// 每到cur位置，从nums中选出一个没有被taken的数，放到cur位置，当一组情况满足返回后
