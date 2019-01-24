package linkedlist

import "fmt"

// 单链表节点
type SingleListNode struct {
	next  *SingleListNode
	value interface{}
}

// 生成一个新的单链表节点
func NewSingleListNode(v interface{}) *SingleListNode {
	// 该节点为哨兵节点
	return &SingleListNode{nil, v}
}

// 获取下一个单链表节点
func (this *SingleListNode) GetNext() *SingleListNode {
	return this.next
}

// 获取当前节点的值
func (this *SingleListNode) GetValue() interface{} {
	return this.value
}

// 定义单链表结构体
type SingleLinkedList struct {
	head   *SingleListNode
	length uint
}

// 双链表节点
type BoubleListNode struct {
	next  *BoubleListNode
	prev  *BoubleListNode
	value interface{}
}

// 生成一个新的双链表节点
func NewBoubleListNode(v interface{}) *BoubleListNode {
	return &BoubleListNode{nil, nil, v}
}

// 获取下一个双链表节点
func (this *BoubleListNode) GetNext() *BoubleListNode {
	return this.next
}

// 获取前一个双链表节点
func (this *BoubleListNode) GetPrev() *BoubleListNode {
	return this.prev
}

// 获取双链表节点的值
func (this *BoubleListNode) GetValue() interface{} {
	return this.value
}

// 初始化一个单链表，返回带哨兵节点的单链表
func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{NewSingleListNode(0), 0}
}

// 在一个节点后插入一个节点到单链表，时间复杂度为o(1)
func (this *SingleLinkedList) InsertAfter(p *SingleListNode, v interface{}) bool {
	if nil == p {
		return false
	}
	newNode := NewSingleListNode(v)
	oldNode := p.next
	p.next = newNode
	newNode.next = oldNode
	this.length++
	return true
}

// 在一个节点前插入一个节点到单链表，时间复杂度为o(n)
func (this *SingleLinkedList) InsertBefore(p *SingleListNode, v interface{}) bool {
	if nil == p || p == this.head {
		return false
	}
	prev := this.head
	cur := this.head.next
	for nil != cur {
		if cur == p {
			break
		}
		prev = cur
		cur = cur.next

	}
	if cur == nil {
		return false
	}
	newNode := NewSingleListNode(v)
	prev.next = newNode
	newNode.next = cur
	this.length++
	return true
}

// 在单链表头部插入节点
func (this *SingleLinkedList) InsertToHead(v interface{}) bool {
	return this.InsertAfter(this.head, v)
}

// 在单链表尾部插入节点
func (this *SingleLinkedList) InsertToTail(v interface{}) bool {
	cur := this.head
	for cur.next != nil {
		cur = cur.next
	}
	return this.InsertAfter(cur, v)
}

// 根据索引来查找节点
func (this *SingleLinkedList) FindByIndex(index uint) *SingleListNode {
	if index >= this.length {
		return nil
	}
	cur := this.head.next
	var i uint = 0
	for ; i < index; i++ {
		cur = cur.next
	}
	return cur
}

// 删除指定节点
func (this *SingleLinkedList) DeleteNode(p *SingleListNode) bool {
	if p == nil {
		return false
	}
	prev := this.head
	cur := this.head.next
	for cur != nil {
		if cur == p {
			break
		}
		prev = cur
		cur = cur.next
	}
	if cur == nil {
		return false
	}
	prev.next = p.next
	p = nil
	this.length--
	return true
}

// 打印出单链表
func (this *SingleLinkedList) Print() {
	format := ""
	cur := this.head.next
	for cur != nil {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if cur != nil {
			format += "->"
		}
	}
	fmt.Println(format)
}
