package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// fmt.Printf("hello world")

type Music struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Generies string    `json:"generies"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var Musics []Music

func getMusics(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Musics)

}

func deleteMusic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range Musics {

		if item.ID == params["id"] {

			Musics = append(Musics[:index], Musics[index+1:]...)

			break
		}

	}

}

func getMusic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

}

func createMusic() {

}

func updateMusic() {

}

func main() {

	Musics = append(Musics, Music{ID: "1", Title: "Take me Away", Generies: "Hip-Hop", Director: &Director{Firstname: "Akintola", Lastname: "Olalekan"}})
	Musics = append(Musics, Music{ID: "2", Title: "I'm Home", Generies: "Afrobeats", Director: &Director{Firstname: "Jonny", Lastname: "Lawreance"}})

	r := mux.NewRouter()

	r.HandleFunc("/musics", getMusics).Methods("GET")
	r.HandleFunc("/musics/{id}", getMusic).Methods("GET")
	r.HandleFunc("/musics", createMusic).Methods("POST")
	r.HandleFunc("/musics/{id}", updateMusic).Methods("PUT")
	r.HandleFunc("/musics/{id}", deleteMusic).Methods("DELETE")

	fmt.Printf("Starting Server at Port :8080")

	log.Fatal(http.ListenAndServe(":8000", r))
}
