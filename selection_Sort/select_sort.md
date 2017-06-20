## 算法原理
选择排序（Selection Sort）是一种简单直观的排序算法。它的工作原理如下，首先在未排序序列中找到    最小（大）元素，存放到排序序列的起始位置，然后，再从剩余未排序元素中继续寻找最小（大）元素，然后    放到已排序序列的末尾。以此类推，直到所有元素均排序完毕。    
    
选择排序的主要优点与数据移动有关。如果某个元素位于正确的最终位置上，则它不会被移动。选择排序每次    交换一对元素，它们当中至少有一个将被移到其最终位置上，因此对n个元素的序列进行排序总共进行至多    n-1次交换。在所有的完全依靠交换去移动元素的排序方法中，选择排序属于非常好的一种。    
    
![](../images/selection_sort_animation.gif)    
图片来源于维基百科    
    
## 实例分析
以数组 arr = [8, 5, 2, 6, 9, 3, 1, 4, 0, 7] 为例，先直观看一下每一步的变化，后面再介绍细节    
    
第一次从数组 [8, 5, 2, 6, 9, 3, 1, 4, 0, 7] 中找到最小的数 0，放到数组的最前面（与第一个元素进行交换）：    
    
```
                               min
                                ↓
8   5   2   6   9   3   1   4   0   7
↑                               ↑
└───────────────────────────────┘


```    
    
交换后：    
    
```
0   5   2   6   9   3   1   4   8   7

```    
    
在剩余的序列中 [5, 2, 6, 9, 3, 1, 4, 8, 7] 中找到最小的数 1，与该序列的第一个个元素进行位置交换：    
    
```
                       min
                        ↓
0   5   2   6   9   3   1   4   8   7
    ↑                   ↑
    └───────────────────┘


```    
    
交换后：    
    
```
0   1   2   6   9   3   5   4   8   7

```    
    
在剩余的序列中 [2, 6, 9, 3, 5, 4, 8, 7] 中找到最小的数 2，与该序列的第一个个元素进行位置交换（实际上不需要交换）：    
    
```
       min
        ↓
0   1   2   6   9   3   5   4   8   7
        ↑


```    
    
重复上述过程，直到最后一个元素就完成了排序。    
    
```
                   min
                    ↓
0   1   2   6   9   3   5   4   8   7
            ↑       ↑
            └───────┘
                           min
                            ↓
0   1   2   3   9   6   5   4   8   7
                ↑           ↑
                └───────────┘
                       min
                        ↓
0   1   2   3   4   6   5   9   8   7
                    ↑   ↑
                    └───┘
                       min
                        ↓
0   1   2   3   4   5   6   9   8   7
                        ↑   
                                   min
                                    ↓
0   1   2   3   4   5   6   9   8   7
                            ↑       ↑
                            └───────┘  
                               min
                                ↓
0   1   2   3   4   5   6   7   8   9
                                ↑      
                                   min
                                    ↓
0   1   2   3   4   5   6   7   8   9
                                    ↑


```    
    
![](../images/Selection-Sort-Animation.gif)    
图片来源于维基百科    

## go语言实现

```go
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


```

## 参考文章
+ [en.wikipedia.org](https://en.wikipedia.org/wiki/Selection_sort)
+ [wikibooks](https://en.wikibooks.org/wiki/Algorithm_Implementation/Sorting/Selection_sort)
+ [维基百科](https://zh.wikipedia.org/wiki/选择排序)
+ [直接选择排序(Straight Selection Sort)](http://student.zjzk.cn/course_ware/data_structure/web/paixu/paixu8.4.1.htm)
+ [经典排序算法 - 选择排序 Selection Sort](http://www.cnblogs.com/kkun/archive/2011/11/23/2260281.html)