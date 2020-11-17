package exchanges

type Order struct {
	Price     string
	Size      string
	Side      string
	ProductID string
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . Client
type Client interface {
	CreateOrder(order Order) error
}

func NewCoinbaseClient() (client Client, err error) {
	return nil, nil
}
