package main

func main() {
	queue := new(Queue)
	// Preallocate memory for 10 items
	queue.create(10)
	// Exercising the queue
	queue.enqueue("abc", "123")
	queue.enqueue("def", "456")
	queue.enqueue("ghi", "789")

	queue.show()
	item := queue.peek()
	println("Peeked at item")
	queue.print_item(item)

	item = queue.pop()
	println("Popped item")
	queue.print_item(item)
	queue.show()
	println("Length of queue is", queue.length())

	queue.empty()
	if queue.isEmpty() {
		println("Que is empty")
	}
	println("Length of queue is", queue.length())
}
