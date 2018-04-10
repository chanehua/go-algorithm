package main

import (
	"fmt"
)

func bubbleSort(data []byte) {
	dataLastIndex := len(data) - 1
	for i := 0; i <= dataLastIndex; i++ {
		for j := 1; j <= dataLastIndex-i; j++ {
			if data[j-1] > data[j] {
				data[j-1], data[j] = data[j], data[j-1]
			}
		}
	}

}

func main() {
	str := "514280"
	data := []byte(str)
	fmt.Printf("not sort: %v \n", string(data))
	bubbleSort(data)
	fmt.Printf("sorted: %v \n", string(data))
}
