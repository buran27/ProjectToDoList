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
	list ToDoList
	events []event
}

type event struct{
	title string
	text string
	time time.Time
}

func NewEvent(title, text string) *event{
	return &event{
		title: title,
		text: text,
		time: time.Now(),
	}
}


func NewScaneer(l *ToDoList) *scaneer{
	return &scaneer{
		list: *l,
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
		title: "",
		text: "",
		time: time.Now(),
	}
	for{
		scr := bufio.NewScanner(os.Stdin)
		fmt.Print("Введите команду: ")
		scr.Scan()
		comm := strings.Split(scr.Text(), " ")
		comm1 := comm[0]
		var title string
		var text string

		if comm1 == "add"{
			if len(comm) < 3{
				fmt.Println(errors.New("Переданы не верные аргументы"))
				ev.text = "Переданы не верные аргументы"
				ev.text = ""
				s.events = append(s.events, *ev)
				continue
			}else{
				title = comm[1]
				text = strings.Join(comm[2:], " ")
			}
		}
		switch comm1{
		case "add":
		s.events = append(s.events, *ev)
		s.list.tasks[comm[1]] = Do{
			title: title,
			text: text,
			startTime: time.Now(),
			endTime: nil,
		}
		case "list":
		s.list.List()
		case "events":
			pp.Println(s.events)
		}
	}
}
