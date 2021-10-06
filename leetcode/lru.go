package main
/* 
least recently used 最久访问的最先淘汰
数据结构使用：
1.hashmap
2.双链表
思路：
1.最先访问的放到链表头部
2.达到cache长度先删掉尾部节点
*/

import "fmt"
// Node k, v节点
type Node struct {
	key   int
	value int
	pre   *Node
	next  *Node
}

// LRUCache cache 结构
type LRUCache struct {
	cap  int           // 容量
	m    map[int]*Node // cache map
	head *Node
	tail *Node
}

// NewLRUCache 构造一个lru cache
func NewLRUCache(cap int) LRUCache {
	cacheMap := make(map[int]*Node, cap)
	head := &Node{0, 0, nil, nil}
	tail := &Node{0, 0, nil, nil}
	head.next = tail
	tail.pre = head
	return LRUCache{cap: cap, m: cacheMap, head: head, tail: tail}
}

// move2Head 该节点移动到双链表的头部
func (l *LRUCache) move2Head(node *Node) {
	// 从当前位置删除该节点
	node.pre.next = node.next
	node.next.pre = node.pre

	// 把该节点移动到链表头部
	l.head.next.pre = node
	node.next = l.head.next
	l.head.next = node
	node.pre = l.head
}

// Get 从lru中取值
func (l *LRUCache) Get(key int) int {
	// 存在取值，并且把该节点移动到头部
	if node, ok := l.m[key]; ok {
		l.move2Head(node)
		return node.value
	}
	return -1
}

// Put 向lru中填充值
func (l *LRUCache) Put(key, value int) {
	if node, ok := l.m[key]; ok {
		node.value = value
		l.move2Head(node)
		return
	}

	tail := l.tail
	// 判断lru的容量, 如果容量满了，就要删掉链表末端的最久未访问的节点
	if len(l.m) > l.cap {
		delete(l.m, tail.pre.key)
		tail.pre.pre.next = tail
		tail.pre = tail.pre.pre
	}

	curNode := &Node{key, value, nil, nil}
	// 加入链表头部
	head := l.head
	curNode.pre = head
	curNode.next = head.next
	head.next.pre = curNode
	head.next = curNode
	l.m[key] = curNode
}

func main() {
}
