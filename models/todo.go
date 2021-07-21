package models

import (

	"todo/dao")


type TODO struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

type User struct {
	Name   string
	Age    int64
	Active bool
}

func CreatTodo(todo TODO) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}

func DeleteTodo(id string)(err error){
	var todo TODO
	err = dao.DB.Where("id=?", id).First(&todo).Error
	return
}

func Update(id string)(err error){
	var todo TODO
	err = dao.DB.Save(&todo).Error
	return
}

func GetTodo(todo TODO)(todoList[]*TODO ,err error){
	err = dao.DB.Find(&todoList).Error
	return
}