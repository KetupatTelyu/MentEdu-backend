package article

import (
	"github.com/gin-gonic/gin"
	"mentedu-backend/internal/app/middleware"
	"mentedu-backend/internal/service/article"
	userService "mentedu-backend/internal/service/user"
	"mentedu-backend/resource"
	"mentedu-backend/utils"
	"strconv"
)

type ArticleUpdater struct {
	articleService article.ArticleUpdaterUsecase
	articleFinder  article.ArticleFinderUsecase
	userFinder     userService.UserFinderUseCase
	cloudStorage   utils.CloudStorage
}

func NewArticleUpdater(articleService article.ArticleUpdaterUsecase, userFinder userService.UserFinderUseCase, cloudStorage utils.CloudStorage) *ArticleUpdater {
	return &ArticleUpdater{
		articleService: articleService,
		userFinder:     userFinder,
		cloudStorage:   cloudStorage,
	}
}

func (au *ArticleUpdater) UpdateArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, err)
		return
	}

	currentArticle, err := au.articleFinder.FindByID(c.Request.Context(), id)

	var request resource.UpdateArticleRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(400, err)
		return
	}

	newImg := currentArticle.Image

	if currentArticle.Image != request.Image.Filename {
		imagePath, err := au.cloudStorage.UploadFile(request.Image, "articles/article/image")

		if err != nil {
			c.JSON(400, err)
			return
		}

		newImg = imagePath
	}

	if err != nil {
		c.JSON(400, err)
		return
	}

	user, err := au.userFinder.FindUser(c, middleware.UserID)

	if err != nil {
		c.JSON(400, err)
		return
	}

	newArticle, err := au.articleService.UpdateArticle(c.Request.Context(), id, request.Title, request.Body, request.Slug, newImg, request.CategoryID, user.Email)

	if err != nil {
		c.JSON(400, err)
		return
	}

	c.JSON(200, newArticle)
}

func (au *ArticleUpdater) UpdateCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, err)
		return
	}

	var request resource.UpdateCategoryRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(400, err)
		return
	}

	user, err := au.userFinder.FindUser(c, middleware.UserID)

	if err != nil {
		c.JSON(400, err)
		return
	}

	err = au.articleService.UpdateCategory(c.Request.Context(), id, request.Name, user.Email)

	if err != nil {
		c.JSON(400, err)
		return
	}

	c.JSON(200, gin.H{
		"message": "Category updated successfully",
	})
}
