package models

type User struct {
	Id       uint   `json:"id" gorm:"PrimaryKey;autoIncrement"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Role     string `json:"role"`
}