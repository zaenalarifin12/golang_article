package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zaenalarifin12/golang_article/config"
	"github.com/zaenalarifin12/golang_article/controller"
	"github.com/zaenalarifin12/golang_article/middleware"
	"github.com/zaenalarifin12/golang_article/repository"
	"github.com/zaenalarifin12/golang_article/service"
	"gorm.io/gorm"
)

var (
	db         *gorm.DB           = config.SetupDatabaseConnection()
	jwtService service.JWTService = service.NewJWTService()
	// repository
	userRepository    repository.UserRepository    = repository.NewUserRepository(db)
	articleRepository repository.ArticleRepository = repository.NewArticleRepository(db)

	// services
	authService    service.AuthService    = service.NewAuthService(userRepository)
	articleService service.ArticleService = service.NewArticleService(articleRepository)
	//controller
	authController    controller.AuthController    = controller.NewAuthController(authService, jwtService)
	articleController controller.ArticleController = controller.NewArticleController(articleService, jwtService)
	imageController controller.ImageController = controller.NewImageController()
)

func main() {
	defer config.CloseDatabaseConnection(db)

	r := gin.Default()

	r.Static("/images", "./public/images")

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/auth/register", authController.Register)
		apiv1.POST("/auth/login", authController.Login)

		apiv1.GET("/articles", middleware.AuthorizeJWT(jwtService), articleController.GetAll)
		apiv1.POST("/article", middleware.AuthorizeJWT(jwtService), articleController.Insert)
		apiv1.PUT("/article/:id", middleware.AuthorizeJWT(jwtService), articleController.Update)
		apiv1.GET("/article/:id", middleware.AuthorizeJWT(jwtService), articleController.FindById)
		apiv1.DELETE("/article/:id", middleware.AuthorizeJWT(jwtService), articleController.Delete)
		// image
		apiv1.POST("/image", imageController.UploadImage)

	}

	r.Run(":4000")

}
