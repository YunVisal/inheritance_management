package controllers

import (
	"net/http"

	"inheritence_management/models"
	"inheritence_management/utils/token"

	"github.com/gin-gonic/gin"
)

type RegisterDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterSuccessMessage struct {
	Message string `json:"message" default:"registration success"`
}

type LoginDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginSuccessMessage struct {
	Token string `json:"token"`
}

// Register 		godoc
// @Summary 		register user
// @Description 	register new user
// @Tags			auth
// @Accept 			json
// @Produce 		json
// @Param			Body body RegisterDTO true "User Info"
// @Success 		200 {object} RegisterSuccessMessage
// @Router 			/api/auth/register [post]
func Register(c *gin.Context) {
	var dto RegisterDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username and Password is required"})
		return
	}

	u := models.User{}
	u.Username = dto.Username
	u.Password = dto.Password

	_, err := u.SaveUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var successMsg RegisterSuccessMessage
	successMsg.Message = "registration success"
	c.JSON(http.StatusOK, successMsg)
}

// Login 			godoc
// @Summary 		login
// @Description 	login user
// @Tags			auth
// @Accept 			json
// @Produce 		json
// @Param			Body body LoginDTO true "User Info"
// @Success 		200 {object} LoginSuccessMessage
// @Router 			/api/auth/login [post]
func Login(c *gin.Context) {
	var dto LoginDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username and Password is required"})
		return
	}

	u := models.User{}
	u.Username = dto.Username
	u.Password = dto.Password

	token, err := models.VerifyUser(u.Username, u.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username and Password is incorrect"})
		return
	}

	var successMsg LoginSuccessMessage
	successMsg.Token = token
	c.JSON(http.StatusOK, successMsg)
}

// GetProfile 		godoc
// @Summary 		profile
// @Description 	get user profile
// @Tags			profile
// @Produce 		json
// @Success 		200 {object} models.User
// @Router 			/api/profile [get]
// @Security Bearer
func CurrentUser(c *gin.Context) {
	user_id, err := token.ExtractUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := models.GetUserByID(user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}
