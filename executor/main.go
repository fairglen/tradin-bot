package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fairglen/tradin-bot/handlers"
	"github.com/fairglen/tradin-bot/internal/exchanges"
)

func main() {
	fmt.Println("TradinBot:Executor")

	coinbase := exchanges.Coinbase{}

	ch := handlers.NewCoinbase(coinbase)
	log.Fatal(http.ListenAndServe(":8000", ch.Router))
}
