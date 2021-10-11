package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/nooderg/pipiSpot/internal/configs"
	"github.com/nooderg/pipiSpot/internal/controllers"
	"github.com/nooderg/pipiSpot/pkg/responses"
)

func main() {
	log.Println("Init routes...")
	initRouter()
}

func initRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/api/user/:id", helloworld).Methods("GET")
	r.HandleFunc("/api/users", helloworld).Methods("GET")
	r.HandleFunc("/api/user", controllers.PostUser).Methods("POST")

	r.HandleFunc("/api/points/:id", helloworld).Methods("GET")
	r.HandleFunc("/api/points", helloworld).Methods("GET")
	r.HandleFunc("/api/point", helloworld).Methods("POST")
	r.HandleFunc("/api/point/:id", helloworld).Methods("PUT")

	r.HandleFunc("/api/point/:id/ratings/standing", helloworld).Methods("PUT")
	r.HandleFunc("/api/point/:id/ratings/sitting", helloworld).Methods("PUT")
	r.HandleFunc("/api/point/:id/comment/:idcomm", helloworld).Methods("POST")

	log.Println("Server ready!")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	responses.WriteData(w, "hello world, I'm gonna change the way we... You know.")
}
