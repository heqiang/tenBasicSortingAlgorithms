package main

import "fmt"

func main() {
	fmt.Println(ChoiceSorting([]int{9, 8, 100, 101, 7, 29, 100}))
}

// ChoiceSorting 选择排序
// 把第一个数的下表当作最小的，然后从后面的数中找出最小的数的下标,找到后和已经排好序的数字进行对比，小就交换
// 因为每次都是找的最小的，不需要和已经排好序的在进行对比
func ChoiceSorting(choiceSlice []int) []int {
	for i := 0; i < len(choiceSlice); i++ {
		min := i
		for j := i + 1; j < len(choiceSlice); j++ {
			if choiceSlice[j] < choiceSlice[j-1] {
				min = j
			}
		}
		if choiceSlice[min] < choiceSlice[i] {
			choiceSlice[min], choiceSlice[i] = choiceSlice[i], choiceSlice[min]
		}
	}
	return choiceSlice
}
