package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       int       `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: 1, ISBN: "43244", Title: "PK", Director: &Director{FirstName: "Sam", LastName: "D"}})
	movies = append(movies, Movie{ID: 2, ISBN: "43234", Title: "John Wick", Director: &Director{FirstName: "Ram", LastName: "K"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Print("Listening at PORT 8080\n")

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Error occured %s", err)
	}
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appication/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for idx, item := range movies {
		if item.ID == id {
			json.NewEncoder(w).Encode(movies[idx])
			break
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = movies[len(movies)-1].ID + 1
	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)

}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, item := range movies {
		if item.ID == id {
			// movies = append(movies[:idx], movies[idx+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = id
			movies[id-1] = movie
			// movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for idx, item := range movies {
		if id == item.ID {
			movies = append(movies[:idx], movies[idx+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}
