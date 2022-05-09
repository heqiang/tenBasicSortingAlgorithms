package main

import "fmt"

func main() {
	fmt.Println(HeadSort([]int{9, 8, 100, 101, 7, 29, 100, 323, 11}))
}

func HeadSort(headSlice []int) []int {
	//初始建堆 从非叶子节点的最后一个开始
	for i := 0; i < (len(headSlice)/2)-1; i++ {
		happily(headSlice, i, len(headSlice)-1)
	}
	//堆排序
	for i := len(headSlice) - 1; i > 0; i-- {
		//交换当前节点和堆顶节点,堆排序会乱 所以需要再次排序
		headSlice[0], headSlice[i] = headSlice[i], headSlice[0]
		happily(headSlice, 0, i-1)
	}
	return headSlice
}

// 堆排序
func happily(arr []int, i, lastIndex int) {
	maxNode := i
	//判断左节点
	if 2*i+1 <= lastIndex && arr[2*i+1] > arr[maxNode] {
		maxNode = 2*i + 1
	}
	//判断右节点
	if 2*i+2 <= lastIndex && arr[2*i+2] > arr[maxNode] {
		maxNode = 2*i + 2
	}
	if maxNode != i {
		arr[i], arr[maxNode] = arr[maxNode], arr[i]
		happily(arr, maxNode, lastIndex)
	}

}
