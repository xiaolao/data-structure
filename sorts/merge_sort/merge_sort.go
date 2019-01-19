package ms

func MergeSort(a []int) {
	arrLen := len(a)
	if arrLen <= 1 {
		return
	}

	mergeSort(a, 0, arrLen-1)
}

func mergeSort(a []int, low int, high int) {
	if low < high {
		mid := (low + high) / 2
		mergeSort(a, low, mid)
		mergeSort(a, mid+1, high)
		merge(a, low, mid, high)
	}
}

func merge(a []int, low int, mid int, high int) {
	i, j, k := low, mid+1, 0
	tmpArr := make([]int, high-low+1)
	for ; i <= mid && j <= high; k++ {
		if a[i] <= a[j] {
			tmpArr[k] = a[i]
			i++
		} else {
			tmpArr[k] = a[j]
			j++
		}
	}
	for ; i <= mid; i++ {
		tmpArr[k] = a[i]
		k++
	}
	for ; j <= high; j++ {
		tmpArr[k] = a[j]
		k++
	}
	copy(a[low:high+1], tmpArr)
}
