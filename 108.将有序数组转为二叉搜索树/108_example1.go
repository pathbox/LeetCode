package LeetCode108

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	left := nums[:mid]
	right := nums[mid+1:]

	node := &TreeNode{nums[mid], sortedArrayToBST(left), sortedArrayToBST(right)} // 不断的二分，mid的值就是当前的父节点
	return node
}
