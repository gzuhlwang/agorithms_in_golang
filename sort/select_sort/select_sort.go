package main

import "fmt"

func main(
	sl := []int{49, 38, 65, 97, 76, 13, 27}
	SelectSort(sl, 7)
	fmt.Println(sl)
)


func SelectSort(arr []int, n int) {
	var min int
	for i := 0; i <= n-1; i++ {
		min = i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != i {
			arr[i], arr[min] = arr[min], arr[i] //交换
		}
	}
}