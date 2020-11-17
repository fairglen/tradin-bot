package handlers

import (
	"fmt"
	"net/http"

	"github.com/fairglen/tradin-bot/internal/exchanges"
	"github.com/gorilla/mux"
)

type Coinbase struct {
	client exchanges.Client
	Router *mux.Router
}

func NewCoinbase(client exchanges.Client) *Coinbase {
	r := mux.NewRouter()
	c := &Coinbase{
		client: client,
		Router: r,
	}
	c.Router.HandleFunc("/coinbase/buy", c.Buy).Methods(http.MethodPost)
	return c
}

func (c *Coinbase) Buy(w http.ResponseWriter, r *http.Request) {
	c.client.CreateOrder(exchanges.Order{})
	fmt.Fprint(w, "Congrats you bough bitcoin!")
}
