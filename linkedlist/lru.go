package linkedlist

// import (
// 	"container/list"
// 	"errors"
// )

type LRUCache interface {
	// add a value to cache
	Add(key, value interface{}) bool
	Get(key interface{}) (value interface{}, ok bool)
	Contains(key interface{}) (ok bool)
	Peek(key interface{}) (value interface{}, ok bool)
	Remove(key interface{}) bool
	RemoveOldest() (interface{}, interface{}, bool)
	GetOldest() (interface{}, interface{}, bool)
	Keys() []interface{}
	Len() int
	Purge()
}
