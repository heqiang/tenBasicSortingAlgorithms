package main

import "fmt"

func main() {
	fmt.Println(ShellSoring([]int{9, 8, 100, 101, 7, 29, 100}))
}

// ShellSoring 希尔排序 不稳定
// 分组 有点归并的思想
func ShellSoring(shellSlice []int) []int {
	for gap := len(shellSlice) / 2; gap > 0; gap /= 2 {
		for i := 0; i < len(shellSlice); i++ {
			for j := i; j >= gap; j -= gap {
				if shellSlice[j] < shellSlice[j-gap] {
					shellSlice[j], shellSlice[j-gap] = shellSlice[j-gap], shellSlice[j]
				}
			}
		}
	}
	return shellSlice
}
