package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fairglen/tradin-bot/handlers"
	"github.com/fairglen/tradin-bot/internal/exchanges/exchangesfakes"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

type route struct {
	path    string
	methods []string
}

func getRoutes(t *testing.T, r *mux.Router) []route {
	routes := []route{}
	err := r.Walk(func(r *mux.Route, rtr *mux.Router, ancestors []*mux.Route) error {
		p, _ := r.GetPathTemplate()
		m, _ := r.GetMethods()

		routes = append(routes, route{path: p, methods: m})
		return nil
	})

	if err != nil {
		t.Fatal(err)
	}
	return routes
}

func TestNewRouter(t *testing.T) {
	scenarios := []struct {
		name  string
		route route
	}{
		{
			name:  "has buy route",
			route: route{path: "/coinbase/buy", methods: []string{http.MethodPost}},
		},
	}
	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			ch := handlers.NewCoinbase(&exchangesfakes.FakeClient{})
			require.Contains(t, getRoutes(t, ch.Router), scenario.route)
		})
	}
}

func TestBuy(t *testing.T) {
	scenarios := []struct {
		name        string
		request     func() *http.Request
		expStatus   int
		expResponse string
	}{
		{name: "responds correctly",
			request: func() *http.Request {
				r, err := http.NewRequest("POST", "/coinbase/buy", nil)
				if err != nil {
					t.Fatal(err)
				}
				return r
			},
			expStatus:   http.StatusOK,
			expResponse: "Congrats you bough bitcoin!",
		},
	}
	for _, scenario := range scenarios {
		t.Run(scenario.name, func(*testing.T) {
			rr := httptest.NewRecorder()

			fc := &exchangesfakes.FakeClient{}
			ch := handlers.NewCoinbase(fc)

			h := http.HandlerFunc(ch.Buy)
			h.ServeHTTP(rr, scenario.request())

			require.Equal(t, scenario.expStatus, rr.Code)
			require.Equal(t, scenario.expStatus, rr.Code)
			require.Equal(t, scenario.expResponse, rr.Body.String())
			require.Equal(t, 1, fc.CreateOrderCallCount())
		})
	}
}
