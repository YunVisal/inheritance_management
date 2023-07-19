package controllers

import (
	"inheritence_management/models"
	"net/http"
	"os"

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
// @Param			file formData file true "file"
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

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(os.Getenv("BUCKET_NAME")),
		ACL:    aws.String("public-read"),
		Key:    aws.String(fileHeader.Filename),
		Body:   file,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filepath := "https://" + os.Getenv("BUCKET_NAME") + "." + "s3-" + os.Getenv("AWS_REGION") + ".amazonaws.com/" + fileHeader.Filename

	image := models.Image{}
	image.ImageUrl = filepath
	image.UserId = int(user_id)

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
// @Success 		200 {array} models.Image
// @Router 			/api/image [get]
// @Security Bearer
func GetImage(c *gin.Context) {
	user_id, err := token.ExtractUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	image, err := models.GetImage(user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": image})
}
