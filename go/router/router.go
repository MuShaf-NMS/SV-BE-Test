package router

import (
	"github.com/MuShaf-NMS/SV-BE-Test/config"
	"github.com/MuShaf-NMS/SV-BE-Test/handler"
	"github.com/MuShaf-NMS/SV-BE-Test/repository"
	"github.com/MuShaf-NMS/SV-BE-Test/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitializeRoute(server *gin.Engine, db *gorm.DB, config config.Config) {
	// Define Repositories
	postRepository := repository.NewPostRepository(db)

	// Define Services
	postService := service.NewPostService(postRepository)

	// Define Handlers
	postHandler := handler.NewPostHandler(postService)

	// post router
	postRouter := server.Group("/post")
	{
		postRouter.GET("/:id/:offset", postHandler.GetAll)               // jika menggunakan :limit akan konflik dengan :id
		postRouter.GET("/:id/:offset/:status", postHandler.GetAllFilter) // jika menggunakan :limit akan konflik dengan :id
		postRouter.POST("", postHandler.Create)
		postRouter.GET("/:id", postHandler.GetOne)
		postRouter.PUT("/:id", postHandler.Update)
		postRouter.DELETE("/:id", postHandler.Delete)
	}

}
