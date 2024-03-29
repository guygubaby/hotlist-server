package router

import (
	"github.com/gin-gonic/gin"
	"github.com/guygubaby/hotlist-server/controller"
	"net/http"
	"strconv"
)

func RootRoute(c *gin.Context)  {
	c.JSON(http.StatusOK,gin.H{
		"code":0,
		"data": "pong",
	})
}

func HotListRoute(c *gin.Context) {
	queryCate := c.DefaultQuery("cate","1")
	cate,_ := strconv.Atoi(queryCate)
	res := controller.GetList(cate)
	c.JSON(http.StatusOK,gin.H{
		"code":0,
		"data":res,
	})
}

func TypesRoute(c *gin.Context) {
	res := controller.InitAllTypes()
	c.JSON(http.StatusOK,gin.H{
		"code":0,
		"data": res,
	})
}

func InitRouter() *gin.Engine {
	r :=gin.New()
	api := r.Group("/api")

	api.GET("/",RootRoute)
	api.GET("/types", TypesRoute)
	api.GET("/list",HotListRoute)

	return r
}
