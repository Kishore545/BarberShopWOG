//Without Using Goroutine Channel:
//------------------------------------

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Customer struct {
	ID int
}

type BarberShop struct {
	queue      []Customer
	maxSeats   int
	barberBusy bool
}

func NewBarberShop(seats int) *BarberShop {
	return &BarberShop{
		queue:    make([]Customer, 0),
		maxSeats: seats,
	}
}

func (bs *BarberShop) Arrive(customer Customer) {
	fmt.Printf("Customer %d arrived.\n", customer.ID)
	if !bs.barberBusy {
		bs.barberBusy = true
		bs.cutHair(customer)
	} else if len(bs.queue) < bs.maxSeats {
		bs.queue = append(bs.queue, customer)
		fmt.Printf("Customer %d is waiting. Queue length: %d\n", customer.ID, len(bs.queue))
	} else {
		fmt.Printf("Customer %d left as no seats were available.\n", customer.ID)
	}
}

func (bs *BarberShop) cutHair(customer Customer) {
	fmt.Printf("Barber is cutting hair for Customer %d.\n", customer.ID)
	time.Sleep(30 * time.Second) // Simulating haircut time
	fmt.Printf("Customer %d's haircut is done.\n", customer.ID)
	bs.barberBusy = false

	if len(bs.queue) > 0 {
		nextCustomer := bs.queue[0]
		bs.queue = bs.queue[1:]
		bs.barberBusy = true
		bs.cutHair(nextCustomer)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	shop := NewBarberShop(3) // Maximum 3 waiting seats

	for i := 1; i <= 10; i++ {
		customer := Customer{ID: i}
		shop.Arrive(customer)
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second) // Random arrival time
	}
}
