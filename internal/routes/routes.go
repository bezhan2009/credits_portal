package routes

import (
	_ "Credits/docs"
	"Credits/internal/controllers"
	"Credits/internal/controllers/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func InitRoutes(r *gin.Engine) *gin.Engine {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Static("/uploads", "./uploads")

	creditsRoute := r.Group("/credits")
	{
		creditsRoute.GET("", controllers.GetAllCredits)
		creditsRoute.GET("/:id", controllers.GetCreditById)
		creditsRoute.POST("", middlewares.SavePassportFiles, controllers.CreateCredit)
		creditsRoute.PATCH("/:id", controllers.UpdateCredit)
		creditsRoute.DELETE("/:id", controllers.DeleteCredit)
	}

	return r
}
