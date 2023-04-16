package router

import (
	"challenge-4/controllers"
	"challenge-4/middlewares"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/user")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	photoRouter := r.Group("/photo")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/", controllers.CreatePhoto)
		photoRouter.GET("/", controllers.GetAllPhoto)
		photoRouter.GET("/:photoId", controllers.GetPhotoById)
		photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	socmedRouter := r.Group("/socmed")
	{
		socmedRouter.Use(middlewares.Authentication())
		socmedRouter.POST("/", controllers.CreateSocmed)
		socmedRouter.GET("/", controllers.GetAllSocmed)
		socmedRouter.GET("/:socmedId", controllers.GetSocmedById)
		socmedRouter.PUT("/:socmedId", middlewares.SocmedAuthorization(), controllers.UpdateSocmed)
		socmedRouter.DELETE("/:socmedId", middlewares.SocmedAuthorization(), controllers.DeleteSocmed)
	}

	commentRouter := r.Group("/comment")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/", controllers.CreateComment)
		commentRouter.GET("/", controllers.GetAllComment)
		commentRouter.GET("/:commentId", controllers.GetCommentById)
		commentRouter.PUT("/:commentId", middlewares.CommentAuthorization(), controllers.UpdateComment)
		commentRouter.DELETE("/:commentId", middlewares.CommentAuthorization(), controllers.DeleteComment)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
