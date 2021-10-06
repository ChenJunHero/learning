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
	cache := NewLRUCache(10)
	tests := [][]int{
		{10},
{10,13},
{3,17},
{6,11},
{10,5},
{9,10},
{13},
{2,19},
{2},
{3},
{5,25},
{8},
{9,22},
{5,5},
{1,30},
{11},
{9,12},
{7},
{5},
{8},
{9},
{4,30},
{9,3},
{9},
{10},
{10},
{6,14},
{3,1},
{3},
{10,11},
{8},
{2,14},
{1},
{5},
{4},
{11,4},
{12,24},
{5,18},
{13},
{7,23},
{8},
{12},
{3,27},
{2,12},
{5},
{2,9},
{13,4},
{8,18},
{1,7},
{6},
{9,29},
{8,21},
{5},
{6,30},
{1,12},
{10},
{4,15},
{7,22},
{11,26},
{8,17},
{9,29},
{5},
{3,4},
{11,30},
{12},
{4,29},
{3},
{9},
{6},
{3,4},
{1},
{10},
{3,29},
{10,28},
{1,20},
{11,13},
{3},
{3,12},
{3,8},
{10,9},
{3,26},
{8},
{7},
{5},
{13,17},
{2,27},
{11,15},
{12},
{9,19},
{2,15},
{3,16},
{1},
{12,17},
{9,1},
{6,19},
{4},
{5},
{5},
{8,1},
{11,7},
{5,2},
{9,28},
{1},
{2,2},
{7,4},
{4,22},
{7,24},
{9,26},
{13,28},
{11,26},
	}
	for i, tt := range tests {
		if len(tt) == 2 {
			cache.Put(tt[0], tt[1])
		} else {
			if i == 49 {
				fmt.Println(cache.Get(tt[0]), tt[0])
				continue
			}
			cache.Get(tt[0])
		}
	}
	// fmt.Println(cache.Get(2))
	// head := cache.head
	// for i := 0; i <20; i++ {
	// 	if head == nil {
	// 		continue
	// 	}
	// 	fmt.Println(i, head)
	// 	head = head.next
	// }
	// fmt.Println(cache.Get(6))
}
