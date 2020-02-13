package main

import (
	"fmt"

	"octopus-project.t-mux/router"
)

func main() {

	handlerA := 17
	handlerB := 18
	handlerC := 19

	router_1 := router.NewRouter()

	router_1.Get("/a", handlerA)
	router_1.Get("/test/of/new/path", handlerB)
	router_1.Get("/test/of/new/:var", handlerC)

	router_2 := router.NewRouter()

	router_2.Get("/ping/1", 1)
	router_2.Get("/ping/2", 2)
	router_2.Get("/ping/3", 3)

	// router_1.Mount("/:mounted", router_2)

	handler1 := router_1.GetHandler("/test/of/new/hello", "GET")
	fmt.Println("handler1: ", handler1)

	handler2 := router_1.GetHandler("/path/to/hell", "GET")
	fmt.Println("handler1: ", handler2)
}
