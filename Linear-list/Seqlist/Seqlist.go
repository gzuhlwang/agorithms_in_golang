/*
implementing a list via sequence list
data:2017/12/26
author:whl
*/

package main

import (
	"fmt"
)

const (
	MaxSize = 20
)

//SeqList describes

type SeqList struct {
	data   [MaxSize]int //静态分配
	length int          //current length of SeqList
}

//初始化线性表
func (l *SeqList) InitSeqList(d, p int) {
	l.data[p] = d
	l.length++
}

//p是i要插入的位置
func (l *SeqList) Insert(d, p int) bool {
	//l.length<MaxSize才能插入啊
	if p >= MaxSize || p < 0 || l.length >= MaxSize {
		return false
	}
	//如果插入的位置p小于当前长度
	if p < l.length {
		for k := l.length - 1; k >= p; k-- {
			l.data[k+1] = l.data[k] //后挪
		}
		//挪出空位后将d插入p
		l.data[p] = d
		l.length++
		return true
	} else {
		l.data[l.length] = d
		l.length++
		return true
	}
}

//p:删除元素的位置
func (l *SeqList) Delete(p int) bool {
	if p < 0 || p > l.length {
		return false
	}

	//将第p个位置以后的元素前移
	for k := p; k < l.length; k++ {
		l.data[k-1] = l.data[k]
	}
	l.length--
	return true

}

func main() {
	var l SeqList
	d, p := 1, 0
	for i := 0; i < 15; i++ {
		l.InitSeqList(d, p)
		d++
		p++
	}

	fmt.Println(l)

	//插入一个元素
	l.Insert(200, 5)
	fmt.Println(l)
	//删除一个元素
	l.Delete(2)
	fmt.Println(l)

}
