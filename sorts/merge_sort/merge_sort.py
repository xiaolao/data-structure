"""
归并排序：采用分治的思想，使用递归来实现。
递推公式：merge_sort(p...r) = merge_sort(merge_sort(p...q), merge_sort(q...r))
终止条件：r <= q
"""
from copy import copy
from typing import List


def merge_sort(a: List[int]):
    _merge_sort_between(a, 0, len(a)-1)


def _merge_sort_between(a: List[int], low: int, high: int):
    if low < high:
        mid = low + (high-low)//2
        _merge_sort_between(a, low, mid)
        _merge_sort_between(a, mid+1, high)
        _merge(a, low, mid, high)


def _merge(a: List[int], low: int, mid: int, high: int):
    tmp = []
    i, j = low, mid + 1
    while i <= mid and j <= high:
        if a[i] <= a[j]:
            tmp.append(a[i])
            i += 1
        else:
            tmp.append(a[j])
            j += 1
    start, end = (i, mid) if i <= mid else (j, high)
    tmp.extend(a[start:end+1])
    a[low:high+1] = tmp


if __name__ == "__main__":
    a1 = [3, 5, 6, 7, 8]
    merge_sort(a1)
    assert a1 == [3, 5, 6, 7, 8]
    a2 = [2, 2, 2, 2]
    merge_sort(a2)
    assert a2 == [2, 2, 2, 2]
    a3 = [4, 3, 2, 1]
    merge_sort(a3)
    assert a3 == [1, 2, 3, 4]
    a4 = [5, -1, 9, 3, 7, 8, 3, -2, 9]
    merge_sort(a4)
    assert a4 == [-2, -1, 3, 3, 5, 7, 8, 9, 9]