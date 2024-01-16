package article

import (
	"github.com/gin-gonic/gin"
	"mentedu-backend/internal/service/article"
	"mentedu-backend/resource"
	article2 "mentedu-backend/responses"
	"net/http"
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

	response := article2.NewResponse(article, http.StatusOK, "success")

	c.JSON(200, response)
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

	var articleResponses []*article2.ArticleResponse

	for _, article := range articles {
		categoryID, err := af.articleFinder.FindArticleCategoryByArticleID(c.Request.Context(), article.ID)

		if err != nil {
			c.JSON(400, err)
			return
		}

		categoryName, err := af.articleFinder.FindCategoryByID(c.Request.Context(), categoryID[0].CategoryID)

		if err != nil {
			c.JSON(400, err)
			return
		}

		newArticle := article2.FromArticle(article)

		newArticle.Category = categoryName.Name

		articleResponses = append(articleResponses, newArticle)
	}

	response := article2.NewResponse(articleResponses, http.StatusOK, "success")

	c.JSON(200, response)
}

func (af *ArticleFinder) FindArticleBySlug(c *gin.Context) {
	slug := c.Param("slug")

	article, err := af.articleFinder.FindBySlug(c.Request.Context(), slug)

	if err != nil {
		c.JSON(400, err)
		return
	}

	c.JSON(200, article)
}

func (af *ArticleFinder) FindCategoryByID(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, err)
		return
	}

	category, err := af.articleFinder.FindCategoryByID(c.Request.Context(), categoryID)

	if err != nil {
		c.JSON(400, err)
		return
	}

	response := article2.NewResponse(category, http.StatusOK, "success")

	c.JSON(200, response)
}

func (af *ArticleFinder) FindCategories(c *gin.Context) {
	var request resource.QueryRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(400, err)
		return
	}

	categories, err := af.articleFinder.FindAllCategory(c.Request.Context(), request.Query, request.Sort, request.Order, request.Limit, request.Offset)

	if err != nil {
		c.JSON(400, err)
		return
	}

	response := article2.NewResponse(categories, http.StatusOK, "success")

	c.JSON(200, response)
}

func (af *ArticleFinder) FindArticleByCategoryID(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, err)
		return
	}

	articles, err := af.articleFinder.FindArticleByCategoryID(c.Request.Context(), categoryID)

	if err != nil {
		c.JSON(400, err)
		return
	}

	var articleResponses []*article2.ArticleResponse

	for _, article := range articles {
		categoryName, err := af.articleFinder.FindCategoryByID(c.Request.Context(), categoryID)

		if err != nil {
			c.JSON(400, err)
			return
		}

		newArticle := article2.FromArticle(article)

		newArticle.Category = categoryName.Name

		articleResponses = append(articleResponses, newArticle)
	}

	response := article2.NewResponse(articleResponses, http.StatusOK, "success")

	c.JSON(200, response)
}
