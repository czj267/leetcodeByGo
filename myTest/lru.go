package main

import (
	"fmt"
)

func main() {

	op := []string{"put", "put", "put", "put", "put", "get", "put", "get", "get", "put", "get",
		"put", "put", "put", "get", "put", "get", "get", "get", "get", "put", "put", "get", "get",
		"get", "put", "put", "get", "put", "get", "put", "get", "get", "get", "put", "put", "put", "get", "put", "get", "get", "put", "put", "get", "put", "put", "put", "put", "get", "put", "put", "get", "put", "put", "get", "put", "put", "put", "put", "put", "get", "put", "put", "get", "put", "get", "get", "get", "put", "get", "get", "put", "put", "put", "put", "get", "put", "put", "put", "put", "get", "get", "get", "put", "put", "put", "get", "put", "put", "put", "get", "put", "put", "put", "get", "get", "get", "put", "put", "put", "put", "get", "put", "put", "put", "put", "put", "put", "put"}
	vals := [][]int{
		{10, 13}, {3, 17}, {6, 11}, {10, 5}, {9, 10}, {13}, {2, 19}, {2}, {3}, {5, 25}, {8},
		{9, 22}, {5, 5}, {1, 30}, {11}, {9, 12}, {7}, {5}, {8}, {9}, {4, 30}, {9, 3}, {9}, {10}, {10}, {6, 14}, {3, 1}, {3}, {10, 11}, {8}, {2, 14}, {1}, {5}, {4}, {11, 4}, {12, 24}, {5, 18}, {13}, {7, 23}, {8}, {12}, {3, 27}, {2, 12}, {5}, {2, 9}, {13, 4}, {8, 18}, {1, 7}, {6}, {9, 29}, {8, 21}, {5}, {6, 30}, {1, 12}, {10}, {4, 15}, {7, 22}, {11, 26}, {8, 17}, {9, 29}, {5}, {3, 4}, {11, 30}, {12}, {4, 29}, {3}, {9}, {6}, {3, 4}, {1}, {10}, {3, 29}, {10, 28}, {1, 20}, {11, 13}, {3}, {3, 12}, {3, 8}, {10, 9}, {3, 26}, {8}, {7}, {5}, {13, 17}, {2, 27}, {11, 15}, {12}, {9, 19}, {2, 15}, {3, 16}, {1}, {12, 17}, {9, 1}, {6, 19}, {4}, {5}, {5}, {8, 1}, {11, 7}, {5, 2}, {9, 28}, {1}, {2, 2}, {7, 4}, {4, 22}, {7, 24}, {9, 26}, {13, 28}, {11, 26},
	}
	cache := Constructor(10)
	for k, o := range op {
		if o == "put" {
			cache.Put(vals[k][0], vals[k][1])
		} else {
			fmt.Println(cache.Get(vals[k][0]))
		}
	}
}

type LRUCache struct {
	keys     []int       //保存最近最常访问的key，越往后越经常访问
	data     map[int]int //缓存数据
	capacity int         //缓存容量
	len      int         //最近最长访问的key的数量
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity: capacity, data: map[int]int{},}
}

func (this *LRUCache) Get(key int) int {
	val, ok := this.data[key]
	if !ok {
		return -1
	}
	//遍历最近最常访问的key数组，如果找到key，那么删除，再加到最后面
	site, ok := inSlice(key, this.keys)
	if ok {
		this.keys = append(this.keys[0:site], this.keys[site+1:]...)
		this.len--
	}
	//如果最近最常访问的数量超过了容量，那么删除第一个key
	if this.len == this.capacity {
		this.keys = this.keys[1:]
	}
	this.len++
	this.keys = append(this.keys, key)
	//fmt.Printf("keys: %v\n", this.keys)
	return val
}

func (this *LRUCache) Put(key int, value int) {
	this.len++
	//遍历最近最常访问的key数组，如果找到key，那么删除，再加到最后面
	//因为添加的key也算是一次访问，故需要如此
	site, ok := inSlice(key, this.keys)
	if ok {
		this.keys = append(this.keys[0:site], this.keys[site+1:]...)
		this.len--
	}

	//如果长度大于容量则需要删除数据，
	//同时减少最近最常访问的key
	if this.len > this.capacity {
		delKey := this.keys[0]
		this.keys = this.keys[1:]
		delete(this.data, delKey)
		this.len--
	}
	//添加数据
	this.keys = append(this.keys, key)
	this.data[key] = value
	//fmt.Printf("put %v %v %v\n", this.data, key, value)
	//fmt.Printf("keys: %v \n", this.keys)
}

/**
判断一个元素是否在数组
*/
func inSlice(val int, slice []int) (int, bool) {
	for key, v := range slice {
		if v == val {
			return key, true
		}
	}
	return -1, false
}
