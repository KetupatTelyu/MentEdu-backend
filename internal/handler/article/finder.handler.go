package article

import (
	"github.com/gin-gonic/gin"
	"mentedu-backend/internal/service/article"
	"mentedu-backend/resource"
	"strconv"
)

type ArticleFinder struct {
	articleFinder article.ArticleFinderUsecase
}

func NewArticleFinder(articleFinder article.ArticleFinderUsecase) *ArticleFinder {
	return &ArticleFinder{
		articleFinder: articleFinder,
	}
}

func (af *ArticleFinder) FindArticleByID(c *gin.Context) {
	articleID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, err)
		return
	}

	article, err := af.articleFinder.FindByID(c.Request.Context(), articleID)

	if err != nil {
		c.JSON(400, err)
		return
	}

	c.JSON(200, article)
}

func (af *ArticleFinder) FindArticles(c *gin.Context) {
	var request resource.QueryRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(400, err)
		return
	}

	articles, err := af.articleFinder.FindAll(c.Request.Context(), request.Query, request.Sort, request.Order, request.Limit, request.Offset)

	if err != nil {
		c.JSON(400, err)
		return
	}

	c.JSON(200, articles)
}
