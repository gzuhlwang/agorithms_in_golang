package bubble_sort

func Bubble_sort(des [10]int, nums int) [10]int {

	//	for i := 0; i < nums; i++ {
	//		fmt.Scanf("%d", &des[i])
	//		//fmt.Println(src[i])
	//	}

	length := nums - 1

	//难点在于确定边界（i,j）
	for i := 0; i < length; i++ { //nums个数只需要排nums-1趟

		for j := 0; j < length-i; j++ {
			if des[j] < des[j+1] { //按从大到小的顺序排序
				temp := des[j]
				des[j] = des[j+1]
				des[j+1] = temp
			}
		}
	}
	return des
}
