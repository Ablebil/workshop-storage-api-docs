package model

import (
	"workshop-storage-api-docs/internal/entity"

	"github.com/google/uuid"
)

type CreateRestaurant struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

type EditRestaurant struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

func (e *EditRestaurant) ToMap() map[string]any {
	updates := map[string]any{}

	if e.Name != "" {
		updates["name"] = e.Name
	}
	if e.Location != "" {
		updates["location"] = e.Location
	}

	return updates
}

type RestaurantResponse struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Location string    `json:"location"`
}

func ToRestaurantResponse(restaurant entity.Restaurant) RestaurantResponse {
	return RestaurantResponse{
		Id:       restaurant.Id,
		Name:     restaurant.Name,
		Location: restaurant.Location,
	}
}

func ToRestaurantResponses(restaurants []entity.Restaurant) []RestaurantResponse {
	var responses []RestaurantResponse
	for _, restaurant := range restaurants {
		responses = append(responses, ToRestaurantResponse(restaurant))
	}

	return responses
}
