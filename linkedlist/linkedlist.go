package linkedlist

import (
	"fmt"
	"strconv"
)

type node struct {
	number int
	next   *node
}

// LinkedList ll
type LinkedList struct {
	head *node
}

func (list *LinkedList) Add(n int) {
	if list.head == nil {
		list.head = &node{number: n}
		return
	}

	target := list.head
	for target.next != nil {
		target = target.next
	}

	target.next = &node{number: n}
}

func (list *LinkedList) Prepend(n int) {
	if list.head == nil {
		list.head = &node{number: n}
		return
	}

	node := &node{number: n, next: list.head}
	list.head = node
}

func (list *LinkedList) Remove(n int) {

	if list.head == nil {
		return
	}

	if list.head.number == n {
		list.head = list.head.next
		return
	}

	prev, cur := list.head, list.head.next
	for cur != nil {
		if cur.number == n {
			prev.next = cur.next
		}
		prev, cur = cur, cur.next
	}
}

func (list *LinkedList) Reverse() {
	var prev, next *node
	cur := list.head

	for cur != nil {
		next = cur.next
		cur.next = prev

		prev = cur
		cur = next
	}

	list.head = prev
}

func (list *LinkedList) ShiftRight() {

	if list.head == nil || list.head.next == nil {
		return
	}

	cur := list.head
	for cur.next.next != nil {
		cur = cur.next
	}

	temp := cur.next
	cur.next = nil

	temp.next = list.head
	list.head = temp
}

func (list *LinkedList) ShiftLeft() {

	if list.head == nil || list.head.next == nil {
		return
	}

	cur := list.head
	for cur.next != nil {
		cur = cur.next
	}

	temp := list.head.next
	cur.next = list.head
	cur.next.next = nil
	list.head = temp
}

func (list *LinkedList) Print() {

	target := list.head
	for target != nil {
		fmt.Print(strconv.Itoa(target.number) + " ")
		target = target.next
	}
	fmt.Print("\n")
}
