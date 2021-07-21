package controller

import (
	"net/http"
	"todo/dao"
	"todo/models"

	"github.com/gin-gonic/gin"
)


func CreateTode(c *gin.Context) {
	var todo models.TODO
	c.BindJSON(&todo)
	err := models.CreatTodo(todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": todo,
		})
	}
}

func DelectTodo(c *gin.Context) {
	id, _ := c.Params.Get("id")
	var todo models.TODO
	err := models.DeleteTodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		dao.DB.Delete(todo)
	}

}

func GetTodofunc(c *gin.Context) {
	var todo models.TODO
	todoList,err:= models.GetTodo(todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func UpdateTodo(c *gin.Context) {
	id, _ := c.Params.Get("id")
	var todo models.TODO
	err := models.DeleteTodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(&todo)
	err = models.Update(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}