package main

import (
	"testing"
)

func TestNewListSetsProperties(t *testing.T) {
	myList := newList(48)
	if myList.head.value != 48 {
		t.Errorf("wanted %d, got %d", 48, myList.head.value)
	}
	if myList.tail == nil {
		t.Error("tail not set")
	}
	if myList.length != 1 {
		t.Errorf("wanted %d, got %d", 1, myList.length)
	}
}

func TestPush(t *testing.T) {
	myList := linkedList[int]{}
	myList.push(48)
	if myList.head.value != 48 {
		t.Errorf("wanted %d, got %d", 48, myList.head.value)
	}
	if myList.tail == nil {
		t.Error("tail not set")
	}
	if myList.length != 1 {
		t.Errorf("wanted %d, got %d", 1, myList.length)
	}
	myList.push(49)
	if myList.tail.value != 49 {
		t.Errorf("wanted %d, got %d", 49, myList.tail.value)
	}
	if myList.length != 2 {
		t.Errorf("wanted %d, got %d", 2, myList.length)
	}
}

func TestPop(t *testing.T) {
	myList := newList(48)
	myList.push(49)
	myList.push(50)
	lastNode := myList.pop()
	if lastNode.value != 50 {
		t.Error("expected ", 50, " got ", lastNode.value)
	}
	if myList.length != 2 {
		t.Error("expected length of 2. Got ", myList.length)
	}
	if myList.tail.next != nil {
		t.Error("tail should not have next value")
	}
}

func TestPopReturnsNilWhenEmpty(t *testing.T) {
	myList := linkedList[int]{}
	result := myList.pop()
	if result != nil {
		t.Error("pop should return nil for empty list")
	}
}

func TestPopWorksWithOneNode(t *testing.T) {
	myList := newList(48)
	result := myList.pop()
	if result.value != 48 {
		t.Error("expected 48, got ", result.value)
	}
	if myList.length != 0 {
		t.Error("list should be empty")
	}
}

func TestUnshift(t *testing.T) {
	list := newList(11)
	list.push(3)
	list.push(23)
	list.push(7)
	list.unshift(4)
	if list.head.value != 4 {
		t.Errorf("expected %d, got %d", 4, list.head.value)
	}
	if list.length != 5 {
		t.Errorf("expected length to be %d, got %d", 5, list.length)
	}
	if list.tail.value != 7 {
		t.Errorf("expected %d, got %d", 7, list.tail.value)
	}
}

func TestShift(t *testing.T) {
	list := newList(11)
	list.push(3)
	list.push(23)
	list.push(7)
	first := list.shift()
	if first.value != 11 {
		t.Errorf("expected %d, got %d", 11, list.head.value)
	}
	if list.head.value != 3 {
		t.Errorf("expected %d, got %d", 3, list.head.value)
	}
	if list.length != 3 {
		t.Errorf("expected length to be %d, got %d", 3, list.length)
	}
	if list.head.value != 3 {
		t.Errorf("expected %d, got %d", 3, list.head.value)
	}
}

func TestShiftWithOneElement(t *testing.T) {
	list := newList(11)
	result := list.shift()
	if result.value != 11 {
		t.Errorf("expected %d, got %d", 11, result.value)
	}
	if list.tail != nil {
		t.Errorf("Empty list should have nil tail")
	}
	if list.length != 0 {
		t.Errorf("Empty list should have 0 length")
	}
}

func TestShiftWithEmptyList(t *testing.T) {
	list := linkedList[int]{}
	result := list.shift()
	if result != nil {
		t.Errorf("expected %v, got %v", nil, result)
	}
	if list.length != 0 {
		t.Error("Empty list length should be 0")
	}
}

func TestGet(t *testing.T) {
	list := newList(1)
	list.push(2)
	list.push(3)
	list.push(4)
	result := list.get(2)
	if result.value != 3 {
		t.Errorf("expected %d, got %d", 3, result.value)
	}
	result = list.get(0)
	if result.value != 1 {
		t.Errorf("expected %d, got %d", 1, result.value)
	}

	result = list.get(3)
	if result.value != 4 {
		t.Errorf("expected %d, got %d", 4, result.value)
	}

	result = list.get(4)
	if result != nil {
		t.Errorf("expected %v, got %v", nil, result)
	}
}

func TestSet(t *testing.T) {
	list := newList(1)
	list.push(2)
	list.push(3)
	list.push(4)
	result := list.set(2, 10)
	changedNode := list.get(2)
	if changedNode.value != 10 {
		t.Errorf("expected %d, got %d", 10, changedNode.value)
	}
	if result != changedNode {
		t.Errorf("expected %v to be equal to %v", result, changedNode)
	}
	result = list.set(5, 10)
	if result != nil {
		t.Errorf("setting non-existing item should return nil")
	}
}

func TestInsert(t *testing.T) {
	list := linkedList[int]{}
	for i := 1; i < 5; i++ {
		list.push(i)
	}
	list.insert(2, 5)
	if list.length != 5 {
		t.Errorf("expected length %d, got %d", 5, list.length)
	}
	newVal := list.get(2)
	if newVal.value != 5 {
		t.Errorf("expected %d, got %d", 5, newVal.value)
	}
}

func TestRemove(t *testing.T) {
	list := linkedList[int]{}
	for i := 1; i < 5; i++ {
		list.push(i)
	}
	// test remove in middle
	node := list.remove(2)
	if node.value != 3 {
		t.Errorf("expected %d, got %d", 3, node.value)
	}
	if list.length != 3 {
		t.Errorf("expected length %d, got %d", 3, list.length)
	}
	node = list.remove(2)
	if node.value != 4 {
		t.Errorf("expected %d, got %d", 4, node.value)
	}
}

func TestReverse(t *testing.T) {
	list := linkedList[int]{}
	for i := 1; i < 6; i++ {
		list.push(i)
	}
	list.reverse()
	first := list.get(0)
	if first.value != 5 {
		t.Errorf("expected %d, got %d", 5, first.value)
	}
	last := list.get(list.length - 1)
	if last.value != 1 {
		t.Errorf("expected %d, got %d", 1, last.value)
	}
}
