/* find the k-th largest element in an array with o(n)
time complexity and o(1) space complexity. */
package main

func FindKthLargest(a []int, k int) int {
	if k < 1 || k > len(a) {
		panic("k is out of range.")
	}
	k = len(a) - k
	start, end := 0, len(a)-1
	p := partition(a, start, end)
	for p != k {
		if p < k {
			start = p + 1
		} else {
			end = p - 1
		}
		p = partition(a, start, end)
	}
	return a[k]
}

func partition(a []int, start, end int) int {
	pivot := a[end]
	for start < end {
		for start < end && a[start] <= pivot {
			start++
		}
		a[end] = a[start]
		for start < end && a[end] >= pivot {
			end--
		}
		a[start] = a[end]
	}
	a[start] = pivot
	return start
}
