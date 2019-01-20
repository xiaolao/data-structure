package qs

import "testing"

func TestQuickSort(t *testing.T) {
	a1 := []int{3, 5, 6, 7, 8}
	QuickSort(a1)
	t.Log(a1)
	a2 := []int{2, 2, 2, 2}
	QuickSort(a2)
	t.Log(a2)
	a3 := []int{4, 3, 2, 1}
	QuickSort(a3)
	t.Log(a3)
	a4 := []int{5, -1, 9, 3, 7, 8, 3, -2, 9}
	QuickSort(a4)
	t.Log(a4)
}
