package controllers

import (
	"net/http"

	"ticket-system/config"
	"ticket-system/models"
	"ticket-system/utils"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {

	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existing models.User

	if err := config.DB.Where("email = ?", req.Email).First(&existing).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "email already exists",
		})
		return
	}

	hash, _ := utils.HashPassword(req.Password)

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hash,
	}

	config.DB.Create(&user)

	c.JSON(http.StatusCreated, gin.H{
		"message": "User Registered",
	})
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {

	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	if err := config.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid Credentials",
		})
		return
	}

	if !utils.CheckPassword(req.Password, user.Password) {

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid Credentials",
		})

		return
	}

	token, _ := utils.GenerateJWT(user.ID)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}