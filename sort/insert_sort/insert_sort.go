/*
	description:insertion sort algorithm
	author:whl
	data:2017/12/21
*/

package main

import "fmt"

func main() {

	src := []int{0, 6, 3, 9, 10, 1, 4, 55, 23}
	InsertSort(src)
	//将排完序的序列打印出来
	fmt.Println(src)

}

func InsertSort(src []int) {
	l := len(src)
	j := 0
	/*
		src[0]设置为监视哨，src[0]不存放元素
		第一个数src[i]肯定是有序的（仅有一个记录的表总是有序的），从第2个数开始遍历，依次插入有序序列
		第一个数:6
		第二个数:3
		...
		最后一个数：23
	*/
	for i := 2; i < l; i++ {
		src[0] = src[i] //src[0]设置为监视哨

		//前i-1个数都是由小到大的有序序列，只要监视哨监视的数（src[0]）比有序序列中当前比较的数（src[j]）小，就把这个数后移一位
		//for j = i - 1; j > 0 && src[0] < src[j]; j-- {
		for j = i - 1; src[0] < src[j]; j-- {
			src[j+1] = src[j] //后移
		}
		//监视哨监视的数(src[0])比有序序列中最大的数还大，直接在尾部插入
		src[j+1] = src[0] //插入
	}
}
