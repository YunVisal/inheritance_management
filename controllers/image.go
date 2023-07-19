package controllers

import (
	"inheritence_management/models"
	"net/http"
	"os"
	"strings"
	"time"

	"inheritence_management/utils/token"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
)

// UploadImage 		godoc
// @Summary 		image
// @Description 	upload image
// @Tags			image
// @Produce 		json
// @Param			file formData file true "File"
// @Param			packageId formData string true "Package Id"
// @Success 		200 {object} models.Image
// @Router 			/api/image [post]
// @Security Bearer
func UploadImage(c *gin.Context) {
	user_id, err := token.ExtractUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sess := session.Must(session.NewSession())
	uploader := s3manager.NewUploader(sess)
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fileName := fileHeader.Filename + "_" + strings.Join(strings.Split(time.Now().Format("2006-01-02 15:04:05"), " "), "_")
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(os.Getenv("BUCKET_NAME")),
		ACL:    aws.String("public-read"),
		Key:    aws.String(fileName),
		Body:   file,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filepath := "https://" + os.Getenv("BUCKET_NAME") + "." + "s3-" + os.Getenv("AWS_REGION") + ".amazonaws.com/" + fileName

	image := models.Image{}
	image.ImageUrl = filepath
	image.UserId = int(user_id)
	image.PackageId = c.Request.FormValue("packageId")

	res, err := image.SaveImage()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": res})
}

// GetImage 		godoc
// @Summary 		image
// @Description 	get upload image of user and family member
// @Tags			image
// @Produce 		json
// @Param 			packageId query string false "Package ID"
// @Success 		200 {array} models.Image
// @Router 			/api/image [get]
// @Security Bearer
func GetImage(c *gin.Context) {
	user_id, err := token.ExtractUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var images []models.Image
	if c.Query("packageId") == "" {
		images, err = models.GetImage(user_id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		images, err = models.GetImageByPackageId(user_id, c.Query("packageId"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": images})
}
