## 冒泡排序

### 算法原理    
冒泡排序（Bubble Sort，台湾译为：泡沫排序或气泡排序）是一种简单的排序算法。它重复地走访过要排序的数列，一次比较两个元素，如果他们的顺序错误就把他们交换过来。走访数列的工作是重复地进行    直到没有再需要交换，也就是说该数列已经排序完成。这个算法的名字由来是因为越小的元素会经由交换慢    慢“浮”到数列的顶端。    
冒泡排序算法的流程如下：    

1. 比较相邻的元素。如果第一个比第二个大，就交换他们两个。    
2. 对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对。在这一点，最后的元素应该会是最大的数。    
3. 针对所有的元素重复以上的步骤，除了最后一个。
4. 持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。    

由于它的简洁，冒泡排序通常被用来对于程序设计入门的学生介绍算法的概念。    

![](../images/Bubble_sort_animation.gif)      
图片来自维基百科    


### 实例分析    
以数组 arr = [5, 1, 4, 2, 8] 为例说明，加粗的数字表示每次循环要比较的两个数字：    
第一次外循环    
( **5 1** 4 2 8 ) → ( **1 5** 4 2 8 )， 5 > 1 交换位置    
( 1 **5 4** 2 8 ) → ( 1 **4 5** 2 8 )， 5 > 4 交换位置    
( 1 4 **5 2** 8 ) → ( 1 4 **2 5** 8 )， 5 > 2 交换位置    
( 1 4 2 **5 8** ) → ( 1 4 2 **5 8** )， 5 < 8 位置不变    
第二次外循环（除开最后一个元素8，对剩余的序列）    
( **1 4** 2 5 8 ) → ( **1 4** 2 5 8 )， 1 < 4 位置不变    
( 1 **4 2** 5 8 ) → ( 1 **2 4** 5 8 )， 4 > 2 交换位置    
( 1 2 **4 5** 8 ) → ( 1 2 **4 5** 8 )， 4 < 5 位置不变    
第三次外循环（除开已经排序好的最后两个元素，可以注意到上面的数组其实已经排序完成，但是程序本身并    不知道，所以还要进行后续的循环，直到剩余的序列为 1）    
( **1 2** 4 5 8 ) → ( **1 2** 4 5 8 )    
( 1 **2 4** 5 8 ) → ( 1 **2 4** 5 8 )    
第四次外循环（最后一次）    
( **1 2** 4 5 8 ) → ( **1 2** 4 5 8 )     

### go语言实现    

```go
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


```


### 参考文章    

 + [en.wikipedia.org](https://en.wikipedia.org/wiki/Bubble_sort)    
 + [维基百科，自由的百科全书](https://zh.wikipedia.org/wiki/冒泡排序) 
 + [Bubble Sort](https://www.toptal.com/developers/sorting-algorithms/bubble-sort) 
 + [经典排序算法 - 冒泡排序Bubble sort](http://www.cnblogs.com/kkun/archive/2011/11/23/2260280.html) 
 + [冒泡排序](http://student.zjzk.cn/course_ware/data_structure/web/paixu/paixu8.3.1.1.htm) 