package services

import (
	"errors"
	"go-crud/models"
	"go-crud/repository"
)

type UserService struct{
	repo repository.UserRepository
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.FindAll()
}

func (s *UserService) GetUsersById(id uint) (models.User, error) {
	users, err := s.repo.FindByID(id)

	if err != nil {
		return users, errors.New("user tidak ditemukan")
	}

	return users, nil
}

func (s *UserService) CreateUsers(users *models.User) error {
	return s.repo.Create(users)
}

func (s *UserService) UpdateUsers(users *models.User) error {
	return s.repo.Update(users)
}

func (s *UserService) DeleteUsers(id uint) error {
	return s.repo.Delete(id)
}