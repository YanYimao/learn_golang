package main

// 两数之和
func twoSum(nums []int, target int) []int {
	myMap := map[int]int{}
	for i, n := range nums {
		if v, ok := myMap[target-n]; ok {
			return []int{v, i}
		}
		myMap[n] = i
	}
	return []int{}
}
