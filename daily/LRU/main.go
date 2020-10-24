package main

import "fmt"

type LRUCache struct {
	find     map[int]int
	cache    []int
	capacity int
	length   int
}

func Constructor(capacity int) LRUCache {
	cacheSlice := make([]int, capacity, capacity)
	findMap := make(map[int]int)

	return LRUCache{
		cache:    cacheSlice,
		find:     findMap,
		capacity: capacity,
		length:   0,
	}
}

func (this *LRUCache) Get(key int) int {
	result, ok := this.find[key]
	if !ok {
		return -1
	}

	var i int
	for i = 0; i < this.length; i++ {
		if this.cache[i] == key {
			break
		}
	}
	// 等于时无需调整
	for j := i; j < this.length-1; j++ {
		this.cache[j] = this.cache[j+1]
	}
	this.cache[this.length-1] = key

	return result
}

func (this *LRUCache) Put(key int, value int) {
	getRst := this.Get(key)
	if getRst != -1 {
		this.find[key] = value
		return
	}

	this.find[key] = value

	if this.length == this.capacity {
		delete(this.find, this.cache[0])

		for i := 0; i < this.capacity-1; i++ {
			this.cache[i] = this.cache[i+1]
		}
		this.cache[this.capacity-1] = key
	} else {
		this.cache[this.length] = key
		this.length++
	}
	return
}

// ["LRUCache","put","put","put","put","get","get","get","get","put","get","get","get","get","get"]
// [[3],[1,1],[2,2],[3,3],[4,4],[4],[3],[2],[1],[5,5],[1],[2],[3],[4],[5]]

func main() {
	cache1 := Constructor(3)
	cache1.Put(1, 1)
	cache1.Put(2, 2)
	cache1.Put(3, 3)
	cache1.Put(4, 4)

	fmt.Println(cache1.Get(4))
	fmt.Println(cache1.cache)
	fmt.Println(cache1.find)

	fmt.Println(cache1.Get(3))
	fmt.Println(cache1.cache)

	fmt.Println(cache1.Get(2))
	fmt.Println(cache1.cache)

	fmt.Println(cache1.Get(1))
	fmt.Println(cache1.cache)

	cache1.Put(5, 5)

	fmt.Println(cache1.Get(1))
	fmt.Println(cache1.Get(2))
	fmt.Println(cache1.Get(3))
	fmt.Println(cache1.Get(4))
	fmt.Println(cache1.Get(5))
}
