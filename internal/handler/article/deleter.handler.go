package article

import (
	"github.com/gin-gonic/gin"
	"mentedu-backend/internal/service/article"
	userService "mentedu-backend/internal/service/user"
	"mentedu-backend/utils"
	"strconv"
)

type ArticleDeleterHandler struct {
	articleDeleter article.ArticleDeleterUsecase
	userFinder     userService.UserFinderUseCase
}

func NewArticleDeleterHandler(articleDeleter article.ArticleDeleterUsecase, userFinder userService.UserFinderUseCase) *ArticleDeleterHandler {
	return &ArticleDeleterHandler{
		articleDeleter: articleDeleter,
		userFinder:     userFinder,
	}
}

func (adh *ArticleDeleterHandler) DeleteArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, err)
		return
	}

	err = adh.articleDeleter.DeleteArticle(c.Request.Context(), id)

	if err != nil {
		c.JSON(400, err)
		return
	}

	c.JSON(200, utils.SuccessApiResponse("Article deleted successfully"))
}

func (adh *ArticleDeleterHandler) DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, err)
		return
	}

	err = adh.articleDeleter.DeleteCategory(c.Request.Context(), id)

	if err != nil {
		c.JSON(400, err)
		return
	}

	c.JSON(200, utils.SuccessApiResponse("Category deleted successfully"))
}
