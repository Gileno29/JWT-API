package repository

import "jwt-api/internernal/models"

type UserInterface interface {
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(user *models.User) error
	GetUser(id uint64) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetAll() ([]models.User, error)
}
