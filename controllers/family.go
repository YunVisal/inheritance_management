package controllers

import (
	"net/http"

	"inheritence_management/models"
	"inheritence_management/utils/token"

	"github.com/gin-gonic/gin"
)

type AddRelationshipDTO struct {
	RelativeId int `json:"relativeId" binding:"required"`
}

// AddRelationship 	godoc
// @Summary 		family member
// @Description 	add family member
// @Tags			family
// @Accept 			json
// @Produce 		json
// @Param			Body body AddRelationshipDTO true "Family member Info"
// @Success 		200 {object} models.FamilyMember
// @Router 			/api/family [post]
// @Security Bearer
func AddRelationship(c *gin.Context) {
	var dto AddRelationshipDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username and Password is required"})
		return
	}

	user_id, err := token.ExtractUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	relationship := models.FamilyMember{}
	relationship.UserId = int(user_id)
	relationship.RelativeId = dto.RelativeId

	familyMember, err := relationship.SaveRelationship()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": familyMember})
}

// GetRelationship 	godoc
// @Summary 		family member
// @Description 	get family member
// @Tags			family
// @Produce 		json
// @Success 		200 {array} models.FamilyMemberResponse
// @Router 			/api/family [get]
// @Security Bearer
func GetRelationship(c *gin.Context) {
	user_id, err := token.ExtractUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	familyMember, err := models.GetFamilyMember(user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": familyMember})
}
