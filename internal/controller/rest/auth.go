package rest

import (
	"net/http"
	"os"
	"workshop-storage-api-docs/internal/model"
	"workshop-storage-api-docs/pkg/oauth"

	"github.com/gin-gonic/gin"
)

func (r *V1) Register(c *gin.Context) {
	var registerRequest model.UserRegister

	err := c.ShouldBindBodyWithJSON(&registerRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	err = r.validator.Struct(registerRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	ctx := c.Request.Context()

	err = r.usecase.AuthUsecase.Register(ctx, registerRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (r *V1) Login(c *gin.Context) {
	var loginRequest model.UserLogin

	err := c.ShouldBindBodyWithJSON(&loginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	err = r.validator.Struct(loginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	ctx := c.Request.Context()
	token, err := r.usecase.AuthUsecase.Login(ctx, loginRequest)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, token)
}

func (r *V1) LoginWithGoogle(c *gin.Context) {
	state := oauth.GenerateRandomState()

	c.SetCookie(
		"google_state",
		state,
		3600,
		"/",
		"",
		os.Getenv("APP_ENV") != "development",
		true,
	)

	url := r.usecase.AuthUsecase.GenerateGoogleAuthLink(state)
	c.Redirect(http.StatusFound, url)
}

func (r *V1) HandleGoogleCallback(c *gin.Context) {
	state := c.Query("state")
	code := c.Query("code")

	oauth2state, err := c.Cookie("google_state")
	if err != nil || state != oauth2state {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid oauth state"})
		return
	}

	ctx := c.Request.Context()
	token, err := r.usecase.AuthUsecase.HandleCallback(ctx, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, token)
}
