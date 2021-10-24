package main

/*
回溯 + 哈希表
检查当前节点的后继节点和随机节点创建情况
1.如果没有创建则递归的创建
2.copy完成回到当前节点进行拷贝
*/
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

var cachedNode map[*Node]*Node

func deepCopy(node *Node) *Node {
	if node == nil {
		return nil
	}
	if n, ok := cachedNode[node]; ok {
		return n
	}
	newNode := &Node{
		Val:    node.Val,
	}
	newNode.Next = deepCopy(node.Next)
	newNode.Random = deepCopy(node.Random)
	cachedNode[node] = newNode
	return newNode
}

func copyRandomList(head *Node) *Node {
	cachedNode = make(map[*Node]*Node)
	return deepCopy(head)
}

func main() {

}
