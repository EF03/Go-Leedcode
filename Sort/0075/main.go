package main

import "fmt"

func main() {
	array := []int{1, 0, 2, 1, 0, 1, 2}
	sortColors(array)
	fmt.Printf("%v", array)
}

func sortColors(nums []int) {
	zero, one := 0, 0
	for i, n := range nums {
		nums[i] = 2
		if n <= 1 {
			nums[one] = 1
			one++
		}
		if n == 0 {
			nums[zero] = 0
			zero++
		}
	}
}
