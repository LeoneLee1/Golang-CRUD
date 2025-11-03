package controllers

import (
	"go-crud/database"
	"go-crud/helpers"
	"go-crud/models"
	"go-crud/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {}


// Login godoc
// @Summary Login user
// @Description Login menggunakan email dan password
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body object{email=string,password=string} true "Login request"
// @Success 200 {object} map[string]interface{} "Login berhasil"
// @Failure 400 {object} map[string]interface{} "Input tidak valid"
// @Failure 401 {object} map[string]interface{} "Email atau password salah"
// @Router /login [post]
func (a *AuthController) Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}

	var users models.User
	if err := database.DB.Where("email = ?", input.Email).First(&users).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email tidak ditemukan"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password salah"})
		return
	}

	token, err := helpers.CreateToken(users.Id ,users.Email, users.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login berhasil",
		"token":   token,
	})
}

// Register godoc
// @Summary Register user baru
// @Description Mendaftar akun baru
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body object{name=string,email=string,password=string} true "Register request"
// @Success 200 {object} map[string]interface{} "Akun berhasil dibuat"
// @Failure 400 {object} map[string]interface{} "Input tidak valid"
// @Router /register [post]
func (a *AuthController) Register(c *gin.Context) {
	var input struct {
		Name string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}

	var users models.User

	hashedPassword, err := utils.HashPassword(input.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal hash password!"})
		return
	}

	users.Name = input.Name
	users.Email = input.Email
	users.Password = hashedPassword

	database.DB.Create(&users)

	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil membuat akun",
		"data": users,
	})
}