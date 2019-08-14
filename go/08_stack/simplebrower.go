package _8_stack

import "fmt"

// 模拟浏览器行为
type Browser struct {
	forwardStack Stack
	backStack    Stack
}

// NewBrowser 初始化浏览器
func NewBrowser() *Browser {
	return &Browser{
		forwardStack: NewArrayStack(),
		backStack:    NewLinkedListStack(),
	}
}

// CanForward 是否可以向前翻页
func (this *Browser) CanForward() bool {
	if this.forwardStack.IsEmpty() {
		return false
	}
	return true
}

// CanBack 是否可以向后翻页
func (this *Browser) CanBack() bool {
	if this.backStack.IsEmpty() {
		return false
	}
	return true
}

// Open 打开指定链接
func (this *Browser) Open(addr string) {
	fmt.Printf("Open new addr %+v\n", addr)
	this.forwardStack.Flush()
}

// PushBack 向backStack中压入链接
func (this *Browser) PushBack(addr string) {
	this.backStack.Push(addr)
}

// Forward 向前翻页
func (this *Browser) Forward() {
	if this.forwardStack.IsEmpty() {
		return
	}
	top := this.forwardStack.Pop()
	this.backStack.Push(top)
	fmt.Printf("forward to %+v\n", top)
}

// Back 向后翻页
func (this *Browser) Back() {
	if this.backStack.IsEmpty() {
		return
	}
	top := this.backStack.Pop()
	this.forwardStack.Push(top)
	fmt.Printf("back to %+v\n", top)
}
