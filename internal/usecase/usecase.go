package usecase

import (
	"workshop-storage-api-docs/internal/repository"
	"workshop-storage-api-docs/pkg/bcrypt"
	"workshop-storage-api-docs/pkg/jwt"

	storage_go "github.com/supabase-community/storage-go"
	"golang.org/x/oauth2"
)

type Usecase struct {
	AuthUsecase       IAuthUsecase
	RestaurantUsecase IRestaurantUsecase
	ItemUsecase       IItemUsecase
	MediaUsecase      IMediaUsecase
}

func NewUsecase(jwt jwt.JWT, bcrypt bcrypt.IBcrypt, oauth *oauth2.Config, storageClient *storage_go.Client, repository *repository.Repository) *Usecase {
	return &Usecase{
		AuthUsecase:       NewAuthUsecase(jwt, bcrypt, oauth, repository.UserRepository),
		RestaurantUsecase: NewRestaurantUsecase(repository.RestaurantRepository),
		ItemUsecase:       NewItemUsecase(repository.ItemRepository),
		MediaUsecase:      NewMediaUsecase(storageClient),
	}
}
