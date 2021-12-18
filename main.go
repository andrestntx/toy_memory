package main

import (
	"flink/config"
	"flink/internal/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.Path("/location/{order_id}").Methods("PUT").HandlerFunc(handlers.Put)
	router.Path("/location/{order_id}").Methods("GET").Queries("max", "{[0-9]*?}").HandlerFunc(handlers.Get)
	router.Path("/location/{order_id}").Methods("GET").HandlerFunc(handlers.Get)
	router.Path("/location/{order_id}").Methods("DELETE").HandlerFunc(handlers.Delete)

	fmt.Println("Starting server at port 8080")

	if err := http.ListenAndServe(":"+config.Port, router); err != nil {
		log.Fatal(err)
	}
}
