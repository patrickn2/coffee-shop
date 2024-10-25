package barista

import (
	"log"
	"time"

	"github.com/patrickn2/coffee-shop/order"
)

type Barista struct {
	name  string
	in    <-chan *order.Order
	out   chan<- *order.Order
	done  chan bool
	color string
}

func New(name string, in <-chan *order.Order, out chan<- *order.Order, color string) *Barista {
	log.Printf("Barista %s arrived at the Coffee Shop", name)
	return &Barista{
		name:  name,
		in:    in,
		out:   out,
		done:  make(chan bool),
		color: color,
	}
}

func (c *Barista) StartShift() {
	log.Printf("Starting Barista %s Shift\n", c.name)
	go c.waitOrders()
}

func (c *Barista) EndShift() {
	log.Printf("Ending Barista %s Shift\n", c.name)
	c.done <- true
}

func (c *Barista) waitOrders() {
	for {
		select {
		case order := <-c.in:
			log.Printf("%sPreparring Order for %s\033[0m", c.color, order.ConsumerName)
			for _, item := range order.Items {
				log.Printf("%s-- Preparing Item %s\033[0m", c.color, item.Name)
				time.Sleep(item.PrepareTime)
			}
			log.Printf("%sOrder for %s ready\033[0m", c.color, order.ConsumerName)
			c.out <- order
		case <-c.done:
			log.Printf("Barista %s Shift Ended\n", c.name)
			return
		}
	}
}
