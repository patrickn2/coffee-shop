package handler

import (
	"log"
	"time"

	"github.com/patrickn2/coffee-shop/order"
)

type Handler struct {
	name string
	in   <-chan *order.Order
	done chan bool
}

func New(name string, in <-chan *order.Order) *Handler {
	log.Printf("Handler %s arrived at the Coffee Shop", name)
	return &Handler{
		name: name,
		in:   in,
		done: make(chan bool),
	}
}

func (c *Handler) StartShift() {
	log.Printf("Starting Handler %s Shift\n", c.name)
	go c.waitOrders()
}

func (c *Handler) EndShift() {
	log.Printf("Ending Handler %s Shift\n", c.name)
	c.done <- true
}

func (c *Handler) waitOrders() {
	for {
		select {
		case order := <-c.in:
			log.Printf("\033[0;34mHandling Order for %s\033[0m", order.ConsumerName)
			time.Sleep(time.Second * 2)
			log.Printf("\033[0;34mOrder handled to %s\033[0m", order.ConsumerName)
		case <-c.done:
			log.Printf("Handler %s Shift Ended\n", c.name)
			return
		}
	}
}
