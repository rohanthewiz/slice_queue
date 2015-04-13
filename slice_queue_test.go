package main

import "testing"

func TestSliceQueue( t *testing.T) {
	// Test create and allocate
	queue := new(Queue)
	if queue == nil {
		t.Error("queue should not be nil")
	}
	queue.create(10) // preallocate memory for 10 items
	if c := cap(queue.Que); c != 10 {
		t.Error("The capacity of queue should be 10, got ", c)
	}

	// Test enqueue
	queue.enqueue("abc", "123")
	queue.enqueue("def", "456")
	queue.enqueue("ghi", "789")
	if l := queue.length(); l != 3 {
		t.Error("Queue length should be 3, got ", l)
	}

	// Test peek
	item := queue.peek()
	if item["abc"] != "123" {
		t.Error("Peek failed. Correct map not returned")
	}

	// Test pop
	item = queue.pop()
	if item["abc"] != "123" {
		t.Error("Peek failed. Correct map not returned")
	}
	if queue.length() != 2 {
		t.Error("Length of queue should be 2 after pop")
	}

	// Test Empty
	if queue.isEmpty() {
		t.Error("Queue should not be empty when it has 2 items")
	}
	queue.empty()
	if !queue.isEmpty() {
		t.Error("Queue should now be empty")
	}
}
