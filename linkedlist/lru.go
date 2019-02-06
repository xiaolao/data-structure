package linkedlist

import (
	"container/list"
	"errors"
)

type LRUCache interface {
	Add(key, value interface{}) bool
	Get(key interface{}) (value interface{}, ok bool)
}
