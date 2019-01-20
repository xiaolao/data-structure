package qs

func QuickSort(a []int) {
	arrLen := len(a)
	if arrLen <= 1 {
		return
	}
	quickSort(a, 0, arrLen-1)
}

func quickSort(a []int, low, high int) {
	if low >= high {
		return
	}
	pivotPos := partition(a, low, high)
	quickSort(a, low, pivotPos-1)
	quickSort(a, pivotPos+1, high)
}

func partition(a []int, low, high int) int {
	// 通过游标i将a分成两个部分
	// a[p...i-1]都是小于pivot的数，为已处理区间
	// a[i...r-1]是为处理区间，每次从未处理区间
	// 取一个元素a[j]与pivot作比较，小于pivot则加
	// 到已处理区间尾部,即a[i]的位置。
	var i = low
	pivot := a[high]
	for j := low; j < high; j++ {
		if a[j] < pivot {
			if i != j {
				a[i], a[j] = a[j], a[i]
			}
			i++
		}
	}
	a[i], a[high] = a[high], a[i]
	return i
}
