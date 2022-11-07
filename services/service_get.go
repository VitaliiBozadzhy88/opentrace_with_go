package main

import (
	"fmt"
	"net/http"

	"open_trace_with_go/handlers"
	"open_trace_with_go/initialize_tracer"
)

func main() {
	closer, err := initialize_tracer.InitializeTracer("service_get")
	if err != nil {
		panic(err.Error())
	}
	defer closer.Close()

	http.HandleFunc("/get", handlers.HandleGet)

	fmt.Println("SERVICE START IN PORT 8081")
	http.ListenAndServe(":8081", nil)
}
