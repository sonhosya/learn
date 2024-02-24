package twosum

/*
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
你可以按任意顺序返回答案。

思路：
遍历数组，
获取数组元素与target的差值k，
使用k尝试在map中取值，
有值则直接返回当前两个数据元素的下标（注意先后顺序），
无对应值就放入map中（方便后续直接取值进行快速比对），
重复上述步骤，直到算出结果数据
*/
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := range nums {
		if j, ok := m[target-nums[i]]; ok {
			return []int{j, i}
		}
		m[nums[i]] = i
	}
	return nil
}
