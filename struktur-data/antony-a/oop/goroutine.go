package oop

import (
	"fmt"
	"sync"
	"time"
)

type Order struct {
	id       int
	car      string
	capacity int // kapasitas tangki mobil
}

// stasiun pengisian BBM
type GasStation struct {
	queue chan Order
	wg    sync.WaitGroup
}

func NewGasStation() *GasStation {
	return &GasStation{
		queue: make(chan Order, 10), // stasiun memiliki
	}
}

// available pompa
func (g *GasStation) Tank(tankId int) {
	defer g.wg.Done()

	// dispatch any available queue
	for q := range g.queue {
		fmt.Printf(
			"[Concurrency] tank_id #%d start fill order_id #%d for car #%s",
			tankId,
			q.id,
			q.car,
		)

		// spend time for fill 1 car based on capacity
		capacityTime := time.Duration(q.capacity)
		time.Sleep(capacityTime * time.Second)

		fmt.Printf(
			"[Concurrency] tank_id #%d end fill gas with order_id #%d for car #%s",
			tankId,
			q.id,
			q.car,
		)
	}
}
