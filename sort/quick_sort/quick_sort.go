//package main
package quick_sort

//import "fmt"

//func main() {

//	var arra = [10]int{}
//	//读入待排数据
//	for i := 0; i <= 9; i++ {
//		fmt.Scanf("%d", &arra[i])
//	}

//	//这里必须传指针，否则原始数组是什么就会输出什么。
//	quick_sort(&arra, 0, 9)

//	//将快排后的结果打印出来
//	for i, _ := range arra {
//		fmt.Printf("%d"+"  ", arra[i])
//	}
//}

func quick_sort(origin *[10]int, left, right int) {

	//直接退出
	if left > right {
		return
	}
	//基准数
	base := origin[left]

	i := left
	j := right

	for i != j { //从左右两边交替扫描，知道i==j
		//先从右往左找。找到第一个比基准元素小的元素
		for origin[j] >= base && i < j {
			j--
		}

		//再从左往右找。找到第一个比基准元素大的元素
		for origin[i] <= base && i < j {
			i++
		}

		//交换两个数在数组中的位置
		if i < j {
			temp := origin[i]
			origin[i] = origin[j]
			origin[j] = temp
		}

	}

	//when i meets j
	origin[left] = origin[i]
	origin[i] = base //将基准数归位

	quick_sort(origin, left, i-1)  //继续排左边的
	quick_sort(origin, i+1, right) //继续排右边的

}
