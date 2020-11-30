package main

import (
	"errors"
	"fmt"
)

type Node struct {
	PreviousNode *Node

	NextNode *Node

	Key string

	Value string
}

type LRUCache struct {
	HeadNode *Node

	TailNode *Node

	NodeMap map[string]*Node

	Capacity int
}

func (lruCache LRUCache) New(capacity int) LRUCache {
	headNode := Node{}
	tailNode := Node{
		NextNode: &headNode,
	}

	headNode.PreviousNode = &tailNode

	newLRUCache := LRUCache{
		HeadNode: &headNode,
		TailNode: &tailNode,
		Capacity: capacity,
		NodeMap:  make(map[string]*Node),
	}

	return newLRUCache
}

func (lruCache *LRUCache) moveToHead(node *Node) {
	headNode := lruCache.HeadNode
	currentHeadNode := headNode.PreviousNode
	currentHeadNode.NextNode = node

	headNode.PreviousNode = node
	node.NextNode = lruCache.HeadNode
	node.PreviousNode = currentHeadNode
}

func (lruCache *LRUCache) move(node *Node) {
	node.PreviousNode.NextNode = node.NextNode
	node.NextNode.PreviousNode = node.PreviousNode
}

func (lruCache *LRUCache) removeTail() {
	currentTailNode := lruCache.TailNode.NextNode
	lruCache.TailNode.NextNode = currentTailNode.NextNode
	lruCache.TailNode.NextNode.PreviousNode = lruCache.TailNode

	delete(lruCache.NodeMap, currentTailNode.Key)
}

func (lruCache LRUCache) isFull() bool {
	return len(lruCache.NodeMap) == lruCache.Capacity
}

func (lruCache *LRUCache) get(key string) (string, error) {
	node := lruCache.NodeMap[key]
	if node == nil {
		return "", errors.New("Not found")
	}

	lruCache.move(node)
	lruCache.moveToHead(node)

	return node.Value, nil
}

func (lruCache *LRUCache) insert(key string, value string) {
	if lruCache.isFull() {
		lruCache.removeTail()
	}

	newNode := Node{
		Key:   key,
		Value: value,
	}

	lruCache.NodeMap[key] = &newNode
	lruCache.moveToHead(&newNode)
}

func main() {
	capacity := 2
	lruCache := (LRUCache{}).New(capacity)
	lruCache.insert("key1", "value1")
	lruCache.insert("key2", "value2")

	val1, _ := lruCache.get("key1")
	fmt.Println(val1)

	lruCache.insert("key3", "value3")

	val2, _ := lruCache.get("key2")
	val3, _ := lruCache.get("key3")
	fmt.Println(val2) //should be nil
	fmt.Println(val3)
}
