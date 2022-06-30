package main

type node[T any] struct {
	value T
	next  *node[T]
}

type linkedList[T any] struct {
	head   *node[T]
	tail   *node[T]
	length int
}

func newList[T any](value T) linkedList[T] {
	head := &node[T]{value: value}
	list := linkedList[T]{}
	list.head = head
	list.tail = list.head
	list.length = 1
	return list
}

// push adds value to end of list
func (l *linkedList[T]) push(value T) *linkedList[T] {
	// step one create a new node
	node := &node[T]{value: value}
	// if head/tail nil head/tail point to new node
	if l.head == nil {
		l.head = node
		l.tail = l.head
	} else {
		// 2. last node point to new node
		l.tail.next = node
		l.tail = node
	}
	l.length++
	return l
}

// pop removes last element from list O(n)
func (l *linkedList[T]) pop() *node[T] {
	if l.tail == nil {
		return nil
	}
	currNode := l.head
	prevNode := l.head
	for currNode.next != nil {
		prevNode = currNode
		currNode = currNode.next
	}
	l.tail = prevNode
	l.tail.next = nil
	l.length--
	if l.length == 0 {
		l.head = nil
		l.tail = nil
	}
	return currNode
}

func (l *linkedList[T]) unshift(value T) *linkedList[T] {
	newNode := &node[T]{value: value}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.next = l.head
		l.head = newNode
	}
	l.length++
	return l
}

func (l linkedList[T]) empty() bool {
	return l.head == nil && l.tail == nil && l.length == 0
}

// shift removes element from front of the list O(1)
func (l *linkedList[T]) shift() *node[T] {
	// ll 3, 23, 7
	// move head to next item
	if l.empty() {
		return nil
	}
	first := l.head     // save current head
	l.head = first.next // set head to current head next node
	first.next = nil    // remove node chain from previous head
	l.length--          // decrement length
	if l.length == 0 {  // if length is 0 remove tail
		l.tail = nil
	}
	return first
}

// get returns the node at given index O(n)
func (l *linkedList[T]) get(index int) *node[T] {
	if l.empty() || index < 0 || index >= l.length {
		return nil
	}
	currNode := l.head
	for i := 0; i < index; i++ {
		currNode = currNode.next
	}
	return currNode
}

func (l *linkedList[T]) set(index int, value T) *node[T] {
	node := l.get(index)
	if node == nil {
		return nil
	}
	node.value = value
	return node
}

func (l *linkedList[T]) insert(index int, value T) *linkedList[T] {
	// 1 2 3 4
	if index == 0 {
		return l.unshift(value)
	}
	if index >= l.length {
		return l.push(value)
	}
	newNode := &node[T]{value: value}
	prevNode := l.get(index - 1)
	nextNode := prevNode.next
	prevNode.next = newNode
	newNode.next = nextNode

	l.length++
	return l
}

func (l *linkedList[T]) remove(index int) *node[T] {
	if index == 0 {
		return l.shift()
	}
	if index >= l.length {
		return nil
	}
	// 1 2 3 4
	prevNode := l.get(index - 1)
	if prevNode.next == l.tail {
		return l.pop()
	}
	target := prevNode.next
	prevNode.next = target.next
	target.next = nil
	l.length--
	return target
}

func (l *linkedList[T]) reverse() *linkedList[T] {
	// 1 2 3 4
	// swap head and tail
	head := l.head  // save head head
	l.head = l.tail // point head to tail
	l.tail = head   // point tail to prev head
	var prev *node[T]
	for i := 0; i < l.length; i++ {
		next := head.next
		head.next = prev
		prev = head
		head = next
	}
	return l
}
