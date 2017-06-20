package main

import (
	"fmt"
)

func selectSort(data []int) {
	lenght := len(data)
	var i, minIdenx, minValue int

	for i = 0; i < lenght-1; i++ {
		minIdenx = i
		minValue = data[i]

		// 获取最小值
		for j := i + 1; j < lenght; j++ {
			if data[j] < minValue {
				minIdenx = j
				minValue = data[minIdenx]
			}
		}

		//交换位置
		data[minIdenx], data[i] = data[i], data[minIdenx]

	}
}

func main() {
	data := []int{2, 6, 9, 3, 5, 4, 8, 7, 1, 10, 0}
	fmt.Printf("before sort: %v\n", data)
	selectSort(data)
	fmt.Printf("sorted: %v\n", data)
}
