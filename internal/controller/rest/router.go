package rest

import (
	"github.com/gin-gonic/gin"
)

func NewRouter(app *gin.Engine, v1 *V1) {
	api := app.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", v1.Register)
			auth.POST("/login", v1.Login)
			auth.GET("/google/login", v1.LoginWithGoogle)
			auth.GET("/google/callback", v1.HandleGoogleCallback)
		}

		media := api.Group("/media")
		{
			media.POST("/upload", v1.IMiddleware.Authentication, v1.UploadMedia)
		}

		restaurants := api.Group("/restaurants")
		{
			restaurants.GET("", v1.IMiddleware.Authentication, v1.GetRestaurants)
			restaurants.POST("", v1.IMiddleware.Authentication, v1.IMiddleware.Authorization("admin", "user"), v1.CreateRestaurant) // image upload
			restaurants.DELETE("/:id", v1.DeleteRestaurant)
			restaurants.PATCH("/:id", v1.EditRestaurant) // image upload

			restaurantItems := restaurants.Group("/:id/items")
			{
				restaurantItems.GET("", v1.GetRestaurantItems)
				restaurantItems.POST("", v1.CreateItem) // image upload
			}
		}

		items := api.Group("/items")
		{
			items.DELETE("/:id", v1.DeleteItem)
			items.PATCH("/:id", v1.EditItem) // image upload
		}
	}
}
