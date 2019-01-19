package ms

import "testing"

func TestMergeSort(t *testing.T) {
	a1 := []int{3, 5, 6, 7, 8}
	MergeSort(a1)
	t.Log(a1)
	a2 := []int{2, 2, 2, 2}
	MergeSort(a2)
	t.Log(a2)
	a3 := []int{4, 3, 2, 1}
	MergeSort(a3)
	t.Log(a3)
	a4 := []int{5, -1, 9, 3, 7, 8, 3, -2, 9}
	MergeSort(a4)
	t.Log(a4)
}
