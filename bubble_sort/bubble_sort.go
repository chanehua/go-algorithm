package main

import (
	"fmt"
)

func bubbleSort(arr []int) {
	arrLastIndex := len(arr) - 1
	for i := 0; i < arrLastIndex; i++ {
		for j := 1; j < arrLastIndex-i; j++ {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}

}

func main() {
	arr := []int{5, 1, 4, 2, 8}
	fmt.Printf("not sort: %v \n", arr)
	bubbleSort(arr)
	fmt.Printf("sorted: %v \n", arr)
}
