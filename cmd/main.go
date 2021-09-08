package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nooderg/pipiSpot/pkg/responses"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/hello", helloworld).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	responses.WriteData(w, "hello world, I'm gonna change the way we... You know.")
}
