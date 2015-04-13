package main
import(
	"fmt"
	"reflect"
)

// Implementing a FIFO queue with slices
type Queue struct {
	Que [](map[string]string) // an array of hashes
}

func (q * Queue) create(capacity int32) {
	q.Que = make([](map[string]string), 0, capacity) //length of zero, but capacity of 10 (saves on memory allocs)
}

func (q * Queue) show() {
	println("Current state of the queue")
	for _, item := range(q.Que) {
		q.print_item(item)
	}
}

func (q * Queue) enqueue(key string, payload string) {
	q.Que = append(q.Que, map[string]string{key : payload})
}

// Who's next?
func (q * Queue) peek() map[string]string {
	return q.Que[0]
}

// Get next and remove from Que
func (q * Queue) pop() map[string]string {
	ret := q.Que[0]
	q.Que = q.Que[1:]  // drop first element
	return ret
}

func (q * Queue) isEmpty() bool {
	return len(q.Que) <= 0
}

func (q * Queue) empty() bool {
	q.Que = q.Que[0:0]
	return len(q.Que) <= 0
}

func (q * Queue) length() int {
	return len(q.Que)
}

// Util function
func (q * Queue) print_item(item map[string]string) {
	if fmt.Sprintf("%v", reflect.TypeOf(item)) == "map[string]string" {
		for key, val := range item {
			fmt.Printf("%s: %s\n", key, val)
		}
	}
}
