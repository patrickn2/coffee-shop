package order

import "time"

type Item struct {
	Name        string
	PrepareTime time.Duration
}

type Order struct {
	Id           string
	ConsumerName string
	Items        []Item
	Status       string
}

func New(consumerName string, cashierLine chan<- *Order) {
	Expresso := Item{
		Name:        "Expresso",
		PrepareTime: time.Second * 1,
	}
	Bagel := Item{
		Name:        "Bagel",
		PrepareTime: time.Second * 2,
	}
	OrangeJuice := Item{
		Name:        "Orange Juice",
		PrepareTime: time.Second * 3,
	}
	cashierLine <- &Order{
		ConsumerName: consumerName,
		Items: []Item{
			Expresso,
			Bagel,
			OrangeJuice,
		},
	}
}

func (o *Order) SetId(id string) {
	o.Id = id
}
