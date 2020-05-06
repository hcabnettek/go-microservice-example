package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/hcabnettek/filmapi/errors"
	"github.com/hcabnettek/filmapi/services"
)

type moviecontroller struct{}

var (
	movieService services.MovieService
)

// MovieController interface
type MovieController interface {
	GetMovies(response http.ResponseWriter, request *http.Request)
	AddMovie(response http.ResponseWriter, request *http.Request)
}

// NewMovieController constructor function returns MovieController interface
func NewMovieController(service services.MovieService) MovieController {
	movieService = service
	return &moviecontroller{}
}

func (*moviecontroller) GetMovies(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	movies, err := movieService.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error getting the movies"})
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(movies)
}

func (*moviecontroller) AddMovie(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	/*posts, err := postService.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error getting the posts"})
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts) */
}
