package routers

import (
	"net/http"
	"todo/controller"
	"github.com/gin-gonic/gin"
)

func StartRoute()(*gin.Engine){
	r := gin.Default()
	r.Static("static", "static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1Group := r.Group("v1")
	{
		//增  POST
		v1Group.POST("/todo", controller.CreateTode)
		//删  DELETE
		v1Group.DELETE("/todo/:id",controller.DelectTodo )
		//改  PUT
		v1Group.PUT("/todo/:id", controller.UpdateTodo)
		//查
		v1Group.GET("/todo", controller.GetTodofunc)
	}
	return r
}