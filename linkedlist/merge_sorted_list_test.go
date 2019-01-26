package linkedlist

import "testing"

var l *SingleLinkedList

func init() {
	n5 := &SingleListNode{value: 5}
	n4 := &SingleListNode{value: 4, next: n5}
	n3 := &SingleListNode{value: 3, next: n4}
	n2 := &SingleListNode{value: 2, next: n3}
	n1 := &SingleListNode{value: 1, next: n2}
	l = &SingleLinkedList{head: &SingleListNode{next: n1}}
}

func TestReverse(t *testing.T) {
	l.Print()
	l.Reverse()
	l.Print()
}

func TestHasCycle(t *testing.T) {
	t.Log(l.HasCycle())
	l.head.next.next.next.next.next.next = l.head.next.next.next
	t.Log(l.HasCycle())
}

func TestMergeSortedList(t *testing.T) {
	n5 := &SingleListNode{value: 9}
	n4 := &SingleListNode{value: 7, next: n5}
	n3 := &SingleListNode{value: 5, next: n4}
	n2 := &SingleListNode{value: 3, next: n3}
	n1 := &SingleListNode{value: 1, next: n2}
	l1 := &SingleLinkedList{head: &SingleListNode{next: n1}}

	n10 := &SingleListNode{value: 10}
	n9 := &SingleListNode{value: 8, next: n10}
	n8 := &SingleListNode{value: 6, next: n9}
	n7 := &SingleListNode{value: 4, next: n8}
	n6 := &SingleListNode{value: 2, next: n7}
	l2 := &SingleLinkedList{head: &SingleListNode{next: n6}}

	l1.Print()
	l2.Print()
	l = MergeSortedList(l1, l2)
}

func TestDeleteBottomN(t *testing.T) {
	l.Print()
	l.DeleteBottomN(3)
	l.Print()
}

func TestFindMiddleNode(t *testing.T) {
	l.DeleteBottomN(1)
	l.DeleteBottomN(1)
	l.DeleteBottomN(1)
	l.DeleteBottomN(1)
	l.Print()
	t.Log(l.FindMiddleNode())
}
