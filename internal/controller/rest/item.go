package rest

import (
	"errors"
	"net/http"
	"strconv"
	"workshop-storage-api-docs/internal/model"
	"workshop-storage-api-docs/pkg/apierror"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (r *V1) GetRestaurantItems(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid restaurant id"})
		return
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page parameter"})
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit parameter"})
		return
	}

	pagination := model.Pagination{
		Page:  page,
		Limit: limit,
	}
	pagination.Check()

	ctx := c.Request.Context()
	items, err := r.usecase.ItemUsecase.GetRestaurantItems(ctx, pagination, id)
	if err != nil {
		var apiErr *apierror.APIError
		if errors.As(err, &apiErr) {
			c.JSON(apiErr.Code, gin.H{"error": apiErr.Message})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		return
	}

	response := model.PaginatedResponse[model.ItemResponse]{
		Data:       items,
		Pagination: pagination,
	}

	c.JSON(http.StatusOK, response)
}

func (r *V1) CreateItem(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid restaurant id"})
		return
	}

	var create model.CreateItem
	err = c.ShouldBindBodyWithJSON(&create)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	ctx := c.Request.Context()
	item, err := r.usecase.ItemUsecase.CreateItem(ctx, id, create)
	if err != nil {
		var apiErr *apierror.APIError
		if errors.As(err, &apiErr) {
			c.JSON(apiErr.Code, gin.H{"error": apiErr.Message})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		return
	}

	c.JSON(http.StatusCreated, item)
}

func (r *V1) DeleteItem(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid item id"})
		return
	}

	ctx := c.Request.Context()
	err = r.usecase.ItemUsecase.DeleteItem(ctx, id)
	if err != nil {
		var apiErr *apierror.APIError
		if errors.As(err, &apiErr) {
			c.JSON(apiErr.Code, gin.H{"error": apiErr.Message})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (r *V1) EditItem(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid item id"})
		return
	}

	var edit model.EditItem
	err = c.ShouldBindBodyWithJSON(&edit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	ctx := c.Request.Context()
	err = r.usecase.ItemUsecase.EditItem(ctx, id, edit)
	if err != nil {
		var apiErr *apierror.APIError
		if errors.As(err, &apiErr) {
			c.JSON(apiErr.Code, gin.H{"error": apiErr.Message})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, nil)
}
