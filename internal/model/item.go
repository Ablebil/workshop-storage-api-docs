package model

import (
	"workshop-storage-api-docs/internal/entity"

	"github.com/google/uuid"
)

type ItemResponse struct {
	Id           uuid.UUID `json:"id"`
	RestaurantId uuid.UUID `json:"restaurant_id"`
	Name         string    `json:"name"`
	Price        float64   `json:"price"`
	Available    bool      `json:"available"`
	ImageURL     string    `json:"image_url"`
}

func ToItemResponse(item entity.Item) ItemResponse {
	return ItemResponse{
		Id:           item.Id,
		RestaurantId: item.RestaurantId,
		Name:         item.Name,
		Price:        item.Price,
		Available:    item.Available,
		ImageURL:     item.ImageURL,
	}
}

func ToItemResponses(items []entity.Item) []ItemResponse {
	var responses []ItemResponse
	for _, item := range items {
		responses = append(responses, ToItemResponse(item))
	}

	return responses
}

type CreateItem struct {
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Available bool    `json:"available"`
	ImageURL  string  `json:"image_url"`
}

type EditItem struct {
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Available *bool   `json:"available"`
	ImageURL  string  `json:"image_url"`
}

func (e *EditItem) ToMap() map[string]any {
	updates := map[string]any{}

	if e.Name != "" {
		updates["name"] = e.Name
	}
	if e.Price != 0 {
		updates["price"] = e.Price
	}
	if e.Available != nil {
		updates["available"] = *e.Available
	}
	if e.ImageURL != "" {
		updates["image_url"] = e.ImageURL
	}

	return updates
}
