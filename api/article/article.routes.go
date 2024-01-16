package article

import (
	"github.com/gin-gonic/gin"
	"mentedu-backend/internal/app/config"
	"mentedu-backend/internal/app/middleware"
	article2 "mentedu-backend/internal/handler/article"
	"mentedu-backend/internal/service/article"
	userService "mentedu-backend/internal/service/user"
	"mentedu-backend/utils"
)

func ArticleCreatorHTTPHandler(cfg config.Config, router *gin.Engine, ac article.ArticleCreatorUsecase, uf userService.UserFinderUseCase, cloudStorage utils.CloudStorage) {
	hndlr := article2.NewArticleCreatorHandler(ac, uf, cloudStorage)

	api := router.Group("/api")

	api.Use(middleware.Auth(cfg))
	api.Use(middleware.Admin(cfg))

	{
		api.POST("/cms/article", hndlr.CreateArticle)
	}
}

func ArticleFinderHTTPHandler(cfg config.Config, router *gin.Engine, af article.ArticleFinderUsecase, uf userService.UserFinderUseCase, cloudStorage utils.CloudStorage) {
	hndlr := article2.NewArticleFinder(af)

	api := router.Group("/api")
	{
		api.GET("/article", hndlr.FindArticles)
		api.GET("/article/category/:id", hndlr.FindArticleByCategoryID)
	}

	api.Use(middleware.Auth(cfg))
	{
		api.GET("/cms/article", hndlr.FindArticles)
		api.GET("/cms/article/category/:id", hndlr.FindArticleByCategoryID)
		api.GET("/cms/article/:id", hndlr.FindArticleByID)
	}
}

func ArticleUpdaterHTTPHandler(cfg config.Config, router *gin.Engine, au article.ArticleUpdaterUsecase, uf userService.UserFinderUseCase, cloudStorage utils.CloudStorage) {
	hndlr := article2.NewArticleUpdater(au, uf, cloudStorage)

	api := router.Group("/api")

	api.Use(middleware.Auth(cfg))
	api.Use(middleware.Admin(cfg))

	{
		api.PUT("/cms/article/:id", hndlr.UpdateArticle)
	}
}

func ArticleDeleterHTTPHandler(cfg config.Config, router *gin.Engine, ad article.ArticleDeleterUsecase, uf userService.UserFinderUseCase, cloudStorage utils.CloudStorage) {
	hndlr := article2.NewArticleDeleterHandler(ad, uf)

	api := router.Group("/api")

	api.Use(middleware.Auth(cfg))
	api.Use(middleware.Admin(cfg))

	{
		api.DELETE("/cms/article/:id", hndlr.DeleteArticle)
	}
}

func CategoryCreatorHTTPHandler(cfg config.Config, router *gin.Engine, cc article.ArticleCreatorUsecase, uf userService.UserFinderUseCase, cloudStorage utils.CloudStorage) {
	hndlr := article2.NewArticleCreatorHandler(cc, uf, cloudStorage)

	api := router.Group("/api")

	api.Use(middleware.Auth(cfg))
	api.Use(middleware.Admin(cfg))

	{
		api.POST("/cms/category", hndlr.CreateCategory)
	}
}

func CategoryFinderHTTPHandler(cfg config.Config, router *gin.Engine, cf article.ArticleFinderUsecase, uf userService.UserFinderUseCase, cloudStorage utils.CloudStorage) {
	hndlr := article2.NewArticleFinder(cf)

	api := router.Group("/api")

	api.Use(middleware.Auth(cfg))

	{
		api.GET("/cms/category", hndlr.FindCategories)
		api.GET("/cms/category/:id", hndlr.FindCategoryByID)
	}
}

func CategoryUpdaterHTTPHandler(cfg config.Config, router *gin.Engine, cu article.ArticleUpdaterUsecase, uf userService.UserFinderUseCase, cloudStorage utils.CloudStorage) {
	hndlr := article2.NewArticleUpdater(cu, uf, cloudStorage)

	api := router.Group("/api")

	api.Use(middleware.Auth(cfg))
	api.Use(middleware.Admin(cfg))

	{
		api.PUT("/cms/category/:id", hndlr.UpdateCategory)
	}
}

func CategoryDeleterHTTPHandler(cfg config.Config, router *gin.Engine, cd article.ArticleDeleterUsecase, uf userService.UserFinderUseCase, cloudStorage utils.CloudStorage) {
	hndlr := article2.NewArticleDeleterHandler(cd, uf)

	api := router.Group("/api")

	api.Use(middleware.Auth(cfg))
	api.Use(middleware.Admin(cfg))

	{
		api.DELETE("/cms/category/:id", hndlr.DeleteCategory)
	}
}
