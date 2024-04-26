package queue

import "fmt"

// DynamicLinkQueue Dynamic Queue
type DynamicLinkQueue struct {
	length   int        // queue length
	capacity int        // queue capacity
	data     *queueData // a point to a linked list with header node, the list storage queue element
	tail     *queueData // a point to queue tail
}

// queueData use for linking queue element
type queueData struct {
	elem int        // queue element
	next *queueData // use for linking queueDate
}

// NewDynamicLinkQueue Construct DynamicLinkQueue object
func NewDynamicLinkQueue(cap int) *DynamicLinkQueue {

	if cap <= 0 { // if capacity parameter is illegal, set the capacity into zero
		cap = 0
	}

	q := new(DynamicLinkQueue)
	q.data = new(queueData) // allocate linked list header node
	q.tail = q.data
	q.capacity = cap

	return q
}

// Len return DynamicLinkQueue length
func (q *DynamicLinkQueue) Len() int {
	return q.length
}

// Cap return DynamicLinkQueue capacity
func (q *DynamicLinkQueue) Cap() int {
	return q.capacity
}

// ExpendOrShrink expend queue capacity
func (q *DynamicLinkQueue) ExpendOrShrink(newCap int) error {
	if newCap < 0 {
		return fmt.Errorf("newCap must be greater than zero")
	}

	if newCap < q.length {
		return fmt.Errorf("newCap is bigger than queue length")
	}

	q.capacity = newCap
	return nil
}

// EnQueue enqueue elem into queue
func (q *DynamicLinkQueue) EnQueue(elem int) error {

	if q.length == q.capacity {
		return fmt.Errorf("enqueue queue element: %d failed, the queue is full", elem)
	}

	// add a element into queue tail
	q.tail.next = new(queueData)
	q.tail = q.tail.next
	q.tail.elem = elem
	q.length++

	return nil
}

// EnQueueList enqueue the elements of elem list into queue
func (q *DynamicLinkQueue) EnQueueList(elem ...int) error {
	// 该函数为扩展非核心方法，调用方可以调用EnQueue()方法自己实现相同功能
	// 由于是非核心方法，且该函数目前仅用于演示目的，这里为了图方便，直接使用了golang已经封装好的slice
	if len(elem) > q.capacity-q.length {
		return fmt.Errorf("enqueue queue element: %v failed, the queue capacity is not enough", elem)
	}

	for i := 0; i < len(elem); i++ {
		_ = q.EnQueue(elem[i])
	}

	return nil
}

// DeQueue dequeue a element from the queue
func (q *DynamicLinkQueue) DeQueue() (int, error) {

	if q.length == 0 {
		return 0, fmt.Errorf("the queue is empty")
	}

	result := q.data.next.elem
	q.data.next = q.data.next.next
	q.length--

	return result, nil
}

// DeQueueIntoArray dequeue count elements
func (q *DynamicLinkQueue) DeQueueIntoArray(count int) ([]int, error) {
	// 该函数为扩展非核心方法，调用方可以调用EnQueue()方法自己实现相同功能
	// 由于是非核心方法，且该函数目前仅用于演示目的，这里为了图方便，直接使用了golang已经封装好的slice
	if count < 0 {
		return []int{}, fmt.Errorf("DeQueueIntoArray count must be greater than zero")
	}

	if count > q.capacity {
		return []int{}, fmt.Errorf("the length of the queue is less than need")
	}

	result := make([]int, 0, count)

	for i := 0; i < count; i++ {
		result = append(result, q.data.next.elem)
		q.data.next = q.data.next.next
		q.length--
	}

	return result, nil
}
