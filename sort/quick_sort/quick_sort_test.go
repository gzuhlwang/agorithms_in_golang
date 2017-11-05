package quick_sort

import "testing"

func TestQuick_Sort(t *testing.T) {
	source1 := [10]int{6, 1, 2, 7, 9, 3, 4, 5, 10, 8}

	quick_sort(&source1, 0, 9)

	if source1 == [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		t.Log("Quicksort algorithm is right!")
	}

	source2 := [10]int{6, 1000, 99, 200, 201, 9, 3, 4, 5, 10}

	quick_sort(&source2, 0, 9)

	if source2 == [...]int{3, 4, 5, 6, 9, 10, 99, 200, 201, 1000} {
		t.Log("Quicksort algorithm is right!")
	}
}
