package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for i, num := range nums {
		numsMap[num] = i
	}

	for i, num := range nums {
		j, ok := numsMap[target-num]
		if ok && i != j {
			return []int{i, j}
		}
	}

	return []int{}
}

func main() {
	nums := []int{3, 3}
	target := 6
	result := twoSum(nums, target)

	fmt.Println(result)
}
