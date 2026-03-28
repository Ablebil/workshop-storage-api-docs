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

func (r *V1) GetRestaurants(c *gin.Context) {
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
	restaurants, err := r.usecase.RestaurantUsecase.GetRestaurants(ctx, pagination)
	if err != nil {
		var apiErr *apierror.APIError
		if errors.As(err, &apiErr) {
			c.JSON(apiErr.Code, gin.H{"error": apiErr.Message})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		return
	}

	response := model.PaginatedResponse[model.RestaurantResponse]{
		Data:       restaurants,
		Pagination: pagination,
	}

	c.JSON(http.StatusOK, response)
}

func (r *V1) CreateRestaurant(c *gin.Context) {
	var createRestaurant model.CreateRestaurant

	err := c.ShouldBindBodyWithJSON(&createRestaurant)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	ctx := c.Request.Context()

	restaurant, err := r.usecase.RestaurantUsecase.CreateRestaurant(ctx, createRestaurant)
	if err != nil {
		var apiErr *apierror.APIError
		if errors.As(err, &apiErr) {
			c.JSON(apiErr.Code, gin.H{"error": apiErr.Message})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		return
	}

	c.JSON(http.StatusCreated, restaurant)
}

func (r *V1) DeleteRestaurant(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid restaurant id"})
		return
	}

	ctx := c.Request.Context()

	err = r.usecase.RestaurantUsecase.DeleteRestaurant(ctx, id)
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

func (r *V1) EditRestaurant(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid restaurant id"})
		return
	}

	var edit model.EditRestaurant
	err = c.ShouldBindBodyWithJSON(&edit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	ctx := c.Request.Context()
	err = r.usecase.RestaurantUsecase.EditRestaurant(ctx, id, edit)
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
