package main

import "fmt"

type LRUCache1 struct {
	m       map[int]*Node
	nLen    int
	maxLen  int
	head    *Node
	tail    *Node
}

type Node struct {
	key     int
	value   int
	next    *Node
	pre     *Node
}

func main() {
	op := []string{"put", "put", "put", "put", "put", "get", "put", "get", "get", "put", "get",
		"put", "put", "put", "get", "put", "get", "get", "get", "get", "put", "put", "get", "get",
		"get", "put", "put", "get", "put", "get", "put", "get", "get", "get", "put", "put", "put", "get", "put", "get", "get", "put", "put", "get", "put", "put", "put", "put", "get", "put", "put", "get", "put", "put", "get", "put", "put", "put", "put", "put", "get", "put", "put", "get", "put", "get", "get", "get", "put", "get", "get", "put", "put", "put", "put", "get", "put", "put", "put", "put", "get", "get", "get", "put", "put", "put", "get", "put", "put", "put", "get", "put", "put", "put", "get", "get", "get", "put", "put", "put", "put", "get", "put", "put", "put", "put", "put", "put", "put"}
	vals := [][]int{
		{10, 13}, {3, 17}, {6, 11}, {10, 5}, {9, 10}, {13}, {2, 19}, {2}, {3}, {5, 25}, {8},
		{9, 22}, {5, 5}, {1, 30}, {11}, {9, 12}, {7}, {5}, {8}, {9}, {4, 30}, {9, 3}, {9}, {10}, {10}, {6, 14}, {3, 1}, {3}, {10, 11}, {8}, {2, 14}, {1}, {5}, {4}, {11, 4}, {12, 24}, {5, 18}, {13}, {7, 23}, {8}, {12}, {3, 27}, {2, 12}, {5}, {2, 9}, {13, 4}, {8, 18}, {1, 7}, {6}, {9, 29}, {8, 21}, {5}, {6, 30}, {1, 12}, {10}, {4, 15}, {7, 22}, {11, 26}, {8, 17}, {9, 29}, {5}, {3, 4}, {11, 30}, {12}, {4, 29}, {3}, {9}, {6}, {3, 4}, {1}, {10}, {3, 29}, {10, 28}, {1, 20}, {11, 13}, {3}, {3, 12}, {3, 8}, {10, 9}, {3, 26}, {8}, {7}, {5}, {13, 17}, {2, 27}, {11, 15}, {12}, {9, 19}, {2, 15}, {3, 16}, {1}, {12, 17}, {9, 1}, {6, 19}, {4}, {5}, {5}, {8, 1}, {11, 7}, {5, 2}, {9, 28}, {1}, {2, 2}, {7, 4}, {4, 22}, {7, 24}, {9, 26}, {13, 28}, {11, 26},
	}
	cache := Constructor1(10)
	for k, o := range op {
		if o == "put" {
			cache.Put(vals[k][0], vals[k][1])
		} else {
			fmt.Println(cache.Get(vals[k][0]))
		}
	}
}

func Constructor1(capacity int) LRUCache1 {
	tm := make(map[int]*Node, capacity)
	return LRUCache1{
		m: tm,
		nLen: 0,
		maxLen: capacity,
		head: nil,
		tail: nil,
	}
}


func (this *LRUCache1) Get(key int) int {
	if this.nLen == 0 {
		return -1
	}
	v, ok := this.m[key]
	if !ok {
		return -1
	}
	if this.head == v {
		return v.value
	}
	v.pre.next = v.next
	if v.next != nil {
		v.next.pre = v.pre
	} else {
		this.tail = v.pre
	}
	v.pre = nil
	v.next = this.head
	this.head.pre = v
	this.head = v
	return v.value
}


func (this *LRUCache1) Put(key int, value int) {
	if this.nLen == 0 {
		node := &Node{
			key:    key,
			value:  value,
			next:   nil,
			pre:    nil,
		}
		this.m[key] = node
		this.head, this.tail = node, node
		this.nLen++
		return
	}
	v, ok := this.m[key]
	if ok {
		v.value = value
		if this.head == v {
			return
		}
		v.pre.next = v.next
		if v.next != nil {
			v.next.pre = v.pre
		} else {
			this.tail = v.pre
		}
		v.pre = nil
		this.head.pre = v
		v.next = this.head
		this.head = v
		return
	}
	// 不存在
	node := &Node{
		key:    key,
		value:  value,
		next:   nil,
		pre:    nil,
	}
	this.m[key] = node
	if this.nLen < this.maxLen {
		this.head.pre = node
		node.next = this.head
		this.head = node
		this.nLen++
	} else {
		delete(this.m, this.tail.key)
		this.tail = this.tail.pre
		if this.maxLen == 1 {
			this.tail, this.head = node, node
			return
		}
		this.tail.next = nil
		this.head.pre = node
		node.next = this.head
		this.head = node
	}
}
