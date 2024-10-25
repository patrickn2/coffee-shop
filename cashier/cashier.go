package cashier

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/patrickn2/coffee-shop/order"
)

type Cashier struct {
	name  string
	in    <-chan *order.Order
	out   chan<- *order.Order
	done  chan bool
	speed int
}

func New(name string, in <-chan *order.Order, out chan<- *order.Order, speed int) *Cashier {
	log.Printf("Cashier %s arrived at the Coffee Shop", name)
	return &Cashier{
		name:  name,
		in:    in,
		out:   out,
		done:  make(chan bool),
		speed: speed,
	}
}

func (c *Cashier) StartShift() {
	log.Printf("Starting Cashier %s Shift\n", c.name)
	go c.waitOrders()
}

func (c *Cashier) EndShift() {
	log.Printf("Ending Cashier %s Shift\n", c.name)
	c.done <- true
}

func (c *Cashier) waitOrders() {
	for {
		select {
		case order := <-c.in:
			log.Printf("\033[0;31mRegistering Order for %s\033[0m", order.ConsumerName)
			for _, item := range order.Items {
				log.Printf("\033[0;31mRegistering Item %s\033[0m", item.Name)
				time.Sleep(time.Millisecond * time.Duration(c.speed))
			}
			id := uuid.NewString()
			log.Printf("\033[0;31mOrder for %s registered id %s\033[0m", order.ConsumerName, uuid.NewString())
			order.SetId(id)
			c.out <- order
		case <-c.done:
			log.Printf("Cashier %s Shift Ended\n", c.name)
			return
		}
	}
}
