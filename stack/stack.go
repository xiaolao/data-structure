package stack

import "fmt"

type Stack interface {
	Push(v interface{})
	Pop() interface{}
	Top() interface{}
	IsEmpty() bool
	Flush()
	Print()
}

type ArrayStack struct {
	top  int
	data []interface{}
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		top:  -1,
		data: make([]interface{}, 0, 32),
	}
}

func (this *ArrayStack) IsEmpty() bool {
	if this.top < 0 {
		return true
	}
	return false
}

func (this *ArrayStack) Push(v interface{}) {
	if this.top < 0 {
		this.top = 0
	} else {
		this.top += 1
	}
	if this.top > len(this.data)-1 {
		this.data = append(this.data, v)
	} else {
		this.data[this.top] = v
	}
}

func (this *ArrayStack) Pop() interface{} {
	if this.IsEmpty() {
		return nil
	}
	v := this.data[this.top]
	this.top -= 1
	return v
}

func (this *ArrayStack) Top() interface{} {
	if this.IsEmpty() {
		return nil
	}
	return this.data[this.top]
}

func (this *ArrayStack) Flush() {
	this.top = -1
}

func (this *ArrayStack) Print() {
	if this.IsEmpty() {
		fmt.Println("empty stack")
	} else {
		for i := this.top; i >= 0; i-- {
			fmt.Println(this.data[i])
		}
	}
}

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
