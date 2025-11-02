package models

type User struct {
	Id       int    `json:"id" gorm:"PrimaryKey;autoIncrement"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Role     string `json:"role"`
}

type ValidateUserInput struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Role  string `json:"role" binding:"required"`
}