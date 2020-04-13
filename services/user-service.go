package services

import (
	"errors"

	"github.com/hcabnettek/filmapi/models"
	"github.com/hcabnettek/filmapi/repositories"
)

// UserService interface
type UserService interface {
	Validate(post *models.User) error
	Create(post *models.User) (*models.User, error)
	FindAll() ([]models.User, error)
}

type service struct{}

var (
	repo repositories.UserRepository
)

// NewUserService returns a UserService
func NewUserService(repository repositories.UserRepository) UserService {
	repo = repository
	return &service{}
}

func (*service) Validate(user *models.User) error {
	if user == nil {
		err := errors.New("The post is empty")
		return err
	}
	if user.Name == "" {
		err := errors.New("The name is empty")
		return err
	}
	return nil
}

func (*service) Create(user *models.User) (*models.User, error) {
	//post.ID = rand.Int63()
	//return repo.Save(post)
	return nil, nil
}

func (*service) FindAll() ([]models.User, error) {
	return repo.FindAll()
}
