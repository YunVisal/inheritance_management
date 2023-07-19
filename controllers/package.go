package controllers

import (
	"inheritence_management/models"
	"net/http"
	"strconv"

	"inheritence_management/utils/token"

	"github.com/gin-gonic/gin"
)

type CreatePackageDTO struct {
	PackageName string `json:"packageName" binding:"required"`
}

type GetPackageResponse struct {
	TotalItems int               `json:"totalItems"`
	Packages   []PackageResponse `json:"packages"`
}

type PackageResponse struct {
	Id          uint   `json:"id"`
	PackageName string `json:"packageName"`
	TotalImages int    `json:"totalImages"`
	UserId      uint   `json:"userId"`
}

// CreatePackage 	godoc
// @Summary 		package
// @Description 	create package
// @Tags			package
// @Accept 			json
// @Produce 		json
// @Param			Body body CreatePackageDTO true "Package Info"
// @Success 		200 {object} models.Package
// @Router 			/api/package [post]
// @Security Bearer
func CreatePackage(c *gin.Context) {
	user_id, err := token.ExtractUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var dto CreatePackageDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username and Password is required"})
		return
	}

	_package := models.Package{}
	_package.PackageName = dto.PackageName
	_package.UserId = int(user_id)

	createdPackage, err := _package.CreatePackage()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": createdPackage})
}

// GetPackage 		godoc
// @Summary 		package
// @Description 	get package of current user
// @Tags			package
// @Produce 		json
// @Param 			userId query string false "User Id"
// @Success 		200 {object} GetPackageResponse
// @Router 			/api/package [get]
// @Security Bearer
func GetPackage(c *gin.Context) {
	user_id := uint(0)
	var err error
	user_id, err = token.ExtractUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if c.Query("userId") != "" {
		temp, err := strconv.ParseUint(c.Query("userId"), 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user_id = uint(temp)
	}

	totalPackage, err := models.GetTotalPackage(user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	allPackages, err := models.GetAllPackage(user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var packages []PackageResponse
	for i := 0; i < len(allPackages); i++ {
		totalImage := models.GetTotalImageByPackageId(int(allPackages[i].ID))
		var _package PackageResponse
		_package.Id = allPackages[i].ID
		_package.PackageName = allPackages[i].PackageName
		_package.TotalImages = totalImage
		_package.UserId = user_id
		packages = append(packages, _package)
	}

	var packageResponse GetPackageResponse
	packageResponse.TotalItems = totalPackage
	packageResponse.Packages = packages

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": packageResponse})
}
