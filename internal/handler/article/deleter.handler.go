package article

import (
	"github.com/gin-gonic/gin"
	"mentedu-backend/internal/service/article"
	userService "mentedu-backend/internal/service/user"
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

	c.JSON(200, gin.H{
		"message": "Article deleted successfully",
	})
}
