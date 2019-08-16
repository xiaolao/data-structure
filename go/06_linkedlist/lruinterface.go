package linkedlist

type LRUInterface interface {

	// get the last use element in lru object.
	Get(key interface{}) interface{}

	// push the element in lru object.
	Put(key, value interface{})
}
