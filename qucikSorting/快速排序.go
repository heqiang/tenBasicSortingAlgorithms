package main

import "fmt"

func main() {
	fmt.Println(QuickSort([]int{9, 8, 100, 101, 7, 29, 100}, 0, 6))
}

// QuickSort 不稳定的排序
// 时间复杂度
func QuickSort(quickSlice []int, low, hight int) []int {
	if low < hight {
		i := low
		j := hight
		//pivot := quickSlice[low]
		for i < j {
			//从后往前比较 找到比 基准大的 然后交换
			for i < j && quickSlice[j] >= quickSlice[i] {
				j--
			}
			quickSlice[j], quickSlice[i] = quickSlice[i], quickSlice[j]
			//从前往后比较
			for i < j && quickSlice[i] <= quickSlice[j] {
				i++
			}
			quickSlice[i], quickSlice[j] = quickSlice[j], quickSlice[i]

		}
		QuickSort(quickSlice, low, j-1)
		QuickSort(quickSlice, j+1, hight)
	}
	return quickSlice
}
