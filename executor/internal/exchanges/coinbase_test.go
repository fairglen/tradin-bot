package exchanges_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fairglen/tradin-bot/internal/exchanges"
	"github.com/stretchr/testify/require"
)

func TestCoinbaseCreateOrder(t *testing.T) {
	scenarios := []struct {
		name         string
		order        exchanges.Order
		coinbaseMock http.HandlerFunc
	}{
		{
			name: "creates buy order",
			order: exchanges.Order{
				Price:     "",
				Size:      "",
				Side:      "",
				ProductID: "",
			},
			coinbaseMock: func(w http.ResponseWriter, r *http.Request) {
				require.Equal(t, "/orders", r.URL.Path)
				require.Equal(t, http.MethodPost, r.Method)
				require.Equal(t, "application/json", r.Header.Get("content-type"))
				require.Equal(t, "token", r.Header.Get("Authorization"))
				w.WriteHeader(http.StatusAccepted)
			},
		},
	}
	for _, scenario := range scenarios {
		t.Run(scenario.name, func(*testing.T) {
			coinbase := exchanges.Coinbase{}
			server := httptest.NewServer(http.HandlerFunc(scenario.coinbaseMock))
			defer server.Close()
			coinbase.CreateOrder(scenario.order)
		})
	}
}
