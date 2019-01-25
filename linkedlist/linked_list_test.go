package linkedlist

import "testing"

func TestInsertToHead(t *testing.T) {
	s := NewSingleLinkedList()
	b := NewBoubleListedList()
	for i := 0; i < 10; i++ {
		s.InsertToHead(i + 1)
		b.InsertToHead(i + 1)
	}
	s.Print()
	b.Print()
}

func TestInsertToTail(t *testing.T) {
	s := NewSingleLinkedList()
	b := NewBoubleListedList()
	for i := 0; i < 10; i++ {
		s.InsertToTail(i + 1)
		b.InsertToTail(i + 1)
	}
	s.Print()
	b.Print()
}

func TestFindByIndex(t *testing.T) {
	s := NewSingleLinkedList()
	b := NewBoubleListedList()
	for i := 0; i < 10; i++ {
		s.InsertToTail(i + 1)
		b.InsertToTail(i + 1)
	}
	t.Log(s.FindByIndex(0), b.FindByIndex(0))
	t.Log(s.FindByIndex(9), b.FindByIndex(9))
	t.Log(s.FindByIndex(5), b.FindByIndex(5))
	t.Log(s.FindByIndex(11), b.FindByIndex(11))
}

func TestDeleteNode(t *testing.T) {
	s := NewSingleLinkedList()
	b := NewBoubleListedList()
	for i := 0; i < 3; i++ {
		s.InsertToTail(i + 1)
		b.InsertToTail(i + 1)
	}
	s.Print()
	b.Print()
	t.Log(s.Len(), b.Len())

	t.Log(s.DeleteNode(s.head.next))
	t.Log(b.DeleteNode(b.head.next))
	s.Print()
	b.Print()
	t.Log(s.Len(), b.Len())

	t.Log(s.DeleteNode(s.head.next.next))
	t.Log(b.DeleteNode(b.head.next.next))
	s.Print()
	b.Print()
	t.Log(s.Len(), b.Len())
}
