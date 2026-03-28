package usecase

import (
	"context"
	"errors"
	"workshop-storage-api-docs/internal/entity"
	"workshop-storage-api-docs/internal/model"
	"workshop-storage-api-docs/internal/repository"
	"workshop-storage-api-docs/pkg/apierror"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IItemUsecase interface {
	GetRestaurantItems(ctx context.Context, pagination model.Pagination, restaurantId uuid.UUID) ([]model.ItemResponse, error)
	CreateItem(ctx context.Context, restaurantId uuid.UUID, createItem model.CreateItem) (*model.ItemResponse, error)
	DeleteItem(ctx context.Context, id uuid.UUID) error
	EditItem(ctx context.Context, id uuid.UUID, edit model.EditItem) error
}

type ItemUsecase struct {
	itemRepository repository.IItemRepository
}

func NewItemUsecase(itemRepository repository.IItemRepository) *ItemUsecase {
	return &ItemUsecase{
		itemRepository: itemRepository,
	}
}

func (u *ItemUsecase) GetRestaurantItems(ctx context.Context, pagination model.Pagination, restaurantId uuid.UUID) ([]model.ItemResponse, error) {
	items, err := u.itemRepository.GetRestaurantItems(ctx, pagination, restaurantId)
	if err != nil {
		return nil, apierror.New(500, "failed to get restaurant items")
	}

	responses := model.ToItemResponses(items)
	return responses, nil
}

func (u *ItemUsecase) CreateItem(ctx context.Context, restaurantId uuid.UUID, createItem model.CreateItem) (*model.ItemResponse, error) {
	item := entity.Item{
		Id:           uuid.New(),
		RestaurantId: restaurantId,
		Name:         createItem.Name,
		Price:        createItem.Price,
		Available:    createItem.Available,
	}

	err := u.itemRepository.CreateItem(ctx, item)
	if err != nil {
		return nil, apierror.New(500, "failed to create item")
	}

	response := model.ToItemResponse(item)
	return &response, nil
}

func (u *ItemUsecase) DeleteItem(ctx context.Context, id uuid.UUID) error {
	err := u.itemRepository.DeleteItem(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apierror.New(404, "item not found")
		}

		return apierror.New(500, "failed to delete item")
	}

	return nil
}

func (u *ItemUsecase) EditItem(ctx context.Context, id uuid.UUID, edit model.EditItem) error {
	err := u.itemRepository.EditItem(ctx, id, edit)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apierror.New(404, "item not found")
		}

		return apierror.New(500, "failed to edit item")
	}

	return nil
}
