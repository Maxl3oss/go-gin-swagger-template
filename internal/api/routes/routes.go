package routes

import (
	"net/http"
	"role-management/docs"
	"role-management/internal/api/handlers"
	"role-management/internal/api/repository"
	"role-management/internal/config"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

func SetupRouter(cfg *config.Config, db *gorm.DB) *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/example")
		{
			eg.GET("/helloworld", Helloworld)
		}

		userRepo := repository.NewUserRepository(db)
		userHandler := handlers.NewUserHandler(userRepo)
		ug := v1.Group("/users")
		{
			ug.POST("/", userHandler.Create)
			ug.GET("/", userHandler.GetAll)
			ug.GET("/:id", userHandler.GetByID)
			ug.PUT("/:id", userHandler.Update)
			ug.DELETE("/:id", userHandler.Delete)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
