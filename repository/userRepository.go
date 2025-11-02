package repository

import (
	"go-crud/database"
	"go-crud/models"
)

type UserRepository struct {}

func (r *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User

	result := database.DB.Find(&users)
	return users, result.Error
}

func (r *UserRepository) FindByID(id uint) (models.User, error) {
	var users models.User
	result := database.DB.First(&users, id)
	return users, result.Error
}

func (r *UserRepository) Create(users *models.User) error {
	return database.DB.Create(users).Error
}

func (r *UserRepository) Update(users *models.User) error {
	return database.DB.Save(users).Error
}

func (r *UserRepository) Delete(id uint) error {
	return database.DB.Delete(&models.User{}, id).Error
}

