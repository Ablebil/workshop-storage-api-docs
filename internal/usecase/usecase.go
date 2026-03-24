package usecase

import (
	"workshop-storage-api-docs/internal/repository"
	"workshop-storage-api-docs/pkg/bcrypt"
	"workshop-storage-api-docs/pkg/jwt"

	"golang.org/x/oauth2"
)

type Usecase struct {
	AuthUsecase       IAuthUsecase
	RestaurantUsecase IRestaurantUsecase
	ItemUsecase       IItemUsecase
}

func NewUsecase(jwt jwt.JWT, bcrypt bcrypt.IBcrypt, oauth *oauth2.Config, repository *repository.Repository) *Usecase {
	return &Usecase{
		AuthUsecase:       NewAuthUsecase(jwt, bcrypt, oauth, repository.UserRepository),
		RestaurantUsecase: NewRestaurantUsecase(repository.RestaurantRepository),
		ItemUsecase:       NewItemUsecase(repository.ItemRepository),
	}
}
