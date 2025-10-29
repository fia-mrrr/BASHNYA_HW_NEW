package main

import "fmt"

type Node struct {
	data     int
	next     *Node
	previous *Node
}

type Deque struct {
	head *Node
	tail *Node
	size int
}

func NewDeque() *Deque {
	return &Deque{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (d *Deque) PushFront(data int) {
	newNode := &Node{data: data}

	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		newNode.next = d.head
		d.head.previous = newNode
		d.head = newNode
	}
	d.size++
}

func (d *Deque) PushBack(data int) {
	newNode := &Node{data: data}

	if d.tail == nil {
		d.tail = newNode
		d.head = newNode
	} else {
		newNode.previous = d.tail
		d.tail.next = newNode
		d.tail = newNode
	}
	d.size++
}

func (d *Deque) PopFront() (int, bool) {
	if d.head == nil {
		return 0, false
	}

	data := d.head.data

	if d.head == d.tail {
		d.head = nil
		d.tail = nil
	} else {
		d.head = d.head.next
		d.head.previous = nil
	}
	d.size--

	return data, true
}

func (d *Deque) PopBack() (int, bool) {
	if d.tail == nil {
		return 0, false
	}

	data := d.tail.data

	if d.tail == d.head {
		d.tail = nil
		d.head = nil
	} else {
		d.tail = d.tail.previous
		d.tail.next = nil
	}
	d.size--

	return data, true
}

func (d *Deque) IsEmpty() bool {
	return d.size == 0
}

func (d *Deque) Size() int {
	return d.size
}

func (d *Deque) Clear() {
	d.head = nil
	d.tail = nil
	d.size = 0
}

func (d *Deque) Print() {
	if d.head == nil {
		fmt.Println("Очередь пуста")
		return
	}

	current := d.head
	fmt.Print("Очередь: ")
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

func main() {
	deque := NewDeque()

	deque.PushBack(31)
	deque.PushBack(2)
	deque.PushFront(-4)
	deque.PushFront(9)
	deque.PushBack(82)

	deque.Print()
	fmt.Printf("Размер очереди: %d\n", deque.Size())
	fmt.Printf("Проверка на пустую очередь: %t\n", deque.IsEmpty())

	if data, ok := deque.PopBack(); ok {
		fmt.Printf("Удаление и возвращение элемента из конца очереди: %d\n", data)
	}
	if data, ok := deque.PopFront(); ok {
		fmt.Printf("Удаление и возвращение элемента из начала очереди: %d\n", data)
	}

	deque.Print()
	fmt.Printf("Размер очереди после удаления элементов: %d\n", deque.size)

	deque.Clear()
	fmt.Printf("Размер очереди после очищения: %d\n", deque.size)
	fmt.Printf("Проверка на пустую очередь: %t\n", deque.IsEmpty())
}
