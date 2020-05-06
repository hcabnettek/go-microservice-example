package repositories

import "github.com/hcabnettek/filmapi/models"

// MovieRepository generic interface
type MovieRepository interface {
	Save(post *models.Movie) (*models.Movie, error)
	FindAll() ([]models.Movie, error)
}
