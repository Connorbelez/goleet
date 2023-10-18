package main

type Node struct {
	val  int
	key  int
	next *Node
	prev *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (l *LinkedList) Add(node *Node) {
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		node.prev = l.tail
		l.tail = node
	}
}

func (l *LinkedList) Remove(node *Node) int {
	val := node.key
	if node == l.head {
		l.head = l.head.next
	} else if node == l.tail {
		l.tail = l.tail.prev
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}
	return val
}

type LRUCache struct {
	capacity int
	cache    map[int]*Node
	lru      LinkedList
	size     int
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{capacity: capacity, size: 0}
	lru.cache = make(map[int]*Node)

	return lru
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.lru.Remove(node)
	this.lru.Add(node)

	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.cache[key]
	if !ok {
		var newNode = &Node{val: value, key: key}
		this.lru.Add(newNode)
		this.cache[key] = newNode
		this.size++
	} else {
		node.val = value
		this.lru.Remove(node)
		this.lru.Add(node)
	}
	if this.size > this.capacity {

		val := this.lru.Remove(this.lru.head)
		delete(this.cache, val)
		this.size--
	}

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	lRUCache := Constructor(2)
	lRUCache.Put(1, 0)       // cache is {1=1}
	lRUCache.Put(2, 2)       // cache is {1=1, 2=2}
	println(lRUCache.Get(1)) // return 1
	lRUCache.Put(3, 3)       // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	println(lRUCache.Get(2)) // returns -1 (not found)
	lRUCache.Put(4, 4)       // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	println("items: ", lRUCache.cache)
	println("Size: ", lRUCache.size)

	println(lRUCache.Get(1)) // return -1 (not found)
	println(lRUCache.Get(3)) // return 3
	println(lRUCache.Get(4)) // return 4

}
