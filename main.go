package main

import(
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"ID"`
	Isbn string `json:"isbn"`
	title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

var Movies []Movie 

func main(){
	r :=mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn:"438227", title:"Movie One", Director: &Director{FirstName: "Stanley", LastName: "Kubrick"}})
	movies = append(movies, Movie{ID: "2", Isbn:"438227", title:"Movie Two", Director: &Director{FirstName: "Winston", LastName: "Churchill"}})
	r.HandleFunc("/movies", getMovies).methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).methods("GET")
	r.HandleFunc("/movies", createMovie).methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie ).methods("PUT")
	r.HandleFunc("movies/{id}", deleteMovie).methods("DELETE")

	fmt.PrintF("Starting the server at port: 8000")

	log.Fatal(http.ListenAndServe(":8000", r))
}