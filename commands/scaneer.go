package commands

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/k0kubun/pp"
)

type scaneer struct{
	list *ToDoList
	events []event
}

type event struct{
	Title string
	Text string
	Time time.Time
}

func NewEvent(title, text string) *event{
	return &event{
		Title: title,
		Text: text,
		Time: time.Now(),
	}
}


func NewScaneer(l *ToDoList) *scaneer{
	return &scaneer{
		list: l,
		events: []event{},
	}
}


func (s *scaneer) Start(){
	defer func(){
		if p := recover(); p != nil{
			fmt.Println(p)
		}
	}()
	ev := &event{
		Title: "",
		Text: "",
		Time: time.Now(),
	}
	for{
		scr := bufio.NewScanner(os.Stdin)
		fmt.Print("Введите команду: ")
		scr.Scan()
		comm := strings.Split(scr.Text(), " ")
		comm1 := comm[0]
		var title string
		var text string


		switch comm1{
		case "add":
		if len(comm) < 3{
			fmt.Println(errors.New("Переданы не верные аргументы"))
			ev.Text = "Переданы не верные аргументы"
			ev.Title = "add"
			s.events = append(s.events, *ev)
			continue
		}else{
			title = comm[1]
			text = strings.Join(comm[2:], " ")
			ev.Title = title
			ev.Text = ""
			s.events = append(s.events, *ev)
			s.list.tasks[comm[1]] = *NewDo(title, text)
		}
		case "done":
			if len(comm) < 2 || len(comm) >= 3{
				fmt.Println(errors.New("Переданы не верные аргументы"))
				ev.Text = "Переданы не верные аргументы"
				ev.Time = time.Now()
				s.events = append(s.events, *ev)
				continue
			}else{
				title = comm[1]
				for k, v := range s.list.tasks{
					if v.title == title{
						v.Done()
						s.list.tasks[k] = v
					}
				}
				ev.Title = title
				s.events = append(s.events, *ev)
			}
		case "list":
		s.list.List()
		case "events":
			pp.Println(s.events)
		case "del":
			if len(comm) < 2 || len(comm) >= 3{
				fmt.Println(errors.New("Переданы не верные аргументы"))
				ev.Text = "Переданы не верные аргументы"
				ev.Time = time.Now()
				s.events = append(s.events, *ev)
				continue
			}else{
				title = comm[1]
				for k, v := range s.list.tasks{
					if v.title == title{
						delete(s.list.tasks, k)
					}
				}
				ev.Title = title
				s.events = append(s.events, *ev)
			}
		}
	}
}
