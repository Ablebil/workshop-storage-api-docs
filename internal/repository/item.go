package repository

import (
	"context"
	"workshop-storage-api-docs/internal/entity"
	"workshop-storage-api-docs/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IItemRepository interface {
	GetRestaurantItems(ctx context.Context, pagination model.Pagination, restaurantID uuid.UUID) ([]entity.Item, error)
	CreateItem(ctx context.Context, item entity.Item) error
	DeleteItem(ctx context.Context, id uuid.UUID) error
	EditItem(ctx context.Context, id uuid.UUID, edit model.EditItem) error
}

type ItemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{db}
}

func (r *ItemRepository) GetRestaurantItems(ctx context.Context, pagination model.Pagination, restaurantID uuid.UUID) ([]entity.Item, error) {
	Item, err := gorm.G[entity.Item](r.db).
		Limit(pagination.Limit).
		Offset(pagination.Offset()).
		Order("created_at DESC").
		Where("restaurant_id = ?", restaurantID).
		Find(ctx)

	if err != nil {
		return nil, err
	}

	return Item, nil
}

func (r *ItemRepository) CreateItem(ctx context.Context, item entity.Item) error {
	err := gorm.G[entity.Item](r.db).Create(ctx, &item)
	if err != nil {
		return err
	}

	return nil
}

func (r *ItemRepository) DeleteItem(ctx context.Context, id uuid.UUID) error {
	rows, err := gorm.G[entity.Item](r.db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		return err
	}

	if rows == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *ItemRepository) EditItem(ctx context.Context, id uuid.UUID, edit model.EditItem) error {
	result := r.db.WithContext(ctx).Model(&entity.Item{}).
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
