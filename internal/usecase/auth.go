package usecase

import (
	"context"
	"encoding/json"
	"errors"

	"workshop-storage-api-docs/internal/entity"
	"workshop-storage-api-docs/internal/model"
	"workshop-storage-api-docs/internal/repository"
	"workshop-storage-api-docs/pkg/bcrypt"
	"workshop-storage-api-docs/pkg/jwt"

	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

type IAuthUsecase interface {
	Register(ctx context.Context, param model.UserRegister) error
	Login(ctx context.Context, param model.UserLogin) (*model.TokenResponse, error)
	GenerateGoogleAuthLink(state string) string
	HandleCallback(ctx context.Context, code string) (*model.TokenResponse, error)
}

type AuthUsecase struct {
	Jwt            jwt.JWT
	Bcrypt         bcrypt.IBcrypt
	Config         *oauth2.Config
	UserRepository repository.IUserRepository
}

func NewAuthUsecase(jwt jwt.JWT, bcrypt bcrypt.IBcrypt, oAuth2 *oauth2.Config, userRepository repository.IUserRepository) *AuthUsecase {
	return &AuthUsecase{
		Jwt:            jwt,
		Bcrypt:         bcrypt,
		Config:         oAuth2,
		UserRepository: userRepository,
	}
}

func (u *AuthUsecase) Register(ctx context.Context, param model.UserRegister) error {
	hashedPassword, err := u.Bcrypt.GenerateHash(param.Password)
	if err != nil {
		return errors.New("failed to hash password")
	}
	user := entity.User{
		UserId:   uuid.New(),
		Email:    param.Email,
		Password: hashedPassword,
		Role:     model.UserRoleUser,
	}

	err = u.UserRepository.CreateUser(ctx, user)
	if err != nil {
		return errors.New("failed to create user")
	}
	return nil
}

func (u *AuthUsecase) Login(ctx context.Context, param model.UserLogin) (*model.TokenResponse, error) {
	user, err := u.UserRepository.GetUserByEmail(ctx, param.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	err = u.Bcrypt.ValidatePassword(user.Password, param.Password)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	tokenStr, err := u.Jwt.GenerateToken(user.UserId.String(), user.Role)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	token := &model.TokenResponse{Token: tokenStr}

	return token, nil
}

func (u *AuthUsecase) GenerateGoogleAuthLink(state string) string {
	return u.Config.AuthCodeURL(state)
}

func (u *AuthUsecase) HandleCallback(ctx context.Context, code string) (*model.TokenResponse, error) {
	token, err := u.Config.Exchange(ctx, code)
	if err != nil {
		return nil, errors.New("failed to exchange code for token")
	}

	client := u.Config.Client(ctx, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, errors.New("failed to get user info")
	}
	defer resp.Body.Close()

	var userInfo model.UserInfo

	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		return nil, errors.New("failed to decode user info")
	}

	user, err := u.UserRepository.GetUserByEmail(ctx, userInfo.Email)
	if err != nil {
		user = &entity.User{
			UserId:   uuid.New(),
			Email:    userInfo.Email,
			Password: "",
			Role:     model.UserRoleUser,
		}
		err = u.UserRepository.CreateUser(ctx, *user)
		if err != nil {
			return nil, errors.New("failed to create user")
		}
	}

	jwtTokenStr, err := u.Jwt.GenerateToken(user.UserId.String(), user.Role)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	jwtToken := &model.TokenResponse{Token: jwtTokenStr}

	return jwtToken, nil
}
