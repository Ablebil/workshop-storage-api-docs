package rest

import (
	"errors"
	"net/http"
	"strconv"
	"workshop-storage-api-docs/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *V1) GetRestaurants(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
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
		c.AbortWithStatus(http.StatusInternalServerError)
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
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()

	restaurant, err := r.usecase.RestaurantUsecase.CreateRestaurant(ctx, createRestaurant)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, restaurant)
}

func (r *V1) DeleteRestaurant(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()

	err = r.usecase.RestaurantUsecase.DeleteRestaurant(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (r *V1) EditRestaurant(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var edit model.EditRestaurant
	err = c.ShouldBindBodyWithJSON(&edit)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()
	err = r.usecase.RestaurantUsecase.EditRestaurant(ctx, id, edit)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, nil)
}
