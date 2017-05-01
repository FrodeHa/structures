package binarytree

import (
	"fmt"
)

type node struct {
	number int
	left   *node
	right  *node
}

func (n *node) print() {
	fmt.Printf("%d", n.number)

	if n.left != nil {
		fmt.Print("(")
		n.left.print()
		fmt.Print(")")

	}

	if n.right != nil {
		fmt.Print("(")
		n.right.print()
		fmt.Print(")")
	}
}

type BinaryTree struct {
	root *node
	size int
}

func (b *BinaryTree) Add(n int) {

	if b.root == nil {
		b.root = &node{number: n}
		return
	}

	insert(b.root, n)
	b.size++
}

func (b *BinaryTree) Find(n int) bool {
	return find(b.root, n)
}

func find(n *node, number int) bool {
	if n == nil {
		return false
	}

	if n.number == number {
		return true
	}

	if n.number > number {
		return find(n.left, number)
	}

	return find(n.right, number)
}

func insert(n *node, number int) {

	if n.number > number {
		if n.left == nil {
			n.left = &node{number: number}
		} else {
			insert(n.left, number)
		}
	} else {
		if n.right == nil {
			n.right = &node{number: number}
		} else {
			insert(n.right, number)
		}
	}
}

func (b *BinaryTree) Remove(number int) {
	if b.root == nil {
		return
	}

	if b.root.number == number {
		dummy := &node{number: 0, left: b.root}
		remove(b.root, dummy, number)
		b.root = dummy.left
	} else {
		remove(b.root, nil, number)
	}

	b.size--
}

func remove(n *node, prev *node, number int) {

	if n == nil {
		return
	}

	if n.number > number {
		remove(n.left, n, number)
	} else if n.number < number {
		remove(n.right, n, number)
	} else {
		if n.left != nil && n.right != nil { // Two children
			smallest, parent := n.right, n.right

			for smallest.left != nil {
				parent = smallest
				smallest = smallest.left
			}

			n.number = smallest.number
			if n.right == smallest {
				n.right = nil
			} else {
				parent.left = nil
			}
		} else if n.left != nil || n.right != nil { // One child
			c := nonNil(n.left, n.right)

			if prev.left == n {
				prev.left = c
			} else {
				prev.right = c
			}
		} else if n.left == nil && n.right == nil { // No child
			if prev.left == n {
				prev.left = nil
			} else {
				prev.right = nil
			}
		}
	}
}

func nonNil(n1, n2 *node) *node {
	if n1 != nil {
		return n1
	}
	return n2
}

func (b *BinaryTree) Print() {

	b.root.print()
	fmt.Println()
}
