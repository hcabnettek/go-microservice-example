package repositories

import "github.com/hcabnettek/filmapi/models"

// UserRepository generic interface
type UserRepository interface {
	Save(post *models.User) (*models.User, error)
	FindAll() ([]models.User, error)
}
