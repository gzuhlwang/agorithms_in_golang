package bubble_sort

import (
	//	"fmt"
	"testing"
)

func TestBubble_sort(t *testing.T) {
	//fmt.Println("testing begins!")
	var result [10]int
	result = Bubble_sort([...]int{30, 11, 1000, 22, 33, 13, 9, 10, 7, 111}, 10)
	if result == [10]int{1000, 111, 33, 30, 22, 13, 11, 10, 9, 7} {
		t.Log("the result is ok")
	} else {
		t.Fatal("the result is wrong")
	}
}
