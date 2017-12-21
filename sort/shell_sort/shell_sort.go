/*
description:shell sorting algorithm
author:whl
data:2017/12/21

因D.L. Shell 于1959年提出而得名，是直接插入算法的一种加强版。
当gap==1时，希尔排序算法演化为直接插入算法。

对顺序表做Shell插入排序，本算法和直接插入排序相比，作了以下修改：
1、前后记录位置的增量是gap，不是1
2、list[0]只是暂存单元，不是哨兵，当j>0时，插入位置已到
*/

package main

import "fmt"

func main() {
	src := []int{0, 9, 1, 2, 5, 7, 4, 8, 6, 3, 5}
	ShellSort(src)
	fmt.Println(src)
}

func ShellSort(list []int) {

	var gap, j int
	length := len(list)

	//D.L. Shell给出的gap计算方法
	for gap = length / 2; gap >= 1; gap /= 2 {

		for i := gap + 1; i < length; i++ {

			if list[i] < list[i-gap] {

				list[0] = list[i] //暂存在list[0]

				//对距离为gap的元素组进行排序，从后往前查找待插入位置
				for j = i - gap; j > 0 && list[0] < list[j]; j -= gap {
					list[j+gap] = list[j] //后移
				}
				list[j+gap] = list[0] //复制到插入位置
			}
		}
	}
}
