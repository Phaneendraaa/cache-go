package main

import "fmt"

type Node struct {
	Value string
	Prev  *Node
	Next  *Node
}

type Queue struct {
	Head *Node
	Tail *Node
}

func NewQueue() *Queue {
	head := &Node{"Head", nil, nil}
	tail := &Node{"Tail", head, nil}
	head.Next = tail
	return &Queue{head, tail}
}

type Cache struct {
	Queue   *Queue
	HashMap *Hash
}

func NewHash() *Hash {
	return &Hash{Hashmap: make(map[string]*Node)}
}

type Hash struct {
	Hashmap map[string]*Node
}

func NewCache() *Cache {
	return &Cache{NewQueue(), NewHash()}
}

func (c *Cache) Put(str string, max int) {

	node, ok := c.HashMap.Hashmap[str]
	if ok {
		fmt.Println("it is present", node)
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
		delete(c.HashMap.Hashmap, str)
		fmt.Println("Deleted and trying to add ", str)
		c.Put(str, max)

	} else {
		if len(c.HashMap.Hashmap) >= max {
			c.LengthExceed()
		}
		temp := c.Queue.Head.Next
		c.Queue.Head.Next = &Node{str, c.Queue.Head, temp}
		temp.Prev = c.Queue.Head.Next
		c.HashMap.Hashmap[str] = c.Queue.Head.Next
	}
}

func (c *Cache) LengthExceed() {
	str := c.Queue.Tail.Prev.Value
	temp := c.Queue.Tail.Prev.Prev
	c.Queue.Tail.Prev = c.Queue.Tail.Prev.Prev
	temp.Next = c.Queue.Tail
	 
	delete(c.HashMap.Hashmap, str)
	fmt.Println("cache exceeded deleting ",str)
}

func main() {
	var max int
	fmt.Println("Enter the max size of the cache")
	fmt.Scanln(&max)
	cache := NewCache()
	var inp string
	for true{
		temp := cache.Queue.Head
		fmt.Println("")
		fmt.Println("Enter your input to cache ")
		fmt.Scanln(&inp)
		
		cache.Put(inp,max)
		fmt.Println("")
		fmt.Println("Cache Queue =>")
		
		for temp != nil {
			if temp.Value != "Head" && temp.Value != "Tail" {

			
				fmt.Print(temp.Value, " ->  ")
			}
			temp = temp.Next
		}
		
		fmt.Println("")
		fmt.Println("--Hash map--")
		for key, value := range cache.HashMap.Hashmap {
			fmt.Println(key, "-", value)
		}
	}
	
}
