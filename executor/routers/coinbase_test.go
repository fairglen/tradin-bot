package coinbase

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

type route struct {
	path    string
	methods []string
}

func getRoutes(r *mux.Router) []route {
	routes := []route{}
	r.Walk(func(r *mux.Route, rtr *mux.Router, ancestors []*mux.Route) error {
		p, _ := r.GetPathTemplate()
		m, _ := r.GetMethods()
		routes = append(routes, route{path: p, methods: m})
		return nil
	})
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
		{
			name:  "has sell route",
			route: route{path: "/coinbase/sell", methods: []string{http.MethodPost}},
		},
	}
	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			require.Contains(t, getRoutes(NewRouter()), scenario.route)
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
			h := http.HandlerFunc(Buy)
			h.ServeHTTP(rr, scenario.request())

			require.Equal(t, scenario.expStatus, rr.Code)
			require.Equal(t, scenario.expStatus, rr.Code)
			require.Equal(t, scenario.expResponse, rr.Body.String())
		})
	}
}
