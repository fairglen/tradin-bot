package coinbase

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/coinbase/buy", Buy).Methods(http.MethodPost)

	return r
}

func Buy(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Congrats you bough bitcoin!")
}
