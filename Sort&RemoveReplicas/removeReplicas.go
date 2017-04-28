package main

import (
	"fmt"
)

func main() {
	//原序列
	origin := []int{2, 11, 100, 2, 99, 99, 11, 100, 100}

	res := removeReplicas(origin) //返回去重的序列

	length := len(res)
	//倒序输出
	for i := length - 1; i >= 0; i-- {
		fmt.Println(res[i])
	}

}

//消除重复元素
func removeReplicas(org []int) []int {

	midRes := make(map[int]int) //初始化一个map
	length := len(org)

Lable:
	//从后往前遍历
	for i := length - 1; i >= 0; i-- {
		//在midRes中查找一个特定的键origin[i]是否存在
		_, ok := midRes[org[i]]
		//若存在，重新做循环
		if ok {
			continue Lable
		} else {
			//赋值
			midRes[org[i]] = i
		}
	}

	var res []int //声明了一个int型切片，用来存放midRes中的键
	for k, _ := range midRes {
		res = append(res, k) //将键存进res
	}

	return res

}
