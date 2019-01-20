"""
问题描述：o(n)时间复杂度内，找到无序数组中第K大元素。
解题思路：选择数组的最后一个元素作为pivot,将数组分割成
a[0...p], a[p], a[p+1...r]三个部分，若p+1=k这个a[p]即为
第k大元素，若p+1<k则在a[p+1...r]内递归查找，若p+1>k则在
a[0...p]内递归查找。
"""
from typing import List


def find_kth_largest(arr: List[int], k: int) -> int:
    if k < 1 or k > len(arr):
        raise Exception("k is out of range.")
    k = len(arr)-k
    start, end = 0, len(arr) - 1
    position = _partition(arr, start, end)
    while position != k:
        if position < k:
            start = position + 1
        else:
            end = position - 1
        position = _partition(arr, start, end)
    return arr[k]


def _partition(arr: List[int], low: int, high) -> int:
    pivot = arr[low]
    while low < high:
        while low < high and arr[high] >= pivot:
            high -= 1
        arr[low] = arr[high]
        while low < high and arr[low] <= pivot:
            low += 1
        arr[high] = arr[low]
    arr[low] = pivot
    return low


if __name__ == "__main__":
    a = [6, 1, 3, 5, 7, 2, 4, 9, 11, 8]
    assert find_kth_largest(a, 3) == 8
