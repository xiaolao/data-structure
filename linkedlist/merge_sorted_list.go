package linkedlist

func MergeSortedList(l1, l2 *SingleLinkedList) *SingleLinkedList {
	if l1 == nil || l1.head == nil || l1.head.next == nil {
		return l2
	}
	if l2 == nil || l2.head == nil || l2.head.next == nil {
		return l1
	}
	l := NewSingleLinkedList()
	cur := l.head
	cur1 := l1.head.next
	cur2 := l2.head.next
	// 类似归并排序的merge部分
	for cur1 != nil && cur2 != nil {
		if cur1.value.(int) <= cur2.value.(int) {
			cur.next = cur1
			cur1 = cur1.next
		} else {
			cur.next = cur2
			cur2 = cur2.next
		}
	}
	if cur1 != nil {
		cur.next = cur1
	} else {
		cur.next = cur2
	}
	return l
}
