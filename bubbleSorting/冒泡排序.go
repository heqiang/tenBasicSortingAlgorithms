package main

import "fmt"

func main() {
	fmt.Println(BubbleSorting([]int{9, 8, 100, 101, 7, 29, 100}))
	fmt.Println(BubbleSortingPlus([]int{9, 8, 100, 101, 7, 29, 100}))
}

// BubbleSorting 冒泡排序
// 每次排序都是挨着的两个互相比较
func BubbleSorting(bubble []int) []int {
	for i := 0; i < len(bubble); i++ {
		for j := 0; j < len(bubble)-1; j++ {
			if bubble[j+1] < bubble[j] {
				bubble[j], bubble[j+1] = bubble[j+1], bubble[j]
			}
		}
	}
	return bubble
}

func BubbleSortingPlus(bubble []int) []int {
	for i := 0; i < len(bubble); i++ {
		newBubble := len(bubble)
		for j := 0; j < newBubble-1; j++ {
			if bubble[j+1] < bubble[j] {
				bubble[j], bubble[j+1] = bubble[j+1], bubble[j]
				newBubble--
			}
		}
	}
	return bubble
}
