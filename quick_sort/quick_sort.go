package main

import (
	"fmt"
)

var data = []int{41, 24, 76, 11, 45, 64, 21, 69, 19, 36}

func quickSort(data []int, left, right int) {
	if left < right {
		key := data[(right+left)/2]
		i := left
		j := right
		for {
			for data[i] < key {
				i++
			}
			for data[j] > key {
				j--
			}
			if i >= j {
				break
			}
			// 交换值
			data[i], data[j] = data[j], data[i]
		}

		quickSort(data, left, i-1)
		quickSort(data, j+1, right)
	}

}

func main() {
	fmt.Printf("before sort: %v\n", data)
	quickSort(data, 0, len(data)-1)
	fmt.Printf("after sort: %v\n", data)
}
