package main

import (
	"ProjectToDoList/commands"
)



func main(){

	List := commands.NewToDoList()

	scaneer := commands.NewScaneer(List)

	scaneer.Start()
	

}