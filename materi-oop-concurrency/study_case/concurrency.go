package study_case

import (
	"fmt"
	"sync"
	"time"
	"runtime"
)

type Order struct {
	id int
	menuItem string
	cookTime time.Duration
}

type Restaurant struct {
	orderQueue chan Order
	wg sync.WaitGroup
}

type SequentialRestauraurant struct {
	orders []Order
}

func NewRestaurant() *Restaurant {
	return &Restaurant{
		orderQueue: make(chan Order, 10),
	}
}

func NewSequentialRestaurant() *SequentialRestauraurant {
	return &SequentialRestauraurant{
		orders: make([]Order, 0),
	}
}

// case concurrency
func (r *Restaurant) Cook(cookID int) {
	defer r.wg.Done()
	for order := range r.orderQueue {
		fmt.Printf("[Concurrency] Koki %d mulai memasak %s untuk pesanan #%d\n",
		cookID, order.menuItem, order.id)
		time.Sleep(order.cookTime * time.Second) // sebuah jeda
		fmt.Printf("[Concurrency] Koki %d selesai memasak %s untuk pesanan #%d\n",
		cookID, order.menuItem, order.id)
	}
}

// case sequential
func (r *SequentialRestauraurant) Cook(cookID int, order Order) {
	fmt.Printf("[Sequential] Koki %d mulai memasak %s untuk pesanan #%d\n",
		cookID, order.menuItem, order.id)
	time.Sleep(order.cookTime * time.Second) // sebuah jeda
	fmt.Printf("[Sequential] Koki %d selesai memasak %s untuk pesanan #%d\n",
		cookID, order.menuItem, order.id)
}

// case concurrency
func RunConcurrent(orders []Order) time.Duration {
	start := time.Now()

	restaurant := NewRestaurant()

	numCooks := 3
	restaurant.wg.Add(numCooks)

	for i := 1; i <= numCooks; i++ {
		go restaurant.Cook(i)
	}

	for _, order := range orders {
		restaurant.orderQueue <- order
	}

	close(restaurant.orderQueue)
	restaurant.wg.Wait()

	return time.Since(start)
}

// case sequential
func RunSequential(orders []Order) time.Duration {
	start := time.Now()

	restaurant := NewSequentialRestaurant()

	for _, order := range orders {
		restaurant.orders = append(restaurant.orders, order)
		restaurant.Cook(1, order)
	}

	return time.Since(start)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
} 
func RunTime() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("Alloc = %v\n", bToMb(m.Alloc))
	fmt.Printf("TotalAlloc = %v\n", bToMb(m.TotalAlloc))
	fmt.Printf("Sys = %v\n", bToMb(m.TotalAlloc))

}

func ConcurrencyCase() {
	orders := []Order {
		{1, "Nasi Goreng", 2},
		{2, "Mie Goreng", 3},
		{3, "Ayam Goreng", 1},
		{4, "Mie Rebus", 1},
		{5, "Spaggeti", 3},
		{6, "Mie Ayam", 3},
	}

	// runnning runtime
	fmt.Println("Sequential Process")
	seqTime := RunSequential(orders)

	RunTime()

	fmt.Println("Concurrency Process")
	concTime := RunConcurrent(orders)

	fmt.Println("Perbandingan")
	fmt.Printf("Waktu sequential: %v\n", seqTime)
	fmt.Printf("Waktu concurrency: %v\n", concTime)
	fmt.Printf("Selisih waktu: %v\n", seqTime-concTime)

}