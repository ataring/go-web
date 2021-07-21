package main

import (
	"todo/dao"
	"todo/models"
	"todo/routers"
)

func main() {
	//连接数据库 使用gorm
	err := dao.InitMySql()
	if err != nil {
		panic(err)
	}
	defer dao.DB.Close()
	dao.DB.AutoMigrate(&models.TODO{})

	
	r := routers.StartRoute()
	r.Run()
}
