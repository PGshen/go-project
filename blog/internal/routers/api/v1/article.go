package v1

import (
	"github.com/PGshen/go-project/blog/pkg/app"
	"github.com/PGshen/go-project/blog/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Article struct {

}

func NewArticle() Article {
	return Article{}
}

func (a Article)Get(c *gin.Context)  {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
}

func (a Article)List(c *gin.Context)  {
	app.NewResponse(c).ToResponse("Hello")
}

func (a Article)Create(c *gin.Context)  {

}

func (a Article)Update(c *gin.Context)  {

}

func (a Article)Delete(c *gin.Context)  {

}
