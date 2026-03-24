package usecase

import (
	"context"
	"time"
	"workshop-storage-api-docs/internal/entity"
	"workshop-storage-api-docs/internal/model"
	"workshop-storage-api-docs/internal/repository"

	"github.com/google/uuid"
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
		CreatedAt: time.Now(),
	}

	err := r.restaurantRepository.CreateRestaurant(ctx, restaurant)
	if err != nil {
		return nil, err
	}

	response := model.ToRestaurantResponse(restaurant)
	return &response, nil
}

func (r *RestaurantUsecase) GetRestaurants(ctx context.Context, pagination model.Pagination) ([]model.RestaurantResponse, error) {
	restaurants, err := r.restaurantRepository.GetRestaurants(ctx, pagination)
	if err != nil {
		return nil, err
	}

	responses := model.ToRestaurantResponses(restaurants)
	return responses, nil
}

func (r *RestaurantUsecase) DeleteRestaurant(ctx context.Context, id uuid.UUID) error {
	return r.restaurantRepository.DeleteRestaurant(ctx, id)
}

func (r *RestaurantUsecase) EditRestaurant(ctx context.Context, id uuid.UUID, edit model.EditRestaurant) error {
	return r.restaurantRepository.EditRestaurant(ctx, id, edit)
}
