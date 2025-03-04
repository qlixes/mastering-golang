package modules

import (
	"fmt"
	"sync"
	"time"
)

type Order struct {
	car      string
	capacity int // kapasitas tangki mobil
}

// stasiun pengisian BBM
type GasStation struct {
	buffered   chan Order
	unbuffered chan Order
	wg         sync.WaitGroup
}

func NewGasStation() *GasStation {
	return &GasStation{
		// buffered channel
		buffered:   make(chan Order, 10), // setup max capacity on each station
		unbuffered: make(chan Order),
	}
}

// available pompa
func (g *GasStation) Tank(tankId int) {
	defer g.wg.Done()

	// dispatch any available queue
	for q := range g.unbuffered {
		fmt.Printf(
			"[Concurrency] tank_id #%d start fill car #%s \n",
			tankId,
			q.car,
		)

		// spend time for fill 1 car based on capacity
		capacityTime := time.Duration(q.capacity)
		time.Sleep(capacityTime * time.Second)

		fmt.Printf(
			"[Concurrency] tank_id #%d end fill gas with car #%s \n",
			tankId,
			q.car,
		)
	}
}

func RunConcurrent(orders []Order) time.Duration {
	start := time.Now()

	// init station
	station := NewGasStation()

	// max available Tank
	capacity := 3

	// register Tank into station
	station.wg.Add(capacity)

	for i := 1; i <= capacity; i++ {
		go station.Tank(i)
	}

	// load all orders
	for _, order := range orders {
		// load all order into queue
		station.unbuffered <- order
	}

	// close queue after list order was loaded
	close(station.unbuffered)

	// concurrent will wait process
	station.wg.Wait()

	return time.Since(start)
}

func RunCase() {
	orders := []Order{
		{car: "Toyota", capacity: 30},
		{car: "Honda", capacity: 20},
		{car: "BMW", capacity: 25},
		{car: "Toyota", capacity: 10},
		{car: "Toyota", capacity: 20},
		{car: "Honda", capacity: 15},
		{car: "Toyota", capacity: 5},
		{car: "Toyota", capacity: 10},
		{car: "Ducati", capacity: 5},
		{car: "Mercedez", capacity: 10},
		{car: "Toyota", capacity: 10},
		{car: "Honda", capacity: 12},
		{car: "Honda", capacity: 15},
	}

	concurrentTime := RunConcurrent(orders)

	fmt.Printf("Concurrent Spend Time = %v \n", concurrentTime)
}
