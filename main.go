package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type TODO struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

var DB *gorm.DB

type User struct {
	gorm.Model
	Name   string
	Age    int64
	Active bool
}

func initMySql() (err error) {
	dsn := "todo:yh192610..@tcp(124.70.69.229:3306)/todo?charset=utf8mb4&parseTime=True"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func main() {
	//连接数据库 使用gorm
	err := initMySql()
	if err != nil {
		panic(err)
	}
	defer DB.Close()
	DB.AutoMigrate(&TODO{})

	r := gin.Default()
	r.Static("static", "static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1Group := r.Group("v1")
	{
		//增  POST
		v1Group.POST("/todo", func(c *gin.Context) {
			var todo TODO
			c.BindJSON(&todo)
			err = DB.Create(&todo).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"code": 2000,
					"msg":  "success",
					"data": todo,
				})
			}
		})
		//删  DELETE
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id, _ := c.Params.Get("id")
			var todo TODO
			err := DB.Where("id=?", id).First(&todo).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				DB.Delete(todo)
			}

		})
		//改  PUT
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			id, _ := c.Params.Get("id")
			var todo TODO
			err := DB.Where("id=?", id).First(&todo).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			}
			c.BindJSON(&todo)
			if err = DB.Save(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		//查  GET
		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})
		//查所有
		v1Group.GET("/todo", func(c *gin.Context) {
			var todoList []TODO
			err := DB.Find(&todoList).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todoList)
			}
		})
	}

	r.Run()
}
