package article

import (
	"github.com/gin-gonic/gin"
	"mentedu-backend/common"
	"mentedu-backend/internal/app/middleware"
	"mentedu-backend/internal/service/article"
	userService "mentedu-backend/internal/service/user"
	"mentedu-backend/resource"
	"mentedu-backend/utils"
)

type ArticleCreatorHandler struct {
	articleCreatorUsecase article.ArticleCreatorUsecase
	userFinder            userService.UserFinderUseCase
	cloudStorage          utils.CloudStorage
}

func NewArticleCreatorHandler(articleCreatorUsecase article.ArticleCreatorUsecase, userFinder userService.UserFinderUseCase, cloudStorage utils.CloudStorage) *ArticleCreatorHandler {
	return &ArticleCreatorHandler{
		articleCreatorUsecase: articleCreatorUsecase,
		userFinder:            userFinder,
		cloudStorage:          cloudStorage,
	}
}

func (ach *ArticleCreatorHandler) CreateArticle(c *gin.Context) {
	var request resource.CreateArticleRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	imagePath, err := ach.cloudStorage.UploadFile(request.Image, "articles/article/image")

	if err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	user, err := ach.userFinder.FindUser(c, middleware.UserID)

	if err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	a, err := ach.articleCreatorUsecase.CreateArticle(c.Request.Context(), request.Title, request.Body, request.Slug, imagePath, request.CategoryID, user.Email)

	if err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	c.JSON(200, utils.SuccessApiResponse(a))
}

func (ach *ArticleCreatorHandler) CreateCategory(c *gin.Context) {
	var request resource.CreateCategoryRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	user, err := ach.userFinder.FindUser(c, middleware.UserID)

	if err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	err = ach.articleCreatorUsecase.CreateCategory(c.Request.Context(), request.Name, user.Email)

	if err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	c.JSON(200, utils.SuccessApiResponse("Category created successfully"))
}
