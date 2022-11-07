package main

import (
	"fmt"
	"net/http"

	"open_trace_with_go/handlers"
	"open_trace_with_go/initialize_tracer"
)

func main() {
	closer, err := initialize_tracer.InitializeTracer("service_create")
	if err != nil {
		panic(err.Error())
	}
	defer closer.Close()

	http.HandleFunc("/create", handlers.HandleCreate)

	fmt.Println("SERVICE START IN PORT 8080")
	http.ListenAndServe(":8080", nil)
}
