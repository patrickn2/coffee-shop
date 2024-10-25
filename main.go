package main

import (
	"log"
	"time"

	"github.com/patrickn2/coffee-shop/barista"
	"github.com/patrickn2/coffee-shop/cashier"
	"github.com/patrickn2/coffee-shop/handler"
	"github.com/patrickn2/coffee-shop/order"
)

func main() {
	cashierLine := make(chan *order.Order, 1000)
	baristaLine := make(chan *order.Order, 1000)
	handlerLine := make(chan *order.Order, 1000)

	cashier1 := cashier.New("Alex", cashierLine, baristaLine, 1000)
	barista1 := barista.New("Jhonna", baristaLine, handlerLine, "\033[0;32m")
	barista2 := barista.New("Lucas", baristaLine, handlerLine, "\033[0;35m")
	handler1 := handler.New("Robert", handlerLine)

	cashier1.StartShift()
	barista1.StartShift()
	handler1.StartShift()
	barista2.StartShift()
	defer cashier1.EndShift()
	defer barista1.EndShift()
	defer handler1.EndShift()

	log.Println("Opening Coffee Shop")

	order.New("Yuri", cashierLine)
	order.New("Miguel", cashierLine)
	order.New("Alves", cashierLine)
	order.New("Robert", cashierLine)
	order.New("Adamastor", cashierLine)

	time.Sleep(time.Second * 10)
	barista2.EndShift()
	time.Sleep(time.Second * 50)

	log.Println("Closing Coffee Shop")

}
