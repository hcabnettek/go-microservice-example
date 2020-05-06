package services

import (
	"errors"

	"github.com/hcabnettek/filmapi/models"
	"github.com/hcabnettek/filmapi/repositories"
)

// MovieService interface
type MovieService interface {
	Validate(post *models.Movie) error
	Create(post *models.Movie) (*models.Movie, error)
	FindAll() ([]models.Movie, error)
}

type movieservice struct{}

var (
	movierepo repositories.MovieRepository
)

// NewMovieService returns a MovieService
func NewMovieService(repository repositories.MovieRepository) MovieService {
	movierepo = repository
	return &movieservice{}
}

func (*movieservice) Validate(movie *models.Movie) error {
	if movie == nil {
		err := errors.New("The post is empty")
		return err
	}
	if movie.Title == "" {
		err := errors.New("The name is empty")
		return err
	}
	return nil
}

func (*movieservice) Create(movie *models.Movie) (*models.Movie, error) {
	//post.ID = rand.Int63()
	//return repo.Save(post)
	return nil, nil
}

func (*movieservice) FindAll() ([]models.Movie, error) {
	return movierepo.FindAll()
}
