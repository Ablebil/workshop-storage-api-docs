package rest

import (
	"errors"
	"net/http"
	"workshop-storage-api-docs/pkg/apierror"

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
		var apiErr *apierror.APIError
		if errors.As(err, &apiErr) {
			c.JSON(apiErr.Code, gin.H{"error": apiErr.Message})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, response)
}
