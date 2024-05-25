package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var movies []Movie

const PORT string = "8080"

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(movies)
	data, _ := json.Marshal(movies)
	io.WriteString(w, string(data))
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, item := range movies {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	// set json content type
	w.Header().Set("Content-Type", "application/json")
	// params
	params := mux.Vars(r)
	// loop over the movies, range
	for index, item := range movies {
		// delete the movie with the id that you have sent
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			// add a new movie  - the movie that we send in the request
			var movie Movie
			json.NewDecoder(r.Body).Decode(&movie)
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movies)
		}

	}

}

func main() {

	r := mux.NewRouter()

	movies = append(movies, Movie{
		Id:    "1",
		Isbn:  "12345",
		Title: "Supernatural",
		Director: &Director{
			FirstName: "Eric",
			LastName:  "Kripke",
		},
	})

	movies = append(movies, Movie{
		Id:    "2",
		Isbn:  "45678",
		Title: "Harry Potter and the Philosophers stone",
		Director: &Director{
			FirstName: "JK",
			LastName:  "Rowling",
		},
	})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/mvoies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server on PORT: %s\n", PORT)

	if err := http.ListenAndServe(":"+PORT, r); err != nil {
		log.Fatal("Error during server startup", err)
	}

}
