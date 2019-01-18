"""
归并排序：采用分治的思想，使用递归来实现。
递推公式：merge_sort(p...r) = merge_sort(merge_sort(p...q), merge_sort(q...r))
终止条件：r <= q
"""

__author__ = "ceq"


from typing import List


def merge_sort(array: List[int]):
    merge_
