## 算法原理
快速排序是图灵奖得主 [C. R. A. Hoare](https://zh.wikipedia.org/wiki/東尼·霍爾) 于 1960 年提出的一种划分交换排序。它采用了一种分治的策略，通常称其为[分治法(Divide-and-ConquerMethod)](https://en.wikipedia.org/wiki/Quicksort)。    
![](../images/C.R.A.Hoare.jpg)    
C. R. A. Hoare    
    
分治法的基本思想是：将原问题分解为若干个规模更小但结构与原问题相似的子问题。递归地解这些子问题，    然后将这些子问题的解组合为原问题的解。    
    

利用分治法可将快速排序的分为三步：    
    
1. 在数据集之中，选择一个元素作为”基准”（pivot）。    
2. 所有小于”基准”的元素，都移到”基准”的左边；所有大于”基准”的元素，都移到”基准”的右边。这个操作称为分区 (partition) 操作，分区操作结束后，基准元素所处的位置就是最终排序后它的位置。    
3. 对”基准”左边和右边的两个子集，不断重复第一步和第二步，直到所有子集只剩下一个元素为止。    
![](../images/Sorting_quicksort_anim.gif)    
图片来自维基百科    

分区是快速排序的主要内容，用伪代码可以表示如下：   
 
```
function partition(a, left, right, pivotIndex)
     pivotValue := a[pivotIndex]
     swap(a[pivotIndex], a[right]) // 把 pivot 移到結尾
     storeIndex := left
     for i from left to right-1
         if a[i] < pivotValue
             swap(a[storeIndex], a[i])
             storeIndex := storeIndex + 1
     swap(a[right], a[storeIndex]) // 把 pivot 移到它最後的地方
     return storeIndex // 返回 pivot 的最终位置

```    
    
首先，把基准元素移到結尾（如果直接选择最后一个元素为基准元素，那就不用移动），然后从左到右（除了.   最后的基准元素），循环移动小于等于基准元素的元素到数组的开头，每次移动 storeIndex 自增     1，表示下一个小于基准元素将要移动到的位置。循环结束后 storeIndex 所代表的的位置就是基准元素的.   所有摆放的位置。所以最后将基准元素所在位置（这里是 right）与 storeIndex 所代表的的位置.   的元素交换位置。要注意的是，一个元素在到达它的最后位置前，可能会被交换很多次。    

一旦我们有了这个分区算法，要写快速排列本身就很容易：  
  
```
procedure quicksort(a, left, right)
    if right > left
        select a pivot value a[pivotIndex]
        pivotNewIndex := partition(a, left, right, pivotIndex)
        quicksort(a, left, pivotNewIndex-1)
        quicksort(a, pivotNewIndex+1, right)
```    
    
## 实例分析    
举例来说，现有数组 arr = [3,7,8,5,2,1,9,5,4]，分区可以分解成以下步骤：    
1. 首先选定一个基准元素，这里我们元素 5 为基准元素（基准元素可以任意选择）：    

```
pivot
            ↓
3   7   8   5   2   1   9   5   4

```     
1. 将基准元素与数组中最后一个元素交换位置，如果选择最后一个元素为基准元素可以省略该步：    

```
pivot
                                ↓
3   7   8   4   2   1   9   5   5

```    
    
1. 从左到右（除了最后的基准元素），循环移动小于基准元素 5 的所有元素到数组开头，留下大于等于基    准元素的元素接在后面。在这个过程它也为基准元素找寻最后摆放的位置。循环流程如下：    
    
    循环 i == 0 时，storeIndex == 0，找到一个小于基准元素的元素 3，那么将其与 storeIndex 所在.   位置的元素交换位置，这里是 3 自身，交换后将 storeIndex 自增 1，storeIndex == 1：    
    
```
                                pivot
                                  ↓
  3   7   8   4   2   1   9   5   5
  ↑
storeIndex

```    

   循环 i == 3 时，storeIndex == 1，找到一个小于基准元素的元素 4：     
   
```   
        ┌───────┐                 pivot
     ↓       ↓                   ↓
 3   7   8   4   2   1   9   5   5
     ↑       ↑
storeIndex   i

```    
   
   交换位置后，storeIndex 自增 1，storeIndex == 2：    
    
```
                              pivot
                                ↓
3   4   8   7   2   1   9   5   5
        ↑           
   storeIndex    
```    

循环 i == 4 时，storeIndex == 2，找到一个小于基准元素的元素 2：    
    
```
        ┌───────┐             pivot
        ↓       ↓               ↓
3   4   8   7   2   1   9   5   5
        ↑       ↑
   storeIndex   i

```    
    
交换位置后，storeIndex 自增 1，storeIndex == 3：    

```
                              pivot
                                ↓
3   4   2   7   8   1   9   5   5
            ↑           
       storeIndex

```    

循环 i == 5 时，storeIndex == 3，找到一个小于基准元素的元素 1：    

```
            ┌───────┐         pivot
            ↓       ↓           ↓
3   4   2   7   8   1   9   5   5
            ↑       ↑
       storeIndex   i

```

交换后位置后，storeIndex 自增 1，storeIndex == 4：    

```
                              pivot
                                ↓
3   4   2   1   8   7   9   5   5
                ↑           
           storeIndex

```


循环 i == 7 时，storeIndex == 4，找到一个小于等于基准元素的元素 5：    

```
                ┌───────────┐ pivot
                ↓           ↓   ↓
3   4   2   1   8   7   9   5   5
                ↑           ↑
           storeIndex       i

```    

交换后位置后，storeIndex 自增 1，storeIndex == 5：    

```
                              pivot
                                ↓
3   4   2   1   5   7   9   8   5
                    ↑           
               storeIndex

```

2.循环结束后交换基准元素和 storeIndex 位置的元素的位置：     
   
```
                  pivot
                    ↓
3   4   2   1   5   5   9   8   7
                    ↑           
               storeIndex

```

那么 storeIndex 的值就是基准元素的最终位置，这样整个分区过程就完成了。    
    
引用[维基百科](https://en.wikipedia.org/wiki/Quicksort)上的一张图片：    
![](../images/Partition_example.svg.png)    
图片来自维基百科    

## golang实现

```go
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


```




    
## 参考文章    
+ [wiki Quicksort](https://en.wikipedia.org/wiki/Quicksort)
+ [维基百科 - 快速排序](https://zh.wikipedia.org/wiki/快速排序)
+ [快速排序（Quicksort）的Javascript实现](http://www.ruanyifeng.com/blog/2011/04/quicksort_in_javascript.html)
+ [Quicksort in JavaScript](http://www.cnblogs.com/ethanzheng/archive/2013/02/20/quicksort-in-javascript.html)
+ [经典排序算法 - 快速排序Quick sort](http://www.cnblogs.com/kkun/archive/2011/11/23/2260270.html)
+ [快速排序(QuickSort)](http://student.zjzk.cn/course_ware/data_structure/web/paixu/paixu8.3.2.1.htm)
+ [Friday Algorithms: Quicksort – Difference Between PHP and JavaScript](http://www.stoimen.com/blog/2010/06/11/friday-algorithms-quicksort-difference-between-php-and-javascript/)