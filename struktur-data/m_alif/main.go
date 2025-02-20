package main

import (
	"fmt"
	"sync"
)

// Struct: Digunakan untuk menyimpan data kompleks
type Person struct {
	Name string
	Age  int
}

// Interface: Mendefinisikan kontrak method tanpa implementasi spesifik
type Animal interface {
	Speak() string
}

type Dog struct{}
func (d Dog) Speak() string { return "Woof!" }

type Cat struct{}
func (c Cat) Speak() string { return "Meow!" }

// Linked List: Implementasi manual menggunakan struct dan pointer
type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) Append(value int) {
	newNode := &Node{Value: value}
	if l.Head == nil {
		l.Head = newNode
		return
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (l *LinkedList) Print() {
	current := l.Head
	for current != nil {
		fmt.Print(current.Value, " -> ")
		current = current.Next
	}
	fmt.Println("nil")
}

// Queue: Implementasi menggunakan slice
type Queue struct {
	items []int
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		return -1
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// Stack: Implementasi menggunakan slice
type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		return -1
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// Channel: Digunakan untuk komunikasi antar goroutine
func worker(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- "Goroutine finished!"
}

func main() {
	// Array: Ukuran tetap
	var arr = [3]int{1, 2, 3}
	fmt.Println("Array:", arr)

	// Slice: Ukuran dinamis
	slice := []int{1, 2, 3}
	slice = append(slice, 4, 5)
	fmt.Println("Slice:", slice)

	// Map: Key-value pair
	m := make(map[string]int)
	m["Alice"] = 25
	m["Bob"] = 30
	fmt.Println("Map:", m)

	// Struct
	person := Person{Name: "John", Age: 28}
	fmt.Println("Struct:", person)

	// Pointer
	age := 30
	ptr := &age
	fmt.Println("Pointer Value:", *ptr)

	// Interface
	var animals []Animal = []Animal{Dog{}, Cat{}}
	for _, animal := range animals {
		fmt.Println("Animal Sound:", animal.Speak())
	}

	// Linked List
	ll := LinkedList{}
	ll.Append(10)
	ll.Append(20)
	ll.Append(30)
	fmt.Print("Linked List: ")
	ll.Print()

	// Queue
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	fmt.Println("Queue Dequeue:", q.Dequeue())

	// Stack
	s := Stack{}
	s.Push(1)
	s.Push(2)
	fmt.Println("Stack Pop:", s.Pop())

	// Channel (Concurrency)
	ch := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)
	go worker(ch, &wg)
	fmt.Println(<-ch)
	wg.Wait()
}
