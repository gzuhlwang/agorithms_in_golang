package bubble_sort

func Bubble_sort(des [10]int, nums int) [10]int {

	//	for i := 0; i < nums; i++ {
	//		fmt.Scanf("%d", &des[i])
	//		//fmt.Println(src[i])
	//	}

	length := nums - 1

	//难点在于确定边界（i,j）
	for i := 0; i < length; i++ { //nums个数只需要排nums-1趟

		flag := false //表示本趟冒泡是否发生交换的标志

		//一趟冒泡过程
		for j := 0; j < length-i; j++ {
			//按从大到小的顺序排序
			if des[j] < des[j+1] {
				des[j], des[j+1] = des[j+1], des[j] //交换
				flag = true
			}

			if flag == false {
				break //本趟遍历后没有发生交换，说明待排序列已经有序
			}
		}
	}
	return des
}
