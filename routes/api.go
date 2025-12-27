package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jrione/go-app-boilerplate/controller"
	"github.com/jrione/go-app-boilerplate/plugin"
	"github.com/jrione/go-app-boilerplate/repository"
)

func SetupRoutes(r *gin.Engine, logger *plugin.Logger, db *plugin.Database) {
	userRepo := repository.NewGormUserRepository(db.DB)
	userController := controller.NewUserController(userRepo, logger)

	api := r.Group("/api")
	{
		api.GET("/health", func(c *gin.Context) {
			logger.Info("Health check requested")
			c.JSON(200, gin.H{"status": "ok"})
		})

		api.GET("/users/:id", userController.GetUser)
		api.POST("/users", userController.CreateUser)
	}
}
