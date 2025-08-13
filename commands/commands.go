package commands

import (
	"fmt"
	"strconv"
	"time"
)

type ToDoList struct{
	Title string
	Text string
	Status string
	Time time.Time
}



func NewToDoList() *[]ToDoList{
	return &[]ToDoList{}
}
/*Реализация команды help. Команда позволяет узнать все доступные команды*/
func Help(){

}



/*Реализация команды add. Добовляет задачу в лист.
Принимает в себя заголовок задачи, указатель на список дел и опиисание*/

func Add(title, text string, list *[]ToDoList){
	*list = append(*list, ToDoList{
		Title: title,
		Text: text,
		Status: "Не выполнено",
		Time: time.Now(),
	})
}

func List(list []ToDoList){
	for i, v := range list{
		fmt.Println(strconv.Itoa(i+1) + ".         ", v)
	}
}