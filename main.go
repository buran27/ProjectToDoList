package main

import (
	"ProjectToDoList/commands"
	"bufio"
	"fmt"
	"os"

	"github.com/k0kubun/pp"
)



func main(){

	List := commands.NewToDoList()

	scaneer := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите команду: ")
	scaneer.Scan()
	switch scaneer.Text(){
	case "add":
		fmt.Print("Введите заголовок: ")
		scaneer.Scan()
		title := scaneer.Text()
		fmt.Print("Введите текст: ")
		scaneer.Scan()
		text := scaneer.Text()
		commands.Add(title, text, List)
	}
	pp.Println(*List)


}