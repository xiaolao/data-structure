package linkedlist

// ListInterface the interface of a list structure.
type ListInterface interface {
	// 在某节点前插入
	InsertAfter(p *SingleListNode, v interface{}) bool
	// 在某节点后插入
	InsertBefore(p *SingleListNode, v interface{}) bool
	// 在头结点插入
	InsertToHead(v interface{}) bool
	// 在尾节点插入
	InsertToTail(v interface{}) bool
	// 通过下标查询节点
	FindByIndex(index uint) *SingleListNode
	// 删除节点元素
	DeleteNode(p *SingleListNode) bool
}
