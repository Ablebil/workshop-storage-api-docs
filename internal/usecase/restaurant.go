package usecase

import (
	"context"
	"errors"
	"time"
	"workshop-storage-api-docs/internal/entity"
	"workshop-storage-api-docs/internal/model"
	"workshop-storage-api-docs/internal/repository"
	"workshop-storage-api-docs/pkg/apierror"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IRestaurantUsecase interface {
	CreateRestaurant(ctx context.Context, createRestaurant model.CreateRestaurant) (*model.RestaurantResponse, error)
	GetRestaurants(ctx context.Context, pagination model.Pagination) ([]model.RestaurantResponse, error)
	DeleteRestaurant(ctx context.Context, id uuid.UUID) error
	EditRestaurant(ctx context.Context, id uuid.UUID, edit model.EditRestaurant) error
}

type RestaurantUsecase struct {
	restaurantRepository repository.IRestaurantRepository
}

func NewRestaurantUsecase(restaurantRepository repository.IRestaurantRepository) *RestaurantUsecase {
	return &RestaurantUsecase{restaurantRepository}
}

func (r *RestaurantUsecase) CreateRestaurant(ctx context.Context, createRestaurant model.CreateRestaurant) (*model.RestaurantResponse, error) {
	restaurant := entity.Restaurant{
		Id:        uuid.New(),
		Name:      createRestaurant.Name,
		Location:  createRestaurant.Location,
		ImageURL:  createRestaurant.ImageURL,
		CreatedAt: time.Now(),
	}

	err := r.restaurantRepository.CreateRestaurant(ctx, restaurant)
	if err != nil {
		return nil, apierror.New(500, "failed to create restaurant")
	}

	response := model.ToRestaurantResponse(restaurant)
	return &response, nil
}

func (r *RestaurantUsecase) GetRestaurants(ctx context.Context, pagination model.Pagination) ([]model.RestaurantResponse, error) {
	restaurants, err := r.restaurantRepository.GetRestaurants(ctx, pagination)
	if err != nil {
		return nil, apierror.New(500, "failed to get restaurants")
	}

	responses := model.ToRestaurantResponses(restaurants)
	return responses, nil
}

func (r *RestaurantUsecase) DeleteRestaurant(ctx context.Context, id uuid.UUID) error {
	err := r.restaurantRepository.DeleteRestaurant(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apierror.New(404, "restaurant not found")
		}

		return apierror.New(500, "failed to delete restaurant")
	}

	return nil
}

func (r *RestaurantUsecase) EditRestaurant(ctx context.Context, id uuid.UUID, edit model.EditRestaurant) error {
	err := r.restaurantRepository.EditRestaurant(ctx, id, edit)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apierror.New(404, "restaurant not found")
		}

		return apierror.New(500, "failed to edit restaurant")
	}

	return nil
}
