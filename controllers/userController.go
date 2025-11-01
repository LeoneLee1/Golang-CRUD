package controllers

import (
	"go-crud/database"
	"go-crud/models"
	"go-crud/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User

	database.DB.Find(&users)

	c.JSON(http.StatusOK, users)
}

func CreateUsers(c *gin.Context) {
	var users models.User

	if err := c.ShouldBindJSON(&users); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := utils.HashPassword(users.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal hash password!"})
		return
	}

	users.Password = hashedPassword

	database.DB.Create(&users)
	c.JSON(http.StatusOK, users)
}

func ShowUsers(c *gin.Context) {
	var users models.User

	if err := database.DB.Where("id = ?", c.Param("id")).First(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!"})
		return
	}

	c.JSON(http.StatusOK,users)
}

func UpdateUsers(c *gin.Context) {
	var users models.User

	if err := database.DB.Where("id = ?", c.Param("id")).First(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!"})
		return
	}

	var input models.ValidateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error:": "Semua field harus di isi!"})
		return
	}

	database.DB.Model(&users).Updates(input)

	c.JSON(http.StatusOK, users)
}

func DeleteUsers(c *gin.Context) {
	var users models.User

	if err := database.DB.Where("id = ?", c.Param("id")).First(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!"})
		return
	}

	database.DB.Delete(&users)

	c.JSON(http.StatusOK, gin.H{"message": "Berhasil menghapus data!"})
}

