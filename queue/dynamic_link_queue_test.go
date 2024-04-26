package queue

import (
	"reflect"
	"testing"
)

func buildDynamicLinkQueueTestInstance(length, cap, elem int) *DynamicLinkQueue {
	result := NewDynamicLinkQueue(cap)
	for i := 0; i < length; i++ {
		result.tail.next = new(queueData)
		result.tail = result.tail.next
		result.tail.elem = elem
		result.length++
	}

	return result
}

func TestNewDynamicLinkQueue(t *testing.T) {

	tests := []struct {
		name string
		cap  int
		want *DynamicLinkQueue
	}{
		{
			name: "NewDynamicLinkQueue test 1",
			cap:  -1,
			want: &DynamicLinkQueue{
				length:   0,
				capacity: 0,
			},
		},
		{
			name: "NewDynamicLinkQueue test 2",
			cap:  0,
			want: &DynamicLinkQueue{
				length:   0,
				capacity: 0,
			},
		},
		{
			name: "NewDynamicLinkQueue test 3",
			cap:  1,
			want: &DynamicLinkQueue{
				length:   0,
				capacity: 1,
			},
		},
		{
			name: "NewDynamicLinkQueue test 4",
			cap:  5,
			want: &DynamicLinkQueue{
				length:   0,
				capacity: 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := func(got, want *DynamicLinkQueue) bool {
				if got.length != want.length {
					return false
				}
				if got.capacity != want.capacity {
					return false
				}
				return true
			}

			if got := NewDynamicLinkQueue(tt.cap); !f(got, buildDynamicLinkQueueTestInstance(tt.want.length, tt.want.capacity, 1)) {
				t.Errorf("NewDynamicLinkQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDynamicLinkQueue_EnQueue(t *testing.T) {
	type fields struct {
		length   int
		capacity int
	}

	tests := []struct {
		name    string
		fields  fields
		elem    int
		wantErr bool
	}{
		{
			name: "enqueue test1",
			fields: fields{
				length:   0,
				capacity: 0,
			},
			elem:    1,
			wantErr: true,
		},
		{
			name: "enqueue test2",
			fields: fields{
				length:   0,
				capacity: 1,
			},
			elem:    1,
			wantErr: false,
		},
		{
			name: "enqueue test3",
			fields: fields{
				length:   1,
				capacity: 1,
			},
			elem:    1,
			wantErr: true,
		},
		{
			name: "enqueue test4",
			fields: fields{
				length:   1,
				capacity: 2,
			},
			elem:    2,
			wantErr: false,
		},
		{
			name: "enqueue test5",
			fields: fields{
				length:   0,
				capacity: 5,
			},
			elem:    1,
			wantErr: false,
		},
		{
			name: "enqueue test6",
			fields: fields{
				length:   4,
				capacity: 5,
			},
			elem:    5,
			wantErr: false,
		},
		{
			name: "enqueue test7",
			fields: fields{
				length:   5,
				capacity: 5,
			},
			elem:    6,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := buildDynamicLinkQueueTestInstance(tt.fields.length, tt.fields.capacity, tt.elem)
			if err := q.EnQueue(tt.elem); (err != nil) != tt.wantErr {
				t.Errorf("EnQueue() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}

func TestDynamicLinkQueue_DeQueue(t *testing.T) {
	type fields struct {
		length int
		cap    int
	}

	tests := []struct {
		name       string
		fields     fields
		want       int
		wantLength int
		wantErr    bool
	}{
		{
			name: "dequeue test1",
			fields: fields{
				length: 0,
				cap:    0,
			},
			want:       0,
			wantLength: 0,
			wantErr:    true,
		},
		{
			name: "dequeue test2",
			fields: fields{
				length: 0,
				cap:    1,
			},
			want:       0,
			wantLength: 0,
			wantErr:    true,
		},
		{
			name: "dequeue test3",
			fields: fields{
				length: 0,
				cap:    5,
			},
			want:       0,
			wantLength: 0,
			wantErr:    true,
		},
		{
			name: "dequeue test4",
			fields: fields{
				length: 1,
				cap:    1,
			},
			want:       3,
			wantLength: 0,
			wantErr:    false,
		},
		{
			name: "dequeue test6",
			fields: fields{
				length: 1,
				cap:    5,
			},
			want:       4,
			wantLength: 0,
			wantErr:    false,
		},
		{
			name: "dequeue test7",
			fields: fields{
				length: 5,
				cap:    5,
			},
			want:       5,
			wantLength: 4,
			wantErr:    false,
		},
		{
			name: "dequeue test8",
			fields: fields{
				length: 4,
				cap:    5,
			},
			want:       6,
			wantLength: 3,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := buildDynamicLinkQueueTestInstance(tt.fields.length, tt.fields.cap, tt.want)
			got, err := q.DeQueue()
			if (err != nil) != tt.wantErr {
				t.Errorf("DeQueue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DeQueue() got = %v, want %v", got, tt.want)
			}
			if q.length != tt.wantLength {
				t.Errorf("DeQueue() queuelength = %v, wantLength %v", tt.fields.length, tt.wantLength)
			}
		})
	}
}

func TestDynamicLinkQueue_Cap(t *testing.T) {
	type fields struct {
		length int
		cap    int
	}

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "capacity test1",
			fields: fields{
				length: 0,
				cap:    0,
			},
			want: 0,
		},
		{
			name: "capacity test2",
			fields: fields{
				length: 1,
				cap:    1,
			},
			want: 1,
		},
		{
			name: "capacity test3",
			fields: fields{
				length: 0,
				cap:    5,
			},
			want: 5,
		},
		{
			name: "capacity test4",
			fields: fields{
				length: 5,
				cap:    5,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := buildDynamicLinkQueueTestInstance(tt.fields.length, tt.fields.cap, tt.want)
			if got := q.Cap(); got != tt.want {
				t.Errorf("Cap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDynamicLinkQueue_Len(t *testing.T) {
	type fields struct {
		length int
		cap    int
	}

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "length test1",
			fields: fields{
				length: 0,
				cap:    0,
			},
			want: 0,
		},
		{
			name: "length test2",
			fields: fields{
				length: 1,
				cap:    1,
			},
			want: 1,
		},
		{
			name: "length test3",
			fields: fields{
				length: 0,
				cap:    5,
			},
			want: 0,
		},
		{
			name: "length test4",
			fields: fields{
				length: 5,
				cap:    5,
			},
			want: 5,
		},
		{
			name: "length test4",
			fields: fields{
				length: 2,
				cap:    5,
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := buildDynamicLinkQueueTestInstance(tt.fields.length, tt.fields.cap, 1)
			if got := q.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDynamicLinkQueue_ExpendOrShrink(t *testing.T) {
	type fields struct {
		length int
		cap    int
		newCap int
	}

	tests := []struct {
		name       string
		fields     fields
		wantNewCap int
		wantErr    bool
	}{
		{
			name: "expend test1",
			fields: fields{
				length: 0,
				cap:    0,
				newCap: 0,
			},
			wantNewCap: 0,
			wantErr:    false,
		},
		{
			name: "expend test2",
			fields: fields{
				length: 1,
				cap:    1,
				newCap: 1,
			},
			wantNewCap: 1,
			wantErr:    false,
		},
		{
			name: "expend test3",
			fields: fields{
				length: 2,
				cap:    2,
				newCap: 2,
			},
			wantNewCap: 2,
			wantErr:    false,
		},
		{
			name: "expend test4",
			fields: fields{
				length: 1,
				cap:    2,
				newCap: 3,
			},
			wantNewCap: 3,
			wantErr:    false,
		},
		{
			name: "shrink test1",
			fields: fields{
				length: 1,
				cap:    1,
				newCap: -1,
			},
			wantNewCap: 1,
			wantErr:    true,
		},
		{
			name: "shrink test2",
			fields: fields{
				length: 1,
				cap:    2,
				newCap: 0,
			},
			wantNewCap: 0,
			wantErr:    true,
		},
		{
			name: "shrink test3",
			fields: fields{
				length: 1,
				cap:    2,
				newCap: 1,
			},
			wantNewCap: 1,
			wantErr:    false,
		},
		{
			name: "shrink test4",
			fields: fields{
				length: 0,
				cap:    2,
				newCap: 0,
			},
			wantNewCap: 0,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := buildDynamicLinkQueueTestInstance(tt.fields.length, tt.fields.cap, 1)
			err := q.ExpendOrShrink(tt.fields.newCap)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExpendOrShrink() error = %v, wantErr %v", err, tt.wantErr)
			}
			if (err != nil) != tt.wantErr && q.Cap() != tt.wantNewCap {
				t.Errorf("ExpendOrShrink() capacity = %v, want %v", q.Cap(), tt.wantNewCap)
			}
		})
	}
}

func TestDynamicLinkQueue_EnQueueList(t *testing.T) {
	type fields struct {
		length   int
		capacity int
	}

	tests := []struct {
		name    string
		fields  fields
		elem    []int
		wantErr bool
	}{
		{
			name: "enqueue test1",
			fields: fields{
				length:   0,
				capacity: 0,
			},
			elem:    make([]int, 0),
			wantErr: false,
		},
		{
			name: "enqueue test2",
			fields: fields{
				length:   1,
				capacity: 1,
			},
			elem:    make([]int, 1),
			wantErr: true,
		},
		{
			name: "enqueue test3",
			fields: fields{
				length:   0,
				capacity: 5,
			},
			elem:    []int{1, 2, 3, 4, 5},
			wantErr: false,
		},
		{
			name: "enqueue test4",
			fields: fields{
				length:   0,
				capacity: 5,
			},
			elem:    []int{1, 2, 3, 4, 5, 6},
			wantErr: true,
		},
		{
			name: "enqueue test5",
			fields: fields{
				length:   4,
				capacity: 5,
			},
			elem:    []int{1, 2},
			wantErr: true,
		},
		{
			name: "enqueue test6",
			fields: fields{
				length:   4,
				capacity: 5,
			},
			elem:    []int{1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := buildDynamicLinkQueueTestInstance(tt.fields.length, tt.fields.capacity, -1)
			if err := q.EnQueueList(tt.elem...); (err != nil) != tt.wantErr {
				t.Errorf("EnQueueList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDynamicLinkQueue_DeQueueIntoArray(t *testing.T) {
	type fields struct {
		length    int
		capacity  int
		queueElem []int
	}
	buildTestInstance := func(length, capacity int, elem ...int) *DynamicLinkQueue {
		result := NewDynamicLinkQueue(capacity)

		for _, v := range elem {
			result.tail.next = new(queueData)
			result.tail = result.tail.next
			result.tail.elem = v
			result.length++
		}
		return result
	}
	tests := []struct {
		name    string
		fields  fields
		count   int
		want    []int
		wantErr bool
	}{
		{
			name: "dequeue test1",
			fields: fields{
				length:    0,
				capacity:  0,
				queueElem: []int{1},
			},
			count:   1,
			want:    []int{},
			wantErr: true,
		},
		{
			name: "dequeue test2",
			fields: fields{
				length:    1,
				capacity:  1,
				queueElem: []int{1},
			},
			count:   1,
			want:    []int{1},
			wantErr: false,
		},
		{
			name: "dequeue test3",
			fields: fields{
				length:    1,
				capacity:  1,
				queueElem: []int{1},
			},
			want:    []int{},
			count:   2,
			wantErr: true,
		},
		{
			name: "dequeue test4",
			fields: fields{
				length:    5,
				capacity:  5,
				queueElem: []int{1, 2, 3, 4, 5},
			},
			want:    make([]int, 0),
			count:   -1,
			wantErr: true,
		},
		{
			name: "dequeue test5",
			fields: fields{
				length:    5,
				capacity:  5,
				queueElem: []int{1, 2, 3, 4, 5},
			},
			want:    make([]int, 0),
			count:   0,
			wantErr: false,
		},
		{
			name: "dequeue test6",
			fields: fields{
				length:    5,
				capacity:  5,
				queueElem: []int{1, 2, 3, 4, 5},
			},
			count:   2,
			want:    []int{1, 2},
			wantErr: false,
		},
		{
			name: "dequeue test7",
			fields: fields{
				length:    5,
				capacity:  5,
				queueElem: []int{1, 2, 3, 4, 5},
			},
			count:   5,
			want:    []int{1, 2, 3, 4, 5},
			wantErr: false,
		},
		{
			name: "dequeue test8",
			fields: fields{
				length:    5,
				capacity:  5,
				queueElem: []int{1, 2, 3, 4, 5},
			},
			count:   6,
			want:    []int{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := buildTestInstance(tt.fields.length, tt.fields.capacity, tt.fields.queueElem...)
			got, err := q.DeQueueIntoArray(tt.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeQueueIntoArray() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeQueueIntoArray() got = %v, want %v", got, tt.want)
			}
		})
	}
}
