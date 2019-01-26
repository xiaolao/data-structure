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

// 定义单链表结构体
type SingleLinkedList struct {
	head   *SingleListNode
	length uint
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
	newNode.next = p.next
	p.next = newNode
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
	if p == nil || p == this.head {
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

// 获取链表长度
func (this *SingleLinkedList) Len() uint {
	return this.length
}

// 单链表反转
func (this *SingleLinkedList) Reverse() {
	if this.head == nil || this.head.next == nil || this.head.next.next == nil {
		return
	}
	cur := this.head.next
	var prev *SingleListNode = nil
	for cur != nil {
		tmp := cur.next
		cur.next = prev
		prev = cur
		cur = tmp
	}
	this.head.next = prev
}

// 判断单链表是否有环
func (this *SingleLinkedList) HasCycle() bool {
	// 使用快慢两个指针slow,fast，若有环slow,fast比在环内相遇
	if this.head == nil {
		return false
	}
	slow := this.head
	fast := this.head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			return true
		}
	}
	return false
}

// 删除单链表倒数第N个节点
func (this *SingleLinkedList) DeleteBottomN(n int) bool {
	// 使用快慢两个指针，先让fast走n步之后slow跟上
	// 当fast到达链表尾部时slow刚好为要删节点的前驱节点
	if n <= 0 || this.head == nil || this.head.next == nil {
		return false
	}
	fast := this.head
	for i := 1; i <= n && fast != nil; i++ {
		fast = fast.next
	}
	if fast == nil {
		return false
	}
	slow := this.head
	for fast.next != nil {
		slow = slow.next
		fast = fast.next
	}
	slow.next = slow.next.next
	return true
}

// 找到单链表的中间节点
func (this *SingleLinkedList) FindMiddleNode() *SingleListNode {
	// 使用快慢两个节点，fast一次走两步，slow一次走一步
	// 当fast到达链表尾部时，slow的位置即是单链表的中间节点
	if this.head == nil || this.head.next == nil {
		return nil
	}
	if this.head.next.next == nil {
		return this.head.next
	}
	slow := this.head
	fast := this.head
	for fast != nil && slow != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

// 定义双链表
type BoubleLinkedList struct {
	head   *BoubleListNode
	length uint
}

// 初始化一个带有头结点的双链表
func NewBoubleListedList() *BoubleLinkedList {
	return &BoubleLinkedList{NewBoubleListNode(0), 0}
}

// 在双链表指定节点后插入新的节点
func (this *BoubleLinkedList) InsertAfter(p *BoubleListNode, v interface{}) bool {
	if p == nil {
		return false
	}
	newNode := NewBoubleListNode(v)
	newNode.prev = p
	if p.next != nil {
		newNode.next = p.next
		p.next.prev = newNode
	}
	p.next = newNode
	this.length++
	return true
}

// 在双链表头部插入新节点
func (this *BoubleLinkedList) InsertToHead(v interface{}) bool {
	return this.InsertAfter(this.head, v)
}

// 在双链表尾部插入新节点
func (this *BoubleLinkedList) InsertToTail(v interface{}) bool {
	cur := this.head
	for cur.next != nil {
		cur = cur.next
	}
	return this.InsertAfter(cur, v)
}

// 在双链表指定节点前插入新的节点
func (this *BoubleLinkedList) InsertBefore(p *BoubleListNode, v interface{}) bool {
	if p == nil || p == this.head {
		return false
	}
	newNode := NewBoubleListNode(v)
	newNode.next = p
	newNode.prev = p.prev
	p.prev.next = newNode
	p.prev = newNode
	return true
}

// 根据下标获取双链表的节点
func (this *BoubleLinkedList) FindByIndex(index uint) *BoubleListNode {
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

// 删除双链表的一个节点
func (this *BoubleLinkedList) DeleteNode(p *BoubleListNode) bool {
	if p == nil || p == this.head {
		return false
	}
	if p.next != nil {
		p.next.prev = p.prev
	}
	p.prev.next = p.next
	p = nil
	this.length--
	return true
}

// 打印双链表
func (this *BoubleLinkedList) Print() {
	format := ""
	cur := this.head.next
	for cur != nil {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if cur != nil {
			format += "<->"
		}
	}
	fmt.Println(format)
}

// 获取双链表长度
func (this *BoubleLinkedList) Len() uint {
	return this.length
}
