/*
2017/10/5
author:gzuhlwang
堆排序算法主要由两部分：（1）建堆；（2）筛选
*/
package main

import "fmt"

func main() {
	source := []int{0, 16, 4, 10, 14, 7, 9, 3, 2, 10, 1}

	heap_sort(source)
	fmt.Println("排序后：")
	fmt.Println(source[1:])

}

//堆排序核心算法
func heap_sort(source []int) {
	N := len(source) - 1

	if N == 0 {
		return
	}
	//构造堆
	/*
		N个结点的完全二叉树，最后一个结点是N/2个结点的孩子。从第N/2个结点的子树进行筛选，
		使之成为对之后，向前依次对各结点为根的子树进行筛选
	*/
	for i := N / 2; i >= 1; i-- {
		heapify(source, i, N)
	}

	//筛选
	for N >= 1 {
		swap(source, 1, N)
		N--
		heapify(source, 1, N)
	}
}

//建大根堆
/*
  9      ——》      10
8   10         8		9
*/
/*index:从第index个结点的子树开始建堆；
length:堆的大小
*/
func heapify(arra []int, index, length int) {

	for {
		left := Left(index)   //第index个结点的左孩子
		right := Right(index) //第index个结点的右孩子
		largest := index      //用largest记录结点较大的编号

		//第index个结点与自己的左孩子比较
		if left <= length && arra[index] <= arra[left] {
			largest = left
		} else {
			largest = index
		}

		if right <= length && arra[largest] <= arra[right] {
			largest = right
		}

		//如果最大的结点的编号不是自己，交换位置
		if index != largest {
			swap(arra, index, largest)
			index = largest //下一轮筛选
		} else {
			break //已为大根堆
		}
	}

}

func Left(i int) int {
	return i << 1
}

func Right(i int) int {
	return (i << 1) + 1
}

func swap(s []int, x, y int) {

	s[x], s[y] = s[y], s[x] //多重赋值
}
