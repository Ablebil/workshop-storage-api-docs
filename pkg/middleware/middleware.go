package middleware

import (
	"workshop-storage-api-docs/pkg/jwt"

	"github.com/gin-gonic/gin"
)

type IMiddleware interface {
	Authentication(c *gin.Context)
	Authorization(roles ...string) gin.HandlerFunc
}

type Middleware struct {
	jwt jwt.IJWT
}

func NewMiddleware(jwt jwt.IJWT) IMiddleware {
	return &Middleware{
		jwt: jwt,
	}
}
