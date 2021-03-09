package models

import (
	"Lesson26/dao"
)

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// Todo 增删改查
func CreateTodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}

func GetTodolist() (todoList []*Todo, err error) {

	if err = dao.DB.Find(&todoList).Error; err != nil {
		return todoList, err
	}
	return
}
func GetTodo(id string) (todo Todo, err error) {
	todo = Todo{}
	if err = dao.DB.Where("id=?", id).First(&todo).Error; err != nil {
		return Todo{}, err
	}
	return
}
func UpdateTodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return
}
func DeleteTodo(id string) (err error) {
	err = dao.DB.Where("id = ?", id).Delete(&Todo{}).Error
	return err
}
