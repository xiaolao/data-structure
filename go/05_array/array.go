package array

import (
	"errors"
	"fmt"
)

// 数组方法约束
type ArrayInterface interface {
	Len() int
	Find(index uint) interface{}
	Insert(index uint, v interface{}) bool
	Delete(index uint) interface{}
}

type Array struct {
	data   []int
	length uint
}

// allocate an array
func NewArray(capacity uint) *Array {
	if capacity == 0 {
		return nil
	} else {
		return &Array{
			length: 0,
			data:   make([]int, capacity, capacity),
		}
	}
}

// 获取数组长度
func (this *Array) Len() int {
	return this.length
}

// 判断下表是否越界
func (this *Array) isIndexOutOfRange(index uint) bool {
	if index >= uint(cap(this.data)) {
		return true
	}
	return false
}

// 通过下标查找元素,范围为[0, n-1]
func (this *Array) Find(index uint) interface{} {
	if this.isIndexOutOfRange(index) {
		panic("index out of range")
	}
	return this.data[index]
}

// 通过下标向数组中插入元素
func (this *Array) Insert(index uint, v interface{}) bool {
	if this.length == uint(cap(this.data)) {
		panic("full array")
	}
	if index != this.length && this.isIndexOutOfRange(index) {
		panic("index out of range")
	}
	for i := this.length; i > index; i-- {
		this.data[i] = this.data[i-1]
	}
	this.data[index] = v
	this.length++
	return nil
}

// 删除索引index上的值
func (this *Array) Delete(index uint) interface{} {
	if this.isIndexOutOfRange(index) {
		panic("out of index range")
	}
	v := this.data[index]
	for i := index; i < this.Len()-1; i++ {
		this.data[i] = this.data[i+1]
	}
	this.length--
	return v
}

// 打印数列
func (this *Array) Print() {
	var format string
	for i := uint(0); i < this.Len(); i++ {
		format += fmt.Sprintf("|%+v", this.data[i])
	}
	fmt.Println(format)
}
