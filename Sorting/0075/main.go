package main

import "fmt"

func main() {
	slice := []int{1, 0, 2, 1, 0, 1, 2}
	slice = append(slice, 0)
	sortColors(slice)
	fmt.Printf("type = %T ,result = %v\n  ", slice, slice)
}

func sortColors(nums []int) {
	// x-1(种)的变数 此题x=3
	zeroIndex, oneIndex := 0, 0
	// 按照index顺序
	for i, n := range nums {
		fmt.Printf("i = %d, n = %d\n", i, n)
		// 赋值并不会改变 n ，n已经是独立变数
		// 流程一定要从大到小
		// 一开始都是赋予2
		nums[i] = 2
		// 累计"小于等于"1的index 并且替换
		if n <= 1 {
			nums[oneIndex] = 1
			oneIndex++
		}
		// 累计"小于等于"0的index 并且替换
		if n == 0 {
			nums[zeroIndex] = 0
			zeroIndex++
		}
	}
}
