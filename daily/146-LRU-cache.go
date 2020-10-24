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