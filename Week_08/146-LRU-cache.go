// 使用双向链表，以及字典存储链表的节点以直接操纵节点优先/节点删除等

type LRUCache struct {
	capacity int
	length   int
	find     map[int]*storeVal
	head     *cacheNode
	tail     *cacheNode
}

type cacheNode struct {
	pre  *cacheNode
	next *cacheNode
	key  int
}

type storeVal struct {
	value int
	node  *cacheNode
}

func (this *LRUCache) moveFromCache(node *cacheNode) {
	node.pre.next = node.next
	node.next.pre = node.pre

	return
}

func (this *LRUCache) moveToHead(node *cacheNode) {
	node.pre = this.head
	node.next = this.head.next
	this.head.next.pre = node
	this.head.next = node
	return
}

func (this *LRUCache) deleteFromCache() {
	this.tail.pre.pre.next = this.tail
	this.tail.pre = this.tail.pre.pre
}

func Constructor(capacity int) LRUCache {
	result := LRUCache{
		capacity: capacity,
		length:   0,
	}
	result.find = make(map[int]*storeVal)

	head, tail := &cacheNode{}, &cacheNode{}

	head.next = tail
	tail.pre = head

	result.head = head
	result.tail = tail

	return result
}

func (this *LRUCache) Get(key int) int {
	if val, ok := this.find[key]; ok {
		if val.node.pre != this.head {
			this.moveFromCache(val.node)
			this.moveToHead(val.node)
		}
		return val.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if this.Get(key) != -1 {
		val, _ := this.find[key]
		val.value = value
		return
	}

	if this.length == this.capacity {
		delete(this.find, this.tail.pre.key)

		this.deleteFromCache()
		this.length--
	}

	newNode := &cacheNode{
		key:  key,
		pre:  this.tail.pre,
		next: this.tail,
	}

	this.find[key] = &storeVal{
		value: value,
		node:  newNode,
	}

	this.moveToHead(newNode)

	this.length++
}