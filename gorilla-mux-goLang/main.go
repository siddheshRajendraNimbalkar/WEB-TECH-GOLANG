package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"math/rand"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json: "id"`
	Isbn     string    `json: "isbn"`
	Title    string    `json: "title"`
	Director *Director `json: "director"`
}

type Director struct {
	Firstname string `json: "firstname"`
	Lastname  string `json: "lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovieByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, ele := range movies {
		if ele.ID == params["id"] {
			json.NewEncoder(w).Encode(movies[index])
			return
		}
	}
}

func createMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rand.Seed(time.Now().UnixNano())
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movieString := rand.Intn(10000)
	movie.ID = strconv.Itoa(movieString)
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)

}

func deleteMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, ele := range movies {
		if ele.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, ele := range movies {
		if ele.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main() {
	fmt.Println("LEARNING GORILLA MUX")
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "45455", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovieByID).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/delete/{id}", deleteMovies).Methods("PUT")

	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}
}
