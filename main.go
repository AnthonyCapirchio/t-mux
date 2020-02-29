package main

import (
	"fmt"
	"net/http"

	"github.com/AnthonyCapirchio/t-mux/router"
)

func a(w http.ResponseWriter, r *http.Request, params map[string]string) {
	fmt.Println("params: ", params)
}

type DryResponseWriter struct{}

func (d DryResponseWriter) Header() http.Header {
	return http.Header{}
}

func (d DryResponseWriter) Write([]byte) (int, error) {
	return 1, nil
}

func (d DryResponseWriter) WriteHeader(statusCode int) {}

func main() {

	router_1 := router.NewRouter()

	router_1.Get("/a", a)
	router_1.Get("/test/of/new/path", a)
	router_1.Get("/test/of/new/:var", a)

	// router_2 := router.NewRouter()

	// router_2.Get("/ping/1", 1)
	// router_2.Get("/ping/2", 2)
	// router_2.Get("/ping/3", 3)

	// router_1.Mount("/:mounted", router_2)

	w := DryResponseWriter{}
	r := http.Request{}

	handler, params := router_1.GetHandler("/test/of/new/hello", "GET")
	if handler != nil {
		handler(w, &r, params)
	}

	// ----------------

	handler, params = router_1.GetHandler("/test/of/new/hello", "POST")
	if handler != nil {
		handler(w, &r, params)
	}

	// ----------------

	handler, params = router_1.GetHandler("/path/to/heaven", "GET")
	if handler != nil {
		handler(w, &r, params)
	}

	// handler2 := router_1.GetHandler("/path/to/hell", "GET")
	// fmt.Println("handler1: ", handler2)
}
