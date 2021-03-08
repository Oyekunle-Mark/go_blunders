package main

import "bytes"

type MyHashMap struct {
	bucket [1000000]int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {

}

func (this *MyHashMap) Hash(key int) int {

}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {

}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {

}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {

}

func djb2Hash(buf *bytes.Buffer) uint  {
	var h uint = 5381
	for _, r := range buf.Bytes() {
		h = (h << 5) + h + uint(r)
	}

	return h
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
