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

// GetUsers godoc
// @Summary Ambil semua user
// @Description Menampilkan semua data user
// @Tags User
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} map[string]interface{} "Gagal mengambil data"
// @Router /users [get]
func (ctrl *UserController) GetUsers(c *gin.Context) {
	users, err := ctrl.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}

	c.JSON(http.StatusOK, users)
}


// CreateUsers godoc
// @Summary Buat user baru
// @Description Membuat data user
// @Tags User
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} map[string]interface{} "Gagal membuat data user"
// @Router /users/create [post]
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


// GetUsersByID godoc
// @Summary Ambil data user
// @Description Ambil data user bedasarkan id
// @Tags User
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} map[string]interface{} "Gagal mengambil data user"
// @Router /users/:id [get]
func (ctrl *UserController) GetUsersByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	users, err := ctrl.service.GetUsersById(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data!"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// UpdateUsers godoc
// @Summary Ubah data user
// @Description Memperbarui data user bedasarkan id
// @Tags User
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} map[string]interface{} "Gagal mengubah data user"
// @Router /users/update/:id [put]
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

// DeleteUsers godoc
// @Summary Hapus data user
// @Description Menghapus data user bedasarkan id
// @Tags User
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} map[string]interface{} "Gagal menghapus data user"
// @Router /users/delete/:id [delete]
func (ctrl *UserController) DeleteUsers(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := ctrl.service.DeleteUsers(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal hapus data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Berhasil hapus data"})
}

// Profile godoc
// @Summary Ambil data user
// @Description Mengambil data user bedasarkan token
// @Tags User
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} map[string]interface{} "Gagal mengambil data user"
// @Router /users/profile [get]
func (ctrl *UserController) Profile(c *gin.Context) {
	ID, _ := c.Get("id")
	Email, _ := c.Get("email")
	Role, _ := c.Get("role")

	c.JSON(http.StatusOK, gin.H{
		"message": "Data user dari token",
		"id": ID,
		"email": Email,
		"role": Role,
	})
}

