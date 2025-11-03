package controllers

import (
	"go-crud/models"
	"go-crud/services"
	"go-crud/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service services.UserService
}

func (ctrl *UserController) GetUsers(c *gin.Context) {
	users, err := ctrl.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (ctrl *UserController) CreateUsers(c *gin.Context) {
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

	if err := ctrl.service.CreateUsers(&users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (ctrl *UserController) GetUsersByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	users, err := ctrl.service.GetUsersById(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data!"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (ctrl *UserController) UpdateUsers(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var users models.User
	if err := c.ShouldBindJSON(&users); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users.Id = uint(id)

	if err := ctrl.service.UpdateUsers(&users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal update data"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (ctrl *UserController) DeleteUsers(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := ctrl.service.DeleteUsers(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal hapus data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Berhasil hapus data"})
}

