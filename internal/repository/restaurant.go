package repository

import (
	"context"
	"workshop-storage-api-docs/internal/entity"
	"workshop-storage-api-docs/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IRestaurantRepository interface {
	CreateRestaurant(ctx context.Context, restaurant entity.Restaurant) error
	GetRestaurants(ctx context.Context, pagination model.Pagination) ([]entity.Restaurant, error)
	DeleteRestaurant(ctx context.Context, id uuid.UUID) error
	EditRestaurant(ctx context.Context, id uuid.UUID, edit model.EditRestaurant) error
}

type RestaurantRepository struct {
	db *gorm.DB
}

func NewRestaurantRepository(db *gorm.DB) *RestaurantRepository {
	return &RestaurantRepository{db}
}

func (r *RestaurantRepository) CreateRestaurant(ctx context.Context, restaurant entity.Restaurant) error {
	err := gorm.G[entity.Restaurant](r.db).Create(ctx, &restaurant)
	if err != nil {
		return err
	}

	return nil
}

func (r *RestaurantRepository) GetRestaurants(ctx context.Context, pagination model.Pagination) ([]entity.Restaurant, error) {
	restaurants, err := gorm.G[entity.Restaurant](r.db).
		Limit(pagination.Limit).
		Offset(pagination.Offset()).
		Order("created_at DESC").
		Find(ctx)
	if err != nil {
		return nil, err
	}

	return restaurants, nil
}

func (r *RestaurantRepository) DeleteRestaurant(ctx context.Context, id uuid.UUID) error {
	rows, err := gorm.G[entity.Restaurant](r.db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		return err
	}

	if rows == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *RestaurantRepository) EditRestaurant(ctx context.Context, id uuid.UUID, edit model.EditRestaurant) error {
	result := r.db.WithContext(ctx).Model(&entity.Restaurant{}).
		Where("id = ?", id).
		Updates(edit.ToMap())

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
