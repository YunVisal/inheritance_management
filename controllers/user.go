package controllers

import (
	"net/http"

	"inheritence_management/models"

	"github.com/gin-gonic/gin"
)

// GetAllUser 		godoc
// @Summary 		users
// @Description 	get all users by search term
// @Tags			users
// @Produce 		json
// @Param term query string false "Search Term"
// @Success 		200 {array} models.User
// @Router 			/api/user [get]
// @Security Bearer
func GetAllUser(c *gin.Context) {
	u, err := models.GetUsers(c.Query("term"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}
