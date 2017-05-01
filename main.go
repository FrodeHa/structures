package main

import (
	"fmt"
	"structures/binarytree"
	"structures/linkedlist"
	"structures/queue"
	"structures/stack"
)

func main() {
	// listTest()
	// treeTest()
	// stackTest()
	queueTest()
}

func queueTest() {
	queue := new(queue.Queue)

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)

	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())

}

func stackTest() {
	stack := new(stack.Stack)

	r, err := stack.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	r, err = stack.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

	r, err = stack.Peek()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

	r, err = stack.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

	r, err = stack.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

	r, err = stack.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}
}

func treeTest() {
	tree := new(binarytree.BinaryTree)
	tree.Add(5)
	tree.Add(2)
	tree.Add(3)
	tree.Add(-4)
	tree.Add(12)
	tree.Add(9)
	tree.Add(21)
	tree.Add(19)
	tree.Add(25)

	tree.Print()
	tree.Remove(12)

	tree.Print()

	tree.Remove(5)

	fmt.Println(tree.Find(3))
	fmt.Println(tree.Find(-4))
	fmt.Println(tree.Find(9))
	fmt.Println(tree.Find(11))

	tree.Print()

}

func listTest() {
	list := new(linkedlist.LinkedList)
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)
	list.Add(6)
	list.Print()

	list.Remove(3)
	list.Remove(1)
	list.Remove(5)

	list.Print()

	list.Prepend(10)
	list.Print()

	list.Reverse()
	list.Print()

	list.ShiftLeft()
	list.Print()
	list.ShiftRight()
	list.ShiftRight()
	list.Print()
}
