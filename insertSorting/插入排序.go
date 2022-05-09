package main

import "fmt"

func main() {
	fmt.Println(InsertSorting([]int{9, 8, 100, 101, 7, 29, 100}))
}

func InsertSorting(insertSlice []int) []int {
	for i := 1; i < len(insertSlice); i++ {
		//循环前面排好序的切片
		for j := i; j > 0; j-- {
			//只要未排序的数字小于排序的就一直交换
			if insertSlice[j] < insertSlice[j-1] {
				insertSlice[j-1], insertSlice[j] = insertSlice[j], insertSlice[j-1]
			}
		}
	}
	return insertSlice
}
