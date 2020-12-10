package exchanges

import "fmt"

type Coinbase struct {
}

func (c Coinbase) CreateOrder(order Order) error {
	fmt.Println("CreateOrder")
	return nil
}
