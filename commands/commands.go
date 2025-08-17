package commands

import (
	"time"

	"github.com/k0kubun/pp"
)




type Do struct{
	title string
	text string
	isDone bool
	StartTime time.Time
	endTime *time.Time

}

func NewDo(title, text string) *Do{
	return &Do{
		title: title,
		text: text,
		isDone: false,
		StartTime: time.Now(),
		endTime: nil,
	}
}


type ToDoList struct{
	tasks map[string]Do
}



func NewToDoList() *ToDoList{
	return &ToDoList{
		tasks: make(map[string]Do),
	}
}
/*Реализация команды help. Команда позволяет узнать все доступные команды*/
func Help(){

}



/*Реализация команды add. Добовляет задачу в лист.
Принимает в себя заголовок задачи, указатель на список дел и описание*/

func (l *ToDoList) Add(title, text string){
	task := NewDo(title, text)
	l.tasks[title] = *task
}

func (l *ToDoList) List(){
	pp.Println(l.tasks)
}


func (d *Do) Done(){
	doneTime := time.Now()
	d.isDone = true
	d.endTime = &doneTime

}