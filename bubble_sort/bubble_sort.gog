package main

import (
	"fmt"
)

func bubbleSort(data []int) {
	dataLastIndex := len(data) - 1
	for i := 0; i < dataLastIndex; i++ {
		for j := 1; j < dataLastIndex-i; j++ {
			if data[j-1] > data[j] {
				data[j-1], data[j] = data[j], data[j-1]
			}
		}
	}

}

func main() {
	data := []int{5, 1, 4, 2, 8}
	fmt.Printf("not sort: %v \n", data)
	bubbleSort(data)
	fmt.Printf("sorted: %v \n", data)
}
