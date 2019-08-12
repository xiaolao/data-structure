package _8_stack

import "fmt"

type node struct {
	next *node
	data interface{}
}

type LinkedListStack struct {
	topNode *node
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{nil}
}

func (this *LinkedListStack) IsEmpty() bool {
	if this.topNode == nil {
		return true
	}
	return false
}

func (this *LinkedListStack) Push(v interface{}) {
	// 链表头部做为栈顶,这样每次pop操作都是o(1)的。
	this.topNode = &node{
		data: v,
		next: this.topNode,
	}
}

func (this *LinkedListStack) Pop() interface{} {
	if this.IsEmpty() {
		return nil
	}
	v := this.topNode.data
	this.topNode = this.topNode.next
	return v
}

func (this *LinkedListStack) Top() interface{} {
	if this.IsEmpty() {
		return nil
	}
	return this.topNode.data
}

func (this *LinkedListStack) Flush() {
	this.topNode = nil
}

func (this *LinkedListStack) Print() {
	if this.IsEmpty() {
		fmt.Println("stack empty")
	} else {
		cur := this.topNode
		for cur != nil {
			fmt.Println(cur.data)
			cur = cur.next
		}
	}
}
