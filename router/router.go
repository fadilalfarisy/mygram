package router

import (
	"challenge-4/controllers"
	"challenge-4/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", controllers.CreateProduct)
		productRouter.GET("/:productId", middlewares.ProductAuthorization(), controllers.GetProductById)
		productRouter.PUT("/:productId", middlewares.AdminAuthorization(), controllers.UpdateProduct)
		productRouter.DELETE("/:productId", middlewares.AdminAuthorization(), controllers.DeleteProduct)
	}

	return r
}
