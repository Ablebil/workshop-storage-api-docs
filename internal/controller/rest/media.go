package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *V1) UploadMedia(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to get image from request"})
		return
	}

	uploadType := c.DefaultPostForm("type", "others")

	ctx := c.Request.Context()

	response, err := r.usecase.MediaUsecase.UploadImage(ctx, file, uploadType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
