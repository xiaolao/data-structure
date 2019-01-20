"""
稳定性: 不稳定
空间复杂度：o(1)
时间复杂度：o(nlogn)
地推公式：quick_sort(p...r) = quick_sort(p...q-1) + quick_sort(q+1...r)
终止条件：p >= r
算法思想：选择待排数组中的一个数作为pivot,使用pivot将数组划分为三个部分，
p...q-1都是小于pivot的数，q+1...r都是大于pivot的数。
"""
import random
from typing import List


def quick_sort(a: List[int]):
    if a:
        _quick_sort_between(a, 0, len(a)-1)


def _quick_sort_between(a: List[int], low: int, high: int):
    if low < high:
        k = random.randint(low, high)
        a[low], a[k] = a[k], a[low]
        pivot_position = _partition(a, low, high)
        _quick_sort_between(a, low, pivot_position-1)
        _quick_sort_between(a, pivot_position+1, high)


def _partition(a: List[int], low: int, high: int) -> int:
    pivot = a[low]
    while low < high:
        while low < high and a[high] >= pivot:
            high -= 1
        a[low] = a[high]
        while low < high and a[low] <= pivot:
            low += 1
        a[high] = a[low]
    a[low] = pivot
    return low


if __name__ == "__main__":
    a1 = [3, 5, 6, 7, 8]
    quick_sort(a1)
    assert a1 == [3, 5, 6, 7, 8]
    a2 = [2, 2, 2, 2]
    quick_sort(a2)
    assert a2 == [2, 2, 2, 2]
    a3 = [4, 3, 2, 1]
    quick_sort(a3)
    assert a3 == [1, 2, 3, 4]
    a4 = [5, -1, 9, 3, 7, 8, 3, -2, 9]
    quick_sort(a4)
    assert a4 == [-2, -1, 3, 3, 5, 7, 8, 9, 9]
