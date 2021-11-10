package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/nooderg/pipiSpot/internal/configs"
	"github.com/nooderg/pipiSpot/internal/controllers"
)

func main() {
	log.Println("Init routes...")
	initRouter()
}

func initRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/api/users", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/api/users/{user_id:[0-9]+}", controllers.GetUser).Methods("GET")
	r.HandleFunc("/api/users/{user_id:[0-9]+}", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/api/users/{user_id:[0-9]+}", controllers.DeleteUser).Methods("DELETE")

	r.HandleFunc("/api/spots", controllers.CreateSpot).Methods("POST")
	r.HandleFunc("/api/spots/{spot_id:[0-9]+}", controllers.GetSpot).Methods("GET")
	r.HandleFunc("/api/spots/{spot_id:[0-9]+}", controllers.UpdateSpot).Methods("PUT")
	r.HandleFunc("/api/spots/{spot_id:[0-9]+}", controllers.DeleteSpot).Methods("DELETE")

	r.HandleFunc("/api/spots/{spot_id:[0-9]+}/ratings", controllers.CreateRatings).Methods("POST")
	r.HandleFunc("/api/spots/{spot_id:[0-9]+}/ratings/{ratings_id:[0-9]+}", controllers.GetRatings).Methods("GET")
	r.HandleFunc("/api/spots/{spot_id:[0-9]+}/ratings/{ratings_id:[0-9]+}", controllers.UpdateRatings).Methods("PUT")
	r.HandleFunc("/api/spots/{spot_id:[0-9]+}/ratings/{ratings_id:[0-9]+}", controllers.DeleteRatings).Methods("DELETE")

	r.HandleFunc("/api/spots/{spot_id:[0-9]+}/comments", controllers.CreateComment).Methods("POST")
	r.HandleFunc("/api/spots/{spot_id:[0-9]+}/comments/{comm_id:[0-9]+}", controllers.GetComment).Methods("GET")
	r.HandleFunc("/api/spots/{spot_id:[0-9]+}/comments/{comm_id:[0-9]+}", controllers.UpdateComment).Methods("PUT")
	r.HandleFunc("/api/spots/{spot_id:[0-9]+}/comments/{comm_id:[0-9]+}", controllers.DeleteComment).Methods("DELETE")

	log.Println("Server ready!")
	log.Fatal(http.ListenAndServe(":8080", r))
}
