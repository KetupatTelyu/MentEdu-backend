package resource

import "github.com/gin-gonic/gin"

type QueryRequest struct {
	Query  string `json:"query" form:"query"`
	Sort   string `json:"sort" form:"sort"`
	Order  string `json:"order" form:"order"`
	Limit  int    `json:"limit" form:"limit"`
	Offset int    `json:"offset" form:"offset"`
}

func NewQueryRequest(c *gin.Context) *QueryRequest {
	return &QueryRequest{
		Query:  c.Query("query"),
		Sort:   c.Query("sort"),
		Order:  c.Query("order"),
		Limit:  c.GetInt("limit"),
		Offset: c.GetInt("offset"),
	}
}
